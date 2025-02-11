// Copyright (c) 2021 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"github.com/networkservicemesh/integration-tests/extensions/checkout"
	"github.com/networkservicemesh/integration-tests/extensions/logs"
)

// Suite is a base suite for generating tests. Contains extensions that can be used for assertion and automation goals.
type Suite struct {
	checkout.Suite
	storeTestLogs, storeSuiteLogs func()
	// Add other extensions here
}

func (s *Suite) SetupSuite() {
	s.Repository = "networkservicemesh/deployments-k8s"
	s.Version = "dcaced16"
	s.Dir = "../" // Note: this should be synced with input parameters in gen.go file

	s.Suite.SetupSuite()

	s.storeSuiteLogs = logs.Capture(s.T().Name())
}

func (s *Suite) AfterTest(_, _ string) {
	s.storeTestLogs()
}

func (s *Suite) BeforeTest(_, _ string) {
	s.storeTestLogs = logs.Capture(s.T().Name())
}

func (s *Suite) TearDownSuite() {
	s.storeSuiteLogs()
}
