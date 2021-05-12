package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventqueuetopicstatsnotification
type Stateventqueuetopicstatsnotification struct { 
	// Group
	Group *map[string]string `json:"group,omitempty"`


	// Data
	Data *[]Stateventqueuetopicdatum `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Stateventqueuetopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
