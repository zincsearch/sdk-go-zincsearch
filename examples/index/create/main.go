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
	index := *client.NewMetaIndexSimple() // MetaIndexSimple | Index data
	index.SetName("index_example")

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
	resp, r, err := apiClient.Index.Create(ctx).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Index.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: MetaHTTPResponseIndex
	fmt.Fprintf(os.Stdout, "Response from `Index.Create`: %v\n", resp)
	if r.StatusCode == 200 {
		fmt.Fprintf(os.Stdout, "`Index.Create`: %v\n", resp.GetIndex())
	} else {
		e, _ := err.(*client.GenericOpenAPIError)
		me, _ := e.Model().(client.MetaHTTPResponseError)
		fmt.Fprintf(os.Stdout, "`Index.Create` error: %v\n", me.GetError())
	}
}
