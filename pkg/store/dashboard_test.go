/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package store

import (
	"os"
	"testing"

	"github.com/macroadster/hicc/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestLoadDashboard(t *testing.T) {
	assert := assert.New(t)
	dashboard := GetDashboard("nonexisted")
	assert.Equal(len(dashboard.Grid), 1, "Dashboard should contain one widget")
}

func TestSaveDashboard(t *testing.T) {
	err := os.MkdirAll("data/dashboard", 0700)
	if err != nil {
		t.Error("Failed to create test directory")
	}
	defer os.RemoveAll("data")
	dashboard := common.Dashboard{}
	err = SaveDashboard("tmp", dashboard)
	if err != nil {
		t.Error(err)
	}
	_, err = os.Stat("data/dashboard/tmp.json")
	if os.IsNotExist(err) {
		t.Error("Expected tmp dashboard to exist")
	}
}
