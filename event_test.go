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
