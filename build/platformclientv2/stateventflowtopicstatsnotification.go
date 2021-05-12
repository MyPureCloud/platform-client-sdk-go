package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventflowtopicstatsnotification
type Stateventflowtopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventflowtopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventflowtopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
