package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregatemetricdata
type Aggregatemetricdata struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregatemetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
