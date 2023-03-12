package main

import (
	"context"
	"log"

	"github.com/merajsahebdar/bookmarkmanager/pkg/notion"
)

func main() {
	client := notion.New()

	searchResponse, err := client.Search(context.Background(), notion.SearchRequest{
		Query:  "Bookmark List",
		Filter: notion.SearchFilter{Value: "database", Property: "object"},
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%v\n", searchResponse)
}
