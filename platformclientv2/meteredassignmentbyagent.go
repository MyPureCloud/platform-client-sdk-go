package platformclientv2
import (
	"encoding/json"
)

// Meteredassignmentbyagent
type Meteredassignmentbyagent struct { 
	// EvaluationContextId
	EvaluationContextId *string `json:"evaluationContextId,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// MaxNumberEvaluations
	MaxNumberEvaluations *int32 `json:"maxNumberEvaluations,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// TimeInterval
	TimeInterval *Timeinterval `json:"timeInterval,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`

}

// String returns a JSON representation of the model
func (o *Meteredassignmentbyagent) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
