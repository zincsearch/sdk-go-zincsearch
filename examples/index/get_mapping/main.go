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
	resp, r, err := apiClient.Index.GetMapping(ctx, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Index.GetMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMapping`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `Index.GetMapping`: %v\n", resp)
}
