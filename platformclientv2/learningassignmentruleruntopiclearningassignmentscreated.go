package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningassignmentscreated
type Learningassignmentruleruntopiclearningassignmentscreated struct { 
	// Module
	Module *Learningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`

}

func (o *Learningassignmentruleruntopiclearningassignmentscreated) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentruleruntopiclearningassignmentscreated
	
	return json.Marshal(&struct { 
		Module *Learningassignmentruleruntopiclearningmodulereference `json:"module,omitempty"`
		*Alias
	}{ 
		Module: o.Module,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentruleruntopiclearningassignmentscreated) UnmarshalJSON(b []byte) error {
	var LearningassignmentruleruntopiclearningassignmentscreatedMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentruleruntopiclearningassignmentscreatedMap)
	if err != nil {
		return err
	}
	
	if Module, ok := LearningassignmentruleruntopiclearningassignmentscreatedMap["module"].(map[string]interface{}); ok {
		ModuleString, _ := json.Marshal(Module)
		json.Unmarshal(ModuleString, &o.Module)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningassignmentscreated) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
