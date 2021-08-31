package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Stateventflowtopicstatsnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventflowtopicstatsnotification
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Stateventflowtopicdatum `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventflowtopicstatsnotification) UnmarshalJSON(b []byte) error {
	var StateventflowtopicstatsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &StateventflowtopicstatsnotificationMap)
	if err != nil {
		return err
	}
	
	if Group, ok := StateventflowtopicstatsnotificationMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := StateventflowtopicstatsnotificationMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventflowtopicstatsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
