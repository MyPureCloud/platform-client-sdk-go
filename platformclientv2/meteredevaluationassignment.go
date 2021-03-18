package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Meteredevaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
