package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationassignment
type Evaluationassignment struct { 
	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// User
	User *User `json:"user,omitempty"`

}

func (o *Evaluationassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationassignment
	
	return json.Marshal(&struct { 
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		User *User `json:"user,omitempty"`
		*Alias
	}{ 
		EvaluationForm: o.EvaluationForm,
		
		User: o.User,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationassignment) UnmarshalJSON(b []byte) error {
	var EvaluationassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationassignmentMap)
	if err != nil {
		return err
	}
	
	if EvaluationForm, ok := EvaluationassignmentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if User, ok := EvaluationassignmentMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
