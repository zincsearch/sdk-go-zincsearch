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
	query := client.NewMetaJSONIngest()
	query.SetIndex("index_example")
	query.SetRecords([]map[string]interface{}{
		{"name": "John Doe", "age": 30, "address": "123 Main St"},
		{"name": "John Doe", "age": 31, "address": "123 Main St"},
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
	resp, r, err := apiClient.Document.Bulkv2(ctx).Query(*query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Document.Bulkv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Bulkv2`: MetaHTTPResponseRecordCount
	fmt.Fprintf(os.Stdout, "Response from `Document.Bulkv2`: %v\n", resp.GetRecordCount())
}
