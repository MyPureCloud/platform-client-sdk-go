package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulerequest - Learning module request
type Learningmodulerequest struct { 
	// Name - The name of learning module
	Name *string `json:"name,omitempty"`


	// Description - The description of learning module
	Description *string `json:"description,omitempty"`


	// CompletionTimeInDays - The completion time of learning module in days
	CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`


	// InformSteps - The list of inform steps in a learning module
	InformSteps *[]Learningmoduleinformsteprequest `json:"informSteps,omitempty"`


	// VarType - The type for the learning module
	VarType *string `json:"type,omitempty"`


	// AssessmentForm - The assessment form for learning module
	AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`

}

func (u *Learningmodulerequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulerequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CompletionTimeInDays *int `json:"completionTimeInDays,omitempty"`
		
		InformSteps *[]Learningmoduleinformsteprequest `json:"informSteps,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		AssessmentForm *Assessmentform `json:"assessmentForm,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Description: u.Description,
		
		CompletionTimeInDays: u.CompletionTimeInDays,
		
		InformSteps: u.InformSteps,
		
		VarType: u.VarType,
		
		AssessmentForm: u.AssessmentForm,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningmodulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
