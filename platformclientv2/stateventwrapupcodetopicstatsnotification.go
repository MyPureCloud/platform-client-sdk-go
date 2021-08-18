package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Stateventwrapupcodetopicstatsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventwrapupcodetopicstatsnotification

	

	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Stateventwrapupcodetopicdatum `json:"data,omitempty"`
		*Alias
	}{ 
		Group: u.Group,
		
		Data: u.Data,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
