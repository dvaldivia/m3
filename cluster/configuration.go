// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2019 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cluster

import (
	"errors"

	"github.com/minio/m3/cluster/db"
)

type Configuration struct {
	Key       string
	Value     interface{}
	ValueType string
	locked    bool
}

func (c *Configuration) ValString() (*string, error) {
	if c.ValueType == "string" {
		val := c.Value.(string)
		return &val, nil
	}
	return nil, errors.New("Invalid value type")
}

func (c *Configuration) ValBool() bool {
	if c.ValueType == "bool" {
		if c.Value == "true" {
			return true
		}
	}
	return false
}

func SetConfig(ctx *Context, key, val, valType string) error {
	return SetConfigWithLock(ctx, key, val, valType, false)
}

func SetConfigWithLock(ctx *Context, key, val, valType string, locked bool) error {
	// insert the new configuration
	query :=
		`INSERT INTO
				configurations ("key", "value", "type", "locked", "sys_created_by")
			  VALUES
				($1, $2, $3, $4, $5)`
	// If we were provided context, query inside a transaction ¬if ctx != nil {
	if _, err := ctx.MainTx().Exec(query, key, val, valType, locked, ctx.WhoAmI); err != nil {
		return err
	}
	return nil
}

func GetConfig(key string) (*Configuration, error) {
	query :=
		`SELECT 
				c.key, c.value, c.type, c.locked
			FROM 
				configurations c
			WHERE c.key=$1`
	// non-transactional query
	rows, err := db.GetInstance().MainDB().Query(query, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// if we iterate at least once, we found a result
	for rows.Next() {
		config := Configuration{}
		// Save the resulted query on the User struct
		err := rows.Scan(&config.Key, &config.Value, &config.ValueType, &config.locked)
		if err != nil {
			return nil, err
		}
		return &config, nil
	}
	return nil, err
}
