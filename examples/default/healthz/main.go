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
	configuration := client.NewConfiguration()
	apiClient := client.NewAPIClient(configuration)
	resp, r, err := apiClient.Default.Healthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Default.Healthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Healthz`: MetaHTTPResponse
	fmt.Fprintf(os.Stdout, "Response from `Default.Healthz`: %v\n", resp.GetStatus())
}
