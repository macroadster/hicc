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
package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/macroadster/hicc/pkg/common"
	"github.com/macroadster/hicc/pkg/store"
)

func Rest() {
	r := gin.Default()
	r.Static("/hicc", "web")
	r.GET("/", RedirectToHicc)
	r.GET("/v1/session", GetSession)
	r.GET("/v1/dashboard/:id", GetDashboard)
	r.PUT("/v1/dashboard/:id", SaveDashboard)
	r.GET("/v1/widget", GetWidget)
	r.Run()
}

func RedirectToHicc(c *gin.Context) {
	c.Redirect(302, "/hicc/home/")
}

func GetSession(c *gin.Context) {
	c.JSON(200, gin.H{
		"user":      "Guest",
		"period":    "last1hr",
		"dashboard": "default",
	})
}

func GetDashboard(c *gin.Context) {
	dashboard := store.GetDashboard(c.Param("id"))
	c.JSON(200, dashboard)
}

func SaveDashboard(c *gin.Context) {
	dashboard := common.Dashboard{}
	err := c.BindJSON(&dashboard)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
	err = store.SaveDashboard(c.Param("id"), dashboard)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "saved",
		})
	} else {
		c.JSON(400, gin.H{
			"message": err,
		})
	}
}

func GetWidget(c *gin.Context) {
	widgets := store.GetWidgets()
	c.JSON(200, widgets)
}
