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

func (o *Stateventwrapupcodetopicstatsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventwrapupcodetopicstatsnotification
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Stateventwrapupcodetopicdatum `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventwrapupcodetopicstatsnotification) UnmarshalJSON(b []byte) error {
	var StateventwrapupcodetopicstatsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &StateventwrapupcodetopicstatsnotificationMap)
	if err != nil {
		return err
	}
	
	if Group, ok := StateventwrapupcodetopicstatsnotificationMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := StateventwrapupcodetopicstatsnotificationMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventwrapupcodetopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
