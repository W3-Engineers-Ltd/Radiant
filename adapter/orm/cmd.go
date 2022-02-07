// Copyright 2014 beego Author. All Rights Reserved.
//

package orm

import (
	"github.com/W3-Engineers-Ltd/Radiant/client/orm"
)

// RunCommand listen for orm command and then run it if command arguments passed.
func RunCommand() {
	orm.RunCommand()
}

func RunSyncdb(name string, force bool, verbose bool) error {
	return orm.RunSyncdb(name, force, verbose)
}
