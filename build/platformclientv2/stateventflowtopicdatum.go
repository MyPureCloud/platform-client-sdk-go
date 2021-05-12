package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowtopicdatum
type Stateventflowtopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventflowtopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventflowtopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
