//
// Copyright (c) 2018 Tencent
//
// SPDX-License-Identifier: Apache-2.0
//

package client

type ConfigurationStruct struct {
	Hostname             string
	Port                 int
	DBType               string
	MongoURL             string
	MongoUsername        string
	MongoPassword        string
	MongoDatabase        string
	MongoPort            int
	MongoConnectTimeout  int
	MongoSocketTimeout   int
	ConsulHost           string
	ConsulPort           int
	CheckInterval        string
	ConsulProfilesActive string
	DistroHost           string
	DistroPort           int
}

var configuration ConfigurationStruct = ConfigurationStruct{} // Needs to be initialized before used

