// Copyright © 2020 Cisco
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
//
// All rights reserved.

package servregistry

import "errors"

var (
	// ErrServRegNotProvided is returned when the manager has no service
	// registry and, thus, cannot make any changes
	ErrServRegNotProvided error = errors.New("no service registry is provided")
	// ErrNotFound is returned when the resource does not exist
	ErrNotFound error = errors.New("resource not found")
	// ErrAlreadyExists is returned when the resource already exists
	ErrAlreadyExists error = errors.New("resource already exists")
)