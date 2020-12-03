package elasticsearch_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestEs(t *testing.T) {
	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)

	es := esClient(t)

	// 1. Get cluster info
	//
	res, err := es.Info()
	if err != nil {
		t.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		t.Fatalf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}
	if r["version"].(map[string]interface{})["number"] != "6.0.1" {
		t.Fatal("Server should run 6.0.1")
	}

	// 2. Index documents concurrently
	//
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			var b strings.Builder
			b.WriteString(`{"title" : "`)
			b.WriteString(title)
			b.WriteString(`"}`)

			// Set up the request object.
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				DocumentType: "doc",
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				t.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				t.Fatal(res.String())
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					t.Fatalf("Error parsing the response body: %s", err)
				}
				// Print the response status and indexed document version.
				if res.StatusCode < 200 || 300 <= res.StatusCode {
					t.Errorf("res.StatusCode wants 200/201, but got %d", res.StatusCode)
				}
			}
		}(i, title)
	}
	wg.Wait()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		t.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		t.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			t.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			t.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("res.StatusCode wants %d, but got %d", 200, res.StatusCode)
	}
	if r["hits"].(map[string]interface{})["total"].(float64) != 2 {
		t.Errorf("hits want %d, but got %f", 2, r["hits"].(map[string]interface{})["total"].(float64))
	}
}

func esClient(t *testing.T) *elasticsearch.Client {
	t.Helper()

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Fatalf("Error creating the client: %s", err)
	}

	return es
}
