package bigquery_test

import (
	"context"
	"errors"
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
		dataset := client.Dataset(ulid.Make().String())
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
		{
			table := client.Dataset(dataset.DatasetID).Table(ulid.Make().String())
			assert.NoError(
				t,
				table.Create(ctx, &bigquery.TableMetadata{
					Schema: bigquery.Schema{
						{Name: "name", Type: bigquery.StringFieldType},
						{Name: "age", Type: bigquery.IntegerFieldType},
					},
				}),
			)
		}
	})
}
