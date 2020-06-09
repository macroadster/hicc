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
	"encoding/json"
	"io/ioutil"

	"github.com/macroadster/hicc/pkg/common"
)

func GetDashboard(id string) common.Dashboard {
	dashboard := common.Dashboard{
		Grid: []common.Widget{},
	}
	data, err := ioutil.ReadFile("data/dashboard/" + id + ".json")
	if err != nil {
		widget := common.Widget{
			Col:    0,
			Row:    0,
			SizeX:  12,
			SizeY:  7,
			Source: "/hicc/welcome.html",
		}
		dashboard = common.Dashboard{
			Grid: []common.Widget{widget},
		}
	}
	json.Unmarshal(data, &dashboard)
	return dashboard
}

func SaveDashboard(id string, dashboard common.Dashboard) error {
	path := "data/dashboard/" + id + ".json"
	data, _ := json.Marshal(dashboard)
	err := ioutil.WriteFile(path, data, 0600)
	return err
}
