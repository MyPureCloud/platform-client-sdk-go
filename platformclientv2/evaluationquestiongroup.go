package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationquestiongroup
type Evaluationquestiongroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// DefaultAnswersToHighest
	DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`


	// DefaultAnswersToNA
	DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// Weight
	Weight *float32 `json:"weight,omitempty"`


	// ManualWeight
	ManualWeight *bool `json:"manualWeight,omitempty"`


	// Questions
	Questions *[]Evaluationquestion `json:"questions,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`

}

func (o *Evaluationquestiongroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationquestiongroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`
		
		DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		Weight *float32 `json:"weight,omitempty"`
		
		ManualWeight *bool `json:"manualWeight,omitempty"`
		
		Questions *[]Evaluationquestion `json:"questions,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		DefaultAnswersToHighest: o.DefaultAnswersToHighest,
		
		DefaultAnswersToNA: o.DefaultAnswersToNA,
		
		NaEnabled: o.NaEnabled,
		
		Weight: o.Weight,
		
		ManualWeight: o.ManualWeight,
		
		Questions: o.Questions,
		
		VisibilityCondition: o.VisibilityCondition,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationquestiongroup) UnmarshalJSON(b []byte) error {
	var EvaluationquestiongroupMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationquestiongroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EvaluationquestiongroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EvaluationquestiongroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := EvaluationquestiongroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if DefaultAnswersToHighest, ok := EvaluationquestiongroupMap["defaultAnswersToHighest"].(bool); ok {
		o.DefaultAnswersToHighest = &DefaultAnswersToHighest
	}
    
	if DefaultAnswersToNA, ok := EvaluationquestiongroupMap["defaultAnswersToNA"].(bool); ok {
		o.DefaultAnswersToNA = &DefaultAnswersToNA
	}
    
	if NaEnabled, ok := EvaluationquestiongroupMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
    
	if Weight, ok := EvaluationquestiongroupMap["weight"].(float64); ok {
		WeightFloat32 := float32(Weight)
		o.Weight = &WeightFloat32
	}
	
	if ManualWeight, ok := EvaluationquestiongroupMap["manualWeight"].(bool); ok {
		o.ManualWeight = &ManualWeight
	}
    
	if Questions, ok := EvaluationquestiongroupMap["questions"].([]interface{}); ok {
		QuestionsString, _ := json.Marshal(Questions)
		json.Unmarshal(QuestionsString, &o.Questions)
	}
	
	if VisibilityCondition, ok := EvaluationquestiongroupMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
