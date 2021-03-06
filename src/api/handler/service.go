// Copyright 2018 Naftis Authors
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

package handler

import (
	"github.com/xiaomi/naftis/src/api/service"
	"github.com/xiaomi/naftis/src/api/util"

	"github.com/gin-gonic/gin"
)

// Services returns all available services.
func Services(c *gin.Context) {
	t := c.Query("t")
	if t != "tree" {
		uid := c.Param("uid")
		c.JSON(200, gin.H{
			"code": 0,
			"data": service.ServiceInfo.Services(uid).Exclude("kube-system", "istio-system", "naftis").Status(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": service.ServiceInfo.Tree(),
	})
	return
}

// ServicePods queries pods's of specific pod by service UID.
func ServicePods(c *gin.Context) {
	uid := c.Param("uid")
	svcs := service.ServiceInfo.Services(uid)

	if len(svcs) == 0 || len(svcs[0].Labels) == 0 {
		c.JSON(200, util.RetOK)
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": service.ServiceInfo.Pods(svcs[0].Spec.Selector).Status(),
	})
}

// Pods queries pods's of specific pod by name.
func Pods(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"code": 0,
		"data": service.ServiceInfo.PodsByName(name).Exclude("kube-system", "istio-system", "naftis").Status(),
	})
}
