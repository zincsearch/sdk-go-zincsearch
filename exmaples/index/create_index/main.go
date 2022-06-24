package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/zinclabs/sdk-go-zincsearch"
)

func main() {
	index := client.NewMetaIndexSimple()  // MetaIndexSimple | Index data
	index.SetName("index_example_create") // string | Index name
	index.SetStorageType("disk")
	index.SetMappings(map[string]interface{}{
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "text",
			},
		},
	})

	ctx := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
		UserName: "admin",
		Password: "Complexpass#123",
	})

	configuration := client.NewConfiguration()
	configuration.Servers = client.ServerConfigurations{
		client.ServerConfiguration{
			URL: "http://localhost:4080",
		},
	}
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.IndexApi.CreateIndex(ctx).Index(*index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexApi.CreateIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIndex`: MetaHTTPResponseIndex
	fmt.Fprintf(os.Stdout, "Response from `IndexApi.CreateIndex`: %v\n", resp)
}
