// Copyright 2019 annp.cs51@gmail.com
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

package backend

import (
	api "RESTovergRPC/directory"
)

// Backend stores and retrieves entities.
type Backend interface {
	// get user
	GetUser(cmd string) (string, error)

	// CreateDirectory
	CreateDirectory(name string) (string, error)

	// AddEntry and return string "ok" or "fail"
	AddEntry(e *api.EntryRequest) (string, error)

	// SearchEntry
	SearchEntry(query string) ([]*api.Entry, error)

	// Close handles any necessary cleanup
	Close() error
}
