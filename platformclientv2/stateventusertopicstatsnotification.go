package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventusertopicstatsnotification
type Stateventusertopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventusertopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventusertopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
