package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Disallowedentitylearningassignmentreference
type Disallowedentitylearningassignmentreference struct { 
	// ErrorCode - The error code associated with this disallowed entity
	ErrorCode *string `json:"errorCode,omitempty"`


	// Entity - The entity that was disallowed
	Entity *Learningassignmentreference `json:"entity,omitempty"`

}

func (o *Disallowedentitylearningassignmentreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Disallowedentitylearningassignmentreference
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Entity *Learningassignmentreference `json:"entity,omitempty"`
		*Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		Entity: o.Entity,
		Alias:    (*Alias)(o),
	})
}

func (o *Disallowedentitylearningassignmentreference) UnmarshalJSON(b []byte) error {
	var DisallowedentitylearningassignmentreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &DisallowedentitylearningassignmentreferenceMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := DisallowedentitylearningassignmentreferenceMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if Entity, ok := DisallowedentitylearningassignmentreferenceMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
