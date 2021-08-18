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

func (u *Evaluationquestiongroup) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		DefaultAnswersToHighest: u.DefaultAnswersToHighest,
		
		DefaultAnswersToNA: u.DefaultAnswersToNA,
		
		NaEnabled: u.NaEnabled,
		
		Weight: u.Weight,
		
		ManualWeight: u.ManualWeight,
		
		Questions: u.Questions,
		
		VisibilityCondition: u.VisibilityCondition,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evaluationquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
