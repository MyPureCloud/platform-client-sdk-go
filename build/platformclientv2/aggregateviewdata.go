package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregateviewdata
type Aggregateviewdata struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

// String returns a JSON representation of the model
func (o *Aggregateviewdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
