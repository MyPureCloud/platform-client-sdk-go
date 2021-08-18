package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregatequeryresponsegroupeddata
type Learningassignmentaggregatequeryresponsegroupeddata struct { 
	// Group - The group values for this data
	Group *map[string]string `json:"group,omitempty"`


	// Data - The metrics in this group
	Data *[]Learningassignmentaggregatequeryresponsedata `json:"data,omitempty"`

}

func (u *Learningassignmentaggregatequeryresponsegroupeddata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryresponsegroupeddata

	

	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Learningassignmentaggregatequeryresponsedata `json:"data,omitempty"`
		*Alias
	}{ 
		Group: u.Group,
		
		Data: u.Data,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsegroupeddata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
