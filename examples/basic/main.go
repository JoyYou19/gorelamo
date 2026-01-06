package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JoyYou19/gorelamo"
	"github.com/JoyYou19/gorelamo/config"
)

func main() {
	client := gorelamo.NewClient(
		"http://10.10.10.71:82/api",
		config.WithBasicAuth("root", "CPpassword"),
	)

	// set root tag //document
	db := client.Database("gorelamo")

	// Insert a document
	err := db.Insert(map[string]any{
		"level": "info",
		"msg":   "connected successfully",
		"ts":    time.Now().Format(time.RFC3339),
	})
	if err != nil {
		log.Fatalf("insert failed: %v", err)
	}

	fmt.Println("Insert OK")

	// Search for the document
	res, err := db.Search("connected successfully").Do()
	if err != nil {
		log.Fatalf("search failed: %v", err)
	}

	fmt.Println("Search result (raw XML):")
	fmt.Println(res.Raw)
}
