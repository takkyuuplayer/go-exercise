package elasticsearch_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
)

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v6"
)

func TestEs(t *testing.T) {
	var (
		r  map[string]interface{}
		//wg sync.WaitGroup
	)

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		t.Fatal(err)
	}
	res, err := es.Info()
	if err != nil {
		t.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		t.Fatalf("Error: %s", res.String())
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))
}
