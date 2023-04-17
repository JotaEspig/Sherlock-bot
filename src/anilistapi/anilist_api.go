package anilistapi

import (
	"context"
	"strings"

	"github.com/machinebox/graphql"
)

var url string = "https://graphql.anilist.co"

//Make a request to the API
//	Args:
//		query (string): request query
//		variables (map[string]interface{}): request params
//		target (interface{}): Must be a pointer to a response struct
func post(query string, variables map[string]interface{}, target interface{}) error {
	client := graphql.NewClient(url)
	req := graphql.NewRequest(query)
	for key, value := range variables {
		req.Var(key, value)
	}
	ctx := context.Background()
	err := client.Run(ctx, req, target)
	if err != nil {
		return err
	}

	return nil
}

// Check the values of "page" and "perPage", and changes their values if it's needed
func checkPageValues(page *int, perPage *int) {
	if *page < 1 {
		*page = 1
	}
	if *perPage > 20 || *perPage < 1 {
		*perPage = 10
	}
}

// Treats the description
func TreatDescription(description string, target *string) {
	desc := description
	// Treat Spoilers
	desc = strings.Replace(desc, "~!", "||", -1)
	desc = strings.Replace(desc, "!~", "||", -1)
	// HTML tags
	desc = strings.Replace(desc, "<br>", "\n", -1)
	desc = strings.Replace(desc, "<i>", "_", -1)
	desc = strings.Replace(desc, "</i>", "_", -1)
	*target = desc
}
