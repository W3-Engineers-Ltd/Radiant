package orm

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// QueryBuilder is the Query builder interface
type QueryBuilder orm.QueryBuilder

// NewQueryBuilder return the QueryBuilder
func NewQueryBuilder(driver string) (qb QueryBuilder, err error) {
	return orm.NewQueryBuilder(driver)
}
