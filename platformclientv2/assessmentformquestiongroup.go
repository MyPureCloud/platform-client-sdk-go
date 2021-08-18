package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentformquestiongroup
type Assessmentformquestiongroup struct { 
	// Id - The ID of the question group,
	Id *string `json:"id,omitempty"`


	// Name - The question group name
	Name *string `json:"name,omitempty"`


	// VarType - The question group type
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


	// Questions - The list of questions for this question group
	Questions *[]Assessmentformquestion `json:"questions,omitempty"`


	// VisibilityCondition
	VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Assessmentformquestiongroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentformquestiongroup

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		DefaultAnswersToHighest *bool `json:"defaultAnswersToHighest,omitempty"`
		
		DefaultAnswersToNA *bool `json:"defaultAnswersToNA,omitempty"`
		
		NaEnabled *bool `json:"naEnabled,omitempty"`
		
		Weight *float32 `json:"weight,omitempty"`
		
		ManualWeight *bool `json:"manualWeight,omitempty"`
		
		Questions *[]Assessmentformquestion `json:"questions,omitempty"`
		
		VisibilityCondition *Visibilitycondition `json:"visibilityCondition,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
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
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assessmentformquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
