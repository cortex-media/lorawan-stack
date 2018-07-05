// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events_test

import (
	"fmt"

	"go.thethings.network/lorawan-stack/pkg/events"
)

func ExampleHandlerFunc() {
	handler := events.HandlerFunc(func(e events.Event) {
		fmt.Printf("Received event %v\n", e)
	})

	events.Subscribe("example", handler)

	// From this moment on, "example" events will be delivered to the handler func.

	events.Unsubscribe("example", handler)

	// Note that in-transit events may still be delivered after Unsubscribe returns.
}

func ExampleChannel() {
	eventChan := make(events.Channel, 16)

	go func() {
		for e := range eventChan {
			fmt.Printf("Received event %v\n", e)
		}
	}()

	events.Subscribe("example", eventChan)

	// From this moment on, "example" events will be delivered to the channel.

	events.Unsubscribe("example", eventChan)

	// Note that in-transit events may still be delivered after Unsubscribe returns.
	// This means that you can't immediately close the channel after unsubscribing.
}