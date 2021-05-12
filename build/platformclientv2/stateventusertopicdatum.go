package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventusertopicdatum
type Stateventusertopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventusertopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventusertopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
