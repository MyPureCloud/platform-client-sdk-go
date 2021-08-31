package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentbulkaddresponse
type Learningassignmentbulkaddresponse struct { 
	// Entities - The learning assignments that were assigned correctly
	Entities *[]Learningassignment `json:"entities,omitempty"`


	// DisallowedEntities - The items that were not allowed to be assigned
	DisallowedEntities *[]Disallowedentitylearningassignmentitem `json:"disallowedEntities,omitempty"`

}

func (o *Learningassignmentbulkaddresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentbulkaddresponse
	
	return json.Marshal(&struct { 
		Entities *[]Learningassignment `json:"entities,omitempty"`
		
		DisallowedEntities *[]Disallowedentitylearningassignmentitem `json:"disallowedEntities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		DisallowedEntities: o.DisallowedEntities,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentbulkaddresponse) UnmarshalJSON(b []byte) error {
	var LearningassignmentbulkaddresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentbulkaddresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := LearningassignmentbulkaddresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if DisallowedEntities, ok := LearningassignmentbulkaddresponseMap["disallowedEntities"].([]interface{}); ok {
		DisallowedEntitiesString, _ := json.Marshal(DisallowedEntities)
		json.Unmarshal(DisallowedEntitiesString, &o.DisallowedEntities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkaddresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
