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

func (u *Surveyquestiongroup) MarshalJSON() ([]byte, error) {
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
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		NaEnabled: u.NaEnabled,
		
		Questions: u.Questions,
		
		VisibilityCondition: u.VisibilityCondition,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Surveyquestiongroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
