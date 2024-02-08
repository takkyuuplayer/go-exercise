package bigquery_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func TestClient(t *testing.T) {
	t.Parallel()

	t.Run("CRUD", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		client, err := bigquery.NewClient(
			ctx,
			"test",
			option.WithEndpoint(os.Getenv("BIGQUERY_ENDPOINT")),
			option.WithoutAuthentication(),
		)
		assert.NoError(t, err)
		t.Cleanup(func() {
			assert.NoError(t, client.Close())
		})

		// Create dataset
		dataset := client.Dataset("dataset_" + ulid.Make().String())
		assert.NoError(
			t,
			dataset.Create(context.Background(), nil),
		)
		t.Cleanup(func() {
			assert.NoError(t, dataset.Delete(context.Background()))
		})

		// List datasets
		{
			it := client.Datasets(ctx)
			for {
				ds, err := it.Next()
				if errors.Is(err, iterator.Done) {
					break
				}
				assert.NoError(t, err)
				t.Log(ds.DatasetID)
			}
		}

		// Create table
		type item struct {
			ID   int    `bigquery:"id"`
			Name string `bigquery:"name"`
			Age  int    `bigquery:"age"`
		}
		schema, err := bigquery.InferSchema(item{})
		assert.NoError(t, err)

		table := client.Dataset(dataset.DatasetID).Table("tmp_" + ulid.Make().String())
		assert.NoError(
			t,
			table.Create(ctx, &bigquery.TableMetadata{
				Schema: schema,
			}),
		)

		// List tables
		{
			it := dataset.Tables(ctx)
			for {
				tbl, err := it.Next()
				if errors.Is(err, iterator.Done) {
					break
				}
				assert.NoError(t, err)
				t.Log(tbl.TableID)
			}
		}

		//	Insert data
		{
			inserter := table.Inserter()
			assert.NoError(
				t,
				inserter.Put(ctx, []*item{
					{1, "Alice", 20},
					{2, "Bob", 18},
				}),
			)
		}

		// Query data
		{
			query := fmt.Sprintf("SELECT * FROM `%s.%s`", dataset.DatasetID, table.TableID)
			q := client.Query(query)
			it, err := q.Read(ctx)
			assert.NoError(t, err)
			for {
				var row []bigquery.Value
				err := it.Next(&row)
				if errors.Is(err, iterator.Done) {
					break
				}
				assert.NoError(t, err)
				t.Log(row)
			}
		}

		// Merge data
		t.Run("Merge", func(t *testing.T) {
			t.Skip("Not working in emulator")

			t2 := client.Dataset(dataset.DatasetID).Table("tmp_" + ulid.Make().String())
			assert.NoError(
				t,
				t2.Create(ctx, &bigquery.TableMetadata{Schema: schema}),
			)
			inserter := t2.Inserter()
			assert.NoError(
				t,
				inserter.Put(ctx, []*item{
					{2, "Bob", 19},
					{3, "Charlie", 20},
				}),
			)

			{
				q := client.Query(fmt.Sprintf(`CREATE OR REPLACE TABLE %s.copy_table AS SELECT * FROM %s.%s`, dataset.DatasetID, dataset.DatasetID, table.TableID))
				job, err := q.Run(ctx)
				assert.NoError(t, err)

				status, err := job.Wait(ctx)
				assert.NoError(t, err)
				assert.Equal(t, bigquery.Done, status.Err())
			}

			q := client.Query(fmt.Sprintf(`MERGE %s.copy_table dest
USING %s.%s src ON dest.id = src.id
WHEN MATCHED THEN
	UPDATE SET name = src.name, age = src.age
WHEN NOT MATCHED THEN
	INSERT (id, name, age) VALUES (src.id, src.name, src.age)
`, dataset.DatasetID, dataset.DatasetID, t2.TableID))

			job, err := q.Run(ctx)
			assert.NoError(t, err)

			status, err := job.Wait(ctx)
			assert.NoError(t, err)
			assert.Equal(t, bigquery.Done, status.Err())

			it, err := job.Read(ctx)
			for {
				var row []bigquery.Value
				err := it.Next(&row)
				if errors.Is(err, iterator.Done) {
					break
				}
				assert.NoError(t, err)
				fmt.Println(row)
			}

			// Query data
			{
				query := fmt.Sprintf("SELECT * FROM `%s.copy_table`", dataset.DatasetID)
				q := client.Query(query)
				it, err := q.Read(ctx)
				assert.NoError(t, err)
				for {
					var row []bigquery.Value
					err := it.Next(&row)
					if errors.Is(err, iterator.Done) {
						break
					}
					assert.NoError(t, err)
					t.Log(row)
				}
			}
		})
	})
}
