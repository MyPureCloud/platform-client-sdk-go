package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentruleruntopiclearningmodulereference
type Learningassignmentruleruntopiclearningmodulereference struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`

}

func (o *Learningassignmentruleruntopiclearningmodulereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentruleruntopiclearningmodulereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentruleruntopiclearningmodulereference) UnmarshalJSON(b []byte) error {
	var LearningassignmentruleruntopiclearningmodulereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentruleruntopiclearningmodulereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LearningassignmentruleruntopiclearningmodulereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LearningassignmentruleruntopiclearningmodulereferenceMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentruleruntopiclearningmodulereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
