package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentbulkremoveresponse
type Learningassignmentbulkremoveresponse struct { 
	// Entities - The learning assignments that were removed successfully
	Entities *[]Learningassignmententity `json:"entities,omitempty"`


	// DisallowedEntities - The learning assignments that were not removed due to missing permissions
	DisallowedEntities *[]Disallowedentitylearningassignmentreference `json:"disallowedEntities,omitempty"`

}

func (o *Learningassignmentbulkremoveresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentbulkremoveresponse
	
	return json.Marshal(&struct { 
		Entities *[]Learningassignmententity `json:"entities,omitempty"`
		
		DisallowedEntities *[]Disallowedentitylearningassignmentreference `json:"disallowedEntities,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		DisallowedEntities: o.DisallowedEntities,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentbulkremoveresponse) UnmarshalJSON(b []byte) error {
	var LearningassignmentbulkremoveresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentbulkremoveresponseMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := LearningassignmentbulkremoveresponseMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if DisallowedEntities, ok := LearningassignmentbulkremoveresponseMap["disallowedEntities"].([]interface{}); ok {
		DisallowedEntitiesString, _ := json.Marshal(DisallowedEntities)
		json.Unmarshal(DisallowedEntitiesString, &o.DisallowedEntities)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentbulkremoveresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
