package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Meteredassignmentbyagent
type Meteredassignmentbyagent struct { 
	// EvaluationContextId
	EvaluationContextId *string `json:"evaluationContextId,omitempty"`


	// Evaluators
	Evaluators *[]User `json:"evaluators,omitempty"`


	// MaxNumberEvaluations
	MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`


	// EvaluationForm
	EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`


	// TimeInterval
	TimeInterval *Timeinterval `json:"timeInterval,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`

}

func (o *Meteredassignmentbyagent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Meteredassignmentbyagent
	
	return json.Marshal(&struct { 
		EvaluationContextId *string `json:"evaluationContextId,omitempty"`
		
		Evaluators *[]User `json:"evaluators,omitempty"`
		
		MaxNumberEvaluations *int `json:"maxNumberEvaluations,omitempty"`
		
		EvaluationForm *Evaluationform `json:"evaluationForm,omitempty"`
		
		TimeInterval *Timeinterval `json:"timeInterval,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		*Alias
	}{ 
		EvaluationContextId: o.EvaluationContextId,
		
		Evaluators: o.Evaluators,
		
		MaxNumberEvaluations: o.MaxNumberEvaluations,
		
		EvaluationForm: o.EvaluationForm,
		
		TimeInterval: o.TimeInterval,
		
		TimeZone: o.TimeZone,
		Alias:    (*Alias)(o),
	})
}

func (o *Meteredassignmentbyagent) UnmarshalJSON(b []byte) error {
	var MeteredassignmentbyagentMap map[string]interface{}
	err := json.Unmarshal(b, &MeteredassignmentbyagentMap)
	if err != nil {
		return err
	}
	
	if EvaluationContextId, ok := MeteredassignmentbyagentMap["evaluationContextId"].(string); ok {
		o.EvaluationContextId = &EvaluationContextId
	}
	
	if Evaluators, ok := MeteredassignmentbyagentMap["evaluators"].([]interface{}); ok {
		EvaluatorsString, _ := json.Marshal(Evaluators)
		json.Unmarshal(EvaluatorsString, &o.Evaluators)
	}
	
	if MaxNumberEvaluations, ok := MeteredassignmentbyagentMap["maxNumberEvaluations"].(float64); ok {
		MaxNumberEvaluationsInt := int(MaxNumberEvaluations)
		o.MaxNumberEvaluations = &MaxNumberEvaluationsInt
	}
	
	if EvaluationForm, ok := MeteredassignmentbyagentMap["evaluationForm"].(map[string]interface{}); ok {
		EvaluationFormString, _ := json.Marshal(EvaluationForm)
		json.Unmarshal(EvaluationFormString, &o.EvaluationForm)
	}
	
	if TimeInterval, ok := MeteredassignmentbyagentMap["timeInterval"].(map[string]interface{}); ok {
		TimeIntervalString, _ := json.Marshal(TimeInterval)
		json.Unmarshal(TimeIntervalString, &o.TimeInterval)
	}
	
	if TimeZone, ok := MeteredassignmentbyagentMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Meteredassignmentbyagent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
