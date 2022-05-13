package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentupdate
type Learningassignmentupdate struct { 
	// State - The Learning Assignment state
	State *string `json:"state,omitempty"`


	// Assessment - An updated Assessment
	Assessment *Learningassessment `json:"assessment,omitempty"`

}

func (o *Learningassignmentupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentupdate
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		Assessment *Learningassessment `json:"assessment,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Assessment: o.Assessment,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentupdate) UnmarshalJSON(b []byte) error {
	var LearningassignmentupdateMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentupdateMap)
	if err != nil {
		return err
	}
	
	if State, ok := LearningassignmentupdateMap["state"].(string); ok {
		o.State = &State
	}
    
	if Assessment, ok := LearningassignmentupdateMap["assessment"].(map[string]interface{}); ok {
		AssessmentString, _ := json.Marshal(Assessment)
		json.Unmarshal(AssessmentString, &o.Assessment)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
