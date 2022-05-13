package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Disallowedentitylearningassignmentitem
type Disallowedentitylearningassignmentitem struct { 
	// ErrorCode - The error code associated with this disallowed entity
	ErrorCode *string `json:"errorCode,omitempty"`


	// Entity - The entity that was disallowed
	Entity *Learningassignmentitem `json:"entity,omitempty"`

}

func (o *Disallowedentitylearningassignmentitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Disallowedentitylearningassignmentitem
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Entity *Learningassignmentitem `json:"entity,omitempty"`
		*Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		Entity: o.Entity,
		Alias:    (*Alias)(o),
	})
}

func (o *Disallowedentitylearningassignmentitem) UnmarshalJSON(b []byte) error {
	var DisallowedentitylearningassignmentitemMap map[string]interface{}
	err := json.Unmarshal(b, &DisallowedentitylearningassignmentitemMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := DisallowedentitylearningassignmentitemMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if Entity, ok := DisallowedentitylearningassignmentitemMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Disallowedentitylearningassignmentitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
