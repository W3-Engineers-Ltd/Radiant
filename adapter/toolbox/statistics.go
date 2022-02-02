package toolbox

import (
	"time"

	"github.com/W3-Engineers-Ltd/Radiant/server/web"
)

// Statistics struct
type Statistics web.Statistics

// URLMap contains several statistics struct to log different data
type URLMap web.URLMap

// AddStatistics add statistics task.
// it needs request method, request url, request controller and statistics time duration
func (m *URLMap) AddStatistics(requestMethod, requestURL, requestController string, requesttime time.Duration) {
	(*web.URLMap)(m).AddStatistics(requestMethod, requestURL, requestController, requesttime)
}

// GetMap put url statistics result in io.Writer
func (m *URLMap) GetMap() map[string]interface{} {
	return (*web.URLMap)(m).GetMap()
}

// GetMapData return all mapdata
func (m *URLMap) GetMapData() []map[string]interface{} {
	return (*web.URLMap)(m).GetMapData()
}

// StatisticsMap hosld global statistics data map
var StatisticsMap *URLMap

func init() {
	StatisticsMap = (*URLMap)(web.StatisticsMap)
}
