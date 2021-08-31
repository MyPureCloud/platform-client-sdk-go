package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentrulerunnotification
type Learningassignmentruleruntopiclearningassignmentrulerunnotification struct { 
	// Entities
	Entities *[]Learningassignmentruleruntopiclearningassignmentscreated `json:"entities,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`

}

func (o *Learningassignmentruleruntopiclearningassignmentrulerunnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentruleruntopiclearningassignmentrulerunnotification
	
	return json.Marshal(&struct { 
		Entities *[]Learningassignmentruleruntopiclearningassignmentscreated `json:"entities,omitempty"`
		
		Total *int `json:"total,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		Total: o.Total,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentruleruntopiclearningassignmentrulerunnotification) UnmarshalJSON(b []byte) error {
	var LearningassignmentruleruntopiclearningassignmentrulerunnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentruleruntopiclearningassignmentrulerunnotificationMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := LearningassignmentruleruntopiclearningassignmentrulerunnotificationMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Total, ok := LearningassignmentruleruntopiclearningassignmentrulerunnotificationMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentrulerunnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
