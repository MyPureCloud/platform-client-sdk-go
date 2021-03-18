package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowoutcometopicstatsnotification
type Stateventflowoutcometopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventflowoutcometopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventflowoutcometopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
