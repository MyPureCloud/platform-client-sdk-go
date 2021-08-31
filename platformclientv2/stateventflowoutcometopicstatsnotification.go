package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Stateventflowoutcometopicstatsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowoutcometopicstatsnotification
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Stateventflowoutcometopicdatum `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowoutcometopicstatsnotification) UnmarshalJSON(b []byte) error {
	var StateventflowoutcometopicstatsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowoutcometopicstatsnotificationMap)
	if err != nil {
		return err
	}
	
	if Group, ok := StateventflowoutcometopicstatsnotificationMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := StateventflowoutcometopicstatsnotificationMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowoutcometopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
