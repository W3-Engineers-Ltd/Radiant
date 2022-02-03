// Copyright 2020 radiant-dev
//

package orm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRegisterDataBase(t *testing.T) {
	err := RegisterDataBase("test-params", DBARGS.Driver, DBARGS.Source,
		MaxIdleConnections(20),
		MaxOpenConnections(300),
		ConnMaxLifetime(time.Minute))
	assert.Nil(t, err)

	al := getDbAlias("test-params")
	assert.NotNil(t, al)
	assert.Equal(t, al.MaxIdleConns, 20)
	assert.Equal(t, al.MaxOpenConns, 300)
	assert.Equal(t, al.ConnMaxLifetime, time.Minute)
}

func TestRegisterDataBaseMaxStmtCacheSizeNegative1(t *testing.T) {
	aliasName := "TestRegisterDataBase_MaxStmtCacheSizeNegative1"
	err := RegisterDataBase(aliasName, DBARGS.Driver, DBARGS.Source, MaxStmtCacheSize(-1))
	assert.Nil(t, err)

	al := getDbAlias(aliasName)
	assert.NotNil(t, al)
	assert.Equal(t, al.DB.stmtDecoratorsLimit, 0)
}

func TestRegisterDataBaseMaxStmtCacheSize0(t *testing.T) {
	aliasName := "TestRegisterDataBase_MaxStmtCacheSize0"
	err := RegisterDataBase(aliasName, DBARGS.Driver, DBARGS.Source, MaxStmtCacheSize(0))
	assert.Nil(t, err)

	al := getDbAlias(aliasName)
	assert.NotNil(t, al)
	assert.Equal(t, al.DB.stmtDecoratorsLimit, 0)
}

func TestRegisterDataBaseMaxStmtCacheSize1(t *testing.T) {
	aliasName := "TestRegisterDataBase_MaxStmtCacheSize1"
	err := RegisterDataBase(aliasName, DBARGS.Driver, DBARGS.Source, MaxStmtCacheSize(1))
	assert.Nil(t, err)

	al := getDbAlias(aliasName)
	assert.NotNil(t, al)
	assert.Equal(t, al.DB.stmtDecoratorsLimit, 1)
}

func TestRegisterDataBaseMaxStmtCacheSize841(t *testing.T) {
	aliasName := "TestRegisterDataBase_MaxStmtCacheSize841"
	err := RegisterDataBase(aliasName, DBARGS.Driver, DBARGS.Source, MaxStmtCacheSize(841))
	assert.Nil(t, err)

	al := getDbAlias(aliasName)
	assert.NotNil(t, al)
	assert.Equal(t, al.DB.stmtDecoratorsLimit, 841)
}

func TestDBCache(t *testing.T) {
	dataBaseCache.add("test1", &alias{})
	dataBaseCache.add("default", &alias{})
	al := dataBaseCache.getDefault()
	assert.NotNil(t, al)
	al, ok := dataBaseCache.get("test1")
	assert.NotNil(t, al)
	assert.True(t, ok)
}
