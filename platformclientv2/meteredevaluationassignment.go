package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Meteredevaluationassignment
type Meteredevaluationassignment struct { 
	// EvaluationContextId
	EvaluationContextId *string `json:"evaluationContextId,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// MaxNumberEvaluations
	MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// AssignToActiveUser
	AssignToActiveUser *bool `json:"assignToActiveUser,omitempty"`


	// TimeInterval
	TimeInterval *Timeinterval `json:"timeInterval,omitempty"`

}

func (u *Meteredevaluationassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Meteredevaluationassignment

	

	return json.Marshal(&struct { 
		EvaluationContextId *string `json:"evaluationContextId,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		AssignToActiveUser *bool `json:"assignToActiveUser,omitempty"`
		
		TimeInterval *Timeinterval `json:"timeInterval,omitempty"`
		*Alias
	}{ 
		EvaluationContextId: u.EvaluationContextId,
		
		Evaluators: u.Evaluators,
		
		MaxNumberEvaluations: u.MaxNumberEvaluations,
		
		EvaluationForm: u.EvaluationForm,
		
		AssignToActiveUser: u.AssignToActiveUser,
		
		TimeInterval: u.TimeInterval,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Meteredevaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
