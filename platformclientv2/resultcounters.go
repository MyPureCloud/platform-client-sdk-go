package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Resultcounters
type Resultcounters struct { 
	// Success
	Success *int `json:"success,omitempty"`


	// Failure
	Failure *int `json:"failure,omitempty"`

}

// String returns a JSON representation of the model
func (o *Resultcounters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
