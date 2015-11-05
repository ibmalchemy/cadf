// Copyright 2015 IBM Corp.
// 
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package cadf

import (
	"fmt"
)

func ExampleOmitempty() {
	event := Event{}
	fmt.Println(event.String())
	// Output:
	// {"eventTime":"0001-01-01T00:00:00Z","initiator":{"host":{},"credential":{}},"target":{"host":{},"credential":{}},"observer":{"host":{},"credential":{}},"reason":{},"api":{"createdAt":"0001-01-01T00:00:00Z"},"latencies":{}}
}
