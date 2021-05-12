package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventwrapupcodetopicdatum
type Stateventwrapupcodetopicdatum struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventwrapupcodetopicmetric `json:"metrics,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicdatum) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
