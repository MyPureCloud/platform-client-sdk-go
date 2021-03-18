package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventqueuetopicmetric
type Stateventqueuetopicmetric struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *map[string]float32 `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
