package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentitem
type Learningassignmentitem struct { 
	// ModuleId - The Learning Module ID associated with this assignment
	ModuleId *string `json:"moduleId,omitempty"`


	// UserId - The User ID associated with this assignment
	UserId *string `json:"userId,omitempty"`

}

func (o *Learningassignmentitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentitem
	
	return json.Marshal(&struct { 
		ModuleId *string `json:"moduleId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		*Alias
	}{ 
		ModuleId: o.ModuleId,
		
		UserId: o.UserId,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentitem) UnmarshalJSON(b []byte) error {
	var LearningassignmentitemMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentitemMap)
	if err != nil {
		return err
	}
	
	if ModuleId, ok := LearningassignmentitemMap["moduleId"].(string); ok {
		o.ModuleId = &ModuleId
	}
	
	if UserId, ok := LearningassignmentitemMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
