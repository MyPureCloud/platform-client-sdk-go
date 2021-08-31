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

func (o *Meteredevaluationassignment) MarshalJSON() ([]byte, error) {
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
		EvaluationContextId: o.EvaluationContextId,
		
		Evaluators: o.Evaluators,
		
		MaxNumberEvaluations: o.MaxNumberEvaluations,
		
		EvaluationForm: o.EvaluationForm,
		
		AssignToActiveUser: o.AssignToActiveUser,
		
		TimeInterval: o.TimeInterval,
		Alias:    (*Alias)(o),
	})
}

func (o *Meteredevaluationassignment) UnmarshalJSON(b []byte) error {
	var MeteredevaluationassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &MeteredevaluationassignmentMap)
	if err != nil {
		return err
	}
	
	if EvaluationContextId, ok := MeteredevaluationassignmentMap["evaluationContextId"].(string); ok {
		o.EvaluationContextId = &EvaluationContextId
	}
	
	if Evaluators, ok := MeteredevaluationassignmentMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if MaxNumberEvaluations, ok := MeteredevaluationassignmentMap["maxNumberEvaluations"].(float64); ok {
		MaxNumberEvaluationsInt := int(MaxNumberEvaluations)
		o.MaxNumberEvaluations = &MaxNumberEvaluationsInt
	}
	
	if EvaluationForm, ok := MeteredevaluationassignmentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if AssignToActiveUser, ok := MeteredevaluationassignmentMap["assignToActiveUser"].(bool); ok {
		o.AssignToActiveUser = &AssignToActiveUser
	}
	
	if TimeInterval, ok := MeteredevaluationassignmentMap["timeInterval"].(map[string]interface{}); ok {
		TimeIntervalString, _ := json.Marshal(TimeInterval)
		json.Unmarshal(TimeIntervalString, &o.TimeInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Meteredevaluationassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
