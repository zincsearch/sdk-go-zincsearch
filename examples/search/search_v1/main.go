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
	index := "index_example"          // string | Index
	query := *client.NewV1ZincQuery() // V1ZincQuery | Query
	query.SetSearchType("match")
	params := *client.NewV1QueryParams()
	params.SetTerm("John")
	params.SetField("name")
	query.SetQuery(params)
	query.SetSortFields([]string{"-@timestamp"})
	query.SetMaxResults(5)
	query.SetSource([]string{"age", "name"})

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
	resp, r, err := apiClient.Search.SearchV1(ctx, index).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchV1`: V1SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchV1`: %v\n", resp)
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%v %v\n", hit.GetTimestamp(), hit.GetSource())
	}
}
