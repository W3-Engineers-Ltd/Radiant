// Copyright 2014 beego Author. All Rights Reserved.
//

package toolbox

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStatics(t *testing.T) {
	userApi := "/api/user"
	post := "POST"
	adminUser := "&admin.user"
	StatisticsMap.AddStatistics(post, userApi, adminUser, time.Duration(2000))
	StatisticsMap.AddStatistics(post, userApi, adminUser, time.Duration(120000))
	StatisticsMap.AddStatistics("GET", userApi, adminUser, time.Duration(13000))
	StatisticsMap.AddStatistics(post, "/api/admin", adminUser, time.Duration(14000))
	StatisticsMap.AddStatistics(post, "/api/user/astaxie", adminUser, time.Duration(12000))
	StatisticsMap.AddStatistics(post, "/api/user/xiemengjun", adminUser, time.Duration(13000))
	StatisticsMap.AddStatistics("DELETE", userApi, adminUser, time.Duration(1400))
	t.Log(StatisticsMap.GetMap())

	data := StatisticsMap.GetMapData()
	b, err := json.Marshal(data)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Log(string(b))
}
