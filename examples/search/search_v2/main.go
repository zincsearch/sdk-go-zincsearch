/* Copyright 2022 Zinc Labs Inc. and Contributors
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"os"

	client "github.com/zinclabs/sdk-go-zincsearch"
)

func main() {
	index := "index_example" // string | Index
	// eg:
	// {
	// 	"query": {
	// 		"bool": {
	// 			"should": [
	// 				{
	// 					"match": {
	// 						"name": {
	// 							"query": "John"
	// 						}
	// 					}
	// 				}
	// 			]
	// 		}
	// 	}
	// }
	query := *client.NewMetaZincQuery() // V1ZincQuery | Query
	matchQuery := *client.NewMetaMatchQuery()
	matchQuery.SetQuery("John")
	subQuery := *client.NewMetaQuery()
	subQuery.SetMatch(map[string]client.MetaMatchQuery{
		"name": matchQuery,
	})
	boolQuery := *client.NewMetaBoolQuery()
	boolQuery.SetShould([]client.MetaQuery{subQuery})
	queryQuery := *client.NewMetaQuery()
	queryQuery.SetBool(boolQuery)
	query.SetQuery(queryQuery)

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
	resp, r, err := apiClient.Search.Search(ctx, index).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchV1`: V1SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchApi.Search`: %+v\n", resp)
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%v %v\n", hit.GetTimestamp(), hit.GetSource())
	}
}
