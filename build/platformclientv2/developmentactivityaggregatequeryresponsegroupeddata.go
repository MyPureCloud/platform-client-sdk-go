package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivityaggregatequeryresponsegroupeddata
type Developmentactivityaggregatequeryresponsegroupeddata struct { 
	// Group - The group values for this data
	Group *map[string]string `json:"group,omitempty"`


	// Data - The metrics in this group
	Data *[]Developmentactivityaggregatequeryresponsedata `json:"data,omitempty"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsegroupeddata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
