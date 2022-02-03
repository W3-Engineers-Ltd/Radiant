// Copyright 2015 TiDB Author. All Rights Reserved.
//

package orm

// TiDBQueryBuilder is the SQL build
type TiDBQueryBuilder struct {
	MySQLQueryBuilder
	tokens []string
}
