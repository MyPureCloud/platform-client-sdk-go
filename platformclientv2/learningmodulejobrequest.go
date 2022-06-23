package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulejobrequest - Learning module job request
type Learningmodulejobrequest struct { 
	// Action - The type for the learning module job
	Action *string `json:"action,omitempty"`

}

func (o *Learningmodulejobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulejobrequest
	
	return json.Marshal(&struct { 
		Action *string `json:"action,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulejobrequest) UnmarshalJSON(b []byte) error {
	var LearningmodulejobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulejobrequestMap)
	if err != nil {
		return err
	}
	
	if Action, ok := LearningmodulejobrequestMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulejobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
