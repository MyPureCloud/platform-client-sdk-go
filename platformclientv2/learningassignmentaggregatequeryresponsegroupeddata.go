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

func (o *Learningassignmentaggregatequeryresponsegroupeddata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregatequeryresponsegroupeddata
	
	return json.Marshal(&struct { 
		Group *map[string]string `json:"group,omitempty"`
		
		Data *[]Learningassignmentaggregatequeryresponsedata `json:"data,omitempty"`
		*Alias
	}{ 
		Group: o.Group,
		
		Data: o.Data,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregatequeryresponsegroupeddata) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregatequeryresponsegroupeddataMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregatequeryresponsegroupeddataMap)
	if err != nil {
		return err
	}
	
	if Group, ok := LearningassignmentaggregatequeryresponsegroupeddataMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Data, ok := LearningassignmentaggregatequeryresponsegroupeddataMap["data"].([]interface{}); ok {
		DataString, _ := json.Marshal(Data)
		json.Unmarshal(DataString, &o.Data)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsegroupeddata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
