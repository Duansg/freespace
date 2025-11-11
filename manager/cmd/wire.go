//go:build wireinject
// +build wireinject

/*
* Licensed to the Apache Software Foundation (ASF) under one
* or more contributor license agreements.  See the NOTICE file
* distributed with this work for additional information
* regarding copyright ownership.  The ASF licenses this file
* to you under the Apache License, Version 2.0 (the
* "License"); you may not use this file except in compliance
* with the License.  You may obtain a copy of the License at
*
*   http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */
package freespacecmd

import (
	"github.com/Duansg/freespace/manager/base/conf"
	"github.com/Duansg/freespace/manager/base/data"
	"github.com/Duansg/freespace/manager/base/handler"
	"github.com/Duansg/freespace/manager/router"
	"github.com/Duansg/freespace/manager/server"
	"github.com/Duansg/freespace/manager/service"
	service_config "github.com/Duansg/freespace/manager/service/config"
	"github.com/google/wire"
	"github.com/segmentfault/pacman"
	"github.com/segmentfault/pacman/log"
)

// initApplication init application.
func initApplication(
	debug bool,
	serverConf *conf.Server,
	dbConf *data.Database,
	swaggerConf *router.SwaggerConfig,
	serviceConf *service_config.ServiceConfig,
	logConf log.Logger) (*pacman.Application, func(), error) {
	panic(wire.Build(
		data.ProviderSetData,
		data.ProviderMappers,
		router.ProviderSetRouter,
		server.ProviderSetServer,
		service.ProviderSetService,
		handler.ProviderSetHandler,
		newApplication,
	))
}
