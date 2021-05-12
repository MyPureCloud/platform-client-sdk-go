package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventwrapupcodetopicstatsnotification
type Stateventwrapupcodetopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventwrapupcodetopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
