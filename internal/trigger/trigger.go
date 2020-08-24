//
// Copyright (c) 2020 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package trigger

import (
	"context"
	"github.com/edgexfoundry/go-mod-messaging/pkg/types"
	"sync"

	"github.com/edgexfoundry/go-mod-bootstrap/bootstrap"
)

// Trigger interface is used to hold event data and allow function to
type Trigger interface {
	// Initialize performs post creation initializations
	Initialize(wg *sync.WaitGroup, ctx context.Context, background <-chan types.MessageEnvelope) (bootstrap.Deferred, error)
}
