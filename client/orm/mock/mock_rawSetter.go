// Copyright 2020 radiant
//

package mock

import (
	"database/sql"

	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

type DoNothingRawSetter struct{}

func (d *DoNothingRawSetter) Exec() (sql.Result, error) {
	return nil, nil
}

func (d *DoNothingRawSetter) QueryRow(containers ...interface{}) error {
	return nil
}

func (d *DoNothingRawSetter) QueryRows(containers ...interface{}) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) SetArgs(i ...interface{}) orm.RawSeter {
	return d
}

func (d *DoNothingRawSetter) Values(container *[]orm.Params, cols ...string) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) ValuesList(container *[]orm.ParamsList, cols ...string) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) ValuesFlat(container *orm.ParamsList, cols ...string) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) RowsToMap(result *orm.Params, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) RowsToStruct(ptrStruct interface{}, keyCol, valueCol string) (int64, error) {
	return 0, nil
}

func (d *DoNothingRawSetter) Prepare() (orm.RawPreparer, error) {
	return nil, nil
}
