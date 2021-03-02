package platformclientv2
import (
	"encoding/json"
)

// Calibrationassignment
type Calibrationassignment struct { 
	// Calibrator
	Calibrator *User `json:"calibrator,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// ExpertEvaluator
	ExpertEvaluator *User `json:"expertEvaluator,omitempty"`

}

// String returns a JSON representation of the model
func (o *Calibrationassignment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
