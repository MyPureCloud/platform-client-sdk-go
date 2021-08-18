package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Calibrationassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Calibrationassignment

	

	return json.Marshal(&struct { 
		Calibrator *User `json:"calibrator,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		ExpertEvaluator *User `json:"expertEvaluator,omitempty"`
		*Alias
	}{ 
		Calibrator: u.Calibrator,
		
		Evaluators: u.Evaluators,
		
		EvaluationForm: u.EvaluationForm,
		
		ExpertEvaluator: u.ExpertEvaluator,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Calibrationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
