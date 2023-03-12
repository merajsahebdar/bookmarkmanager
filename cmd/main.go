package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/merajsahebdar/bookmarkmanager/pkg/notion"
)

func main() {
	client := http.Client{}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"https://api.notion.com/v1/search",
		bytes.NewReader([]byte(`{"query":"Bookmark List","filter":{"value":"database","property":"object"}}`)),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_API_KEY"))
	req.Header.Add("Notion-Version", "2022-06-28")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	searchResponse := notion.SearchResponse{}

	if unmarshalingErr := json.NewDecoder(response.Body).Decode(&searchResponse); unmarshalingErr != nil {
		panic(unmarshalingErr)
	}

	log.Printf("%v\n", searchResponse)
}
