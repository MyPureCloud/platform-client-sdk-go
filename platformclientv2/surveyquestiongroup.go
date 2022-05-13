package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyquestiongroup
type Surveyquestiongroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// NaEnabled
	NaEnabled *bool `json:"naEnabled,omitempty"`


	// Questions
	Questions *[]Surveyquestion `json:"questions,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`

}

func (o *Surveyquestiongroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyquestiongroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		Questions *[]Surveyquestion `json:"questions,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		NaEnabled: o.NaEnabled,
		
		Questions: o.Questions,
		
		VisibilityCondition: o.VisibilityCondition,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyquestiongroup) UnmarshalJSON(b []byte) error {
	var SurveyquestiongroupMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyquestiongroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SurveyquestiongroupMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SurveyquestiongroupMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := SurveyquestiongroupMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if NaEnabled, ok := SurveyquestiongroupMap["naEnabled"].(bool); ok {
		o.NaEnabled = &NaEnabled
	}
    
	if Questions, ok := SurveyquestiongroupMap["questions"].([]interface{}); ok {
		QuestionsString, _ := json.Marshal(Questions)
		json.Unmarshal(QuestionsString, &o.Questions)
	}
	
	if VisibilityCondition, ok := SurveyquestiongroupMap["visibilityCondition"].(map[string]interface{}); ok {
		VisibilityConditionString, _ := json.Marshal(VisibilityCondition)
		json.Unmarshal(VisibilityConditionString, &o.VisibilityCondition)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
