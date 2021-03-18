package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
