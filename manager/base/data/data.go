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

package data

import (
	"path/filepath"
	"time"

	"github.com/Duansg/freespace/manager/pkg/dir"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/segmentfault/pacman/log"
	_ "modernc.org/sqlite"
	"xorm.io/xorm"
	ormlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
	"xorm.io/xorm/schemas"
)

// Data data
type Data struct {
	DB *xorm.Engine
}

// NewData new data instance
func NewData(db *xorm.Engine) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
		err := db.Close()
		if err != nil {
			return
		}
	}
	return &Data{DB: db}, cleanup, nil
}

// NewDB new database instance
func NewDB(debug bool, dataConf *Database) (*xorm.Engine, error) {
	// The default driver is sqlite3
	if dataConf.Driver == "" {
		dataConf.Driver = string(schemas.SQLITE)
	}
	if dataConf.Driver == string(schemas.SQLITE) {
		dbFileDir := filepath.Dir(dataConf.Connection)
		log.Debugf("try to create database directory %s", dbFileDir)
		if err := dir.CreateDirIfNotExist(dbFileDir); err != nil {
			log.Errorf("create database dir failed: %s", err)
		}
		dataConf.MaxOpenConn = 1
	}
	engine, err := xorm.NewEngine(dataConf.Driver, dataConf.Connection)
	if err != nil {
		return nil, err
	}

	if debug {
		engine.ShowSQL(true)
	} else {
		engine.SetLogLevel(ormlog.LOG_ERR)
	}

	if err = engine.Ping(); err != nil {
		return nil, err
	}

	if dataConf.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(dataConf.MaxIdleConn)
	}
	if dataConf.MaxOpenConn > 0 {
		engine.SetMaxOpenConns(dataConf.MaxOpenConn)
	}
	if dataConf.ConnMaxLifeTime > 0 {
		engine.SetConnMaxLifetime(time.Duration(dataConf.ConnMaxLifeTime) * time.Second)
	}
	engine.SetColumnMapper(names.GonicMapper{})
	return engine, nil
}
