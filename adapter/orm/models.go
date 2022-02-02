package orm

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// ResetModelCache Clean model cache. Then you can re-RegisterModel.
// Common use this api for test case.
func ResetModelCache() {
	orm.ResetModelCache()
}
