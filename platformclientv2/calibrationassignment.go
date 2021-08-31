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

func (o *Calibrationassignment) MarshalJSON() ([]byte, error) {
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
		Calibrator: o.Calibrator,
		
		Evaluators: o.Evaluators,
		
		EvaluationForm: o.EvaluationForm,
		
		ExpertEvaluator: o.ExpertEvaluator,
		Alias:    (*Alias)(o),
	})
}

func (o *Calibrationassignment) UnmarshalJSON(b []byte) error {
	var CalibrationassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &CalibrationassignmentMap)
	if err != nil {
		return err
	}
	
	if Calibrator, ok := CalibrationassignmentMap["calibrator"].(map[string]interface{}); ok {
		CalibratorString, _ := json.Marshal(Calibrator)
		json.Unmarshal(CalibratorString, &o.Calibrator)
	}
	
	if Evaluators, ok := CalibrationassignmentMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if EvaluationForm, ok := CalibrationassignmentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if ExpertEvaluator, ok := CalibrationassignmentMap["expertEvaluator"].(map[string]interface{}); ok {
		ExpertEvaluatorString, _ := json.Marshal(ExpertEvaluator)
		json.Unmarshal(ExpertEvaluatorString, &o.ExpertEvaluator)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Calibrationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
