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


	// CoverArt - The cover art for the learning module
	CoverArt *Learningmodulecoverartrequest `json:"coverArt,omitempty"`

}

func (o *Learningmodulerequest) MarshalJSON() ([]byte, error) {
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
		
		CoverArt *Learningmodulecoverartrequest `json:"coverArt,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		CompletionTimeInDays: o.CompletionTimeInDays,
		
		InformSteps: o.InformSteps,
		
		VarType: o.VarType,
		
		AssessmentForm: o.AssessmentForm,
		
		CoverArt: o.CoverArt,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningmodulerequest) UnmarshalJSON(b []byte) error {
	var LearningmodulerequestMap map[string]interface{}
	err := json.Unmarshal(b, &LearningmodulerequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := LearningmodulerequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := LearningmodulerequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CompletionTimeInDays, ok := LearningmodulerequestMap["completionTimeInDays"].(float64); ok {
		CompletionTimeInDaysInt := int(CompletionTimeInDays)
		o.CompletionTimeInDays = &CompletionTimeInDaysInt
	}
	
	if InformSteps, ok := LearningmodulerequestMap["informSteps"].([]interface{}); ok {
		InformStepsString, _ := json.Marshal(InformSteps)
		json.Unmarshal(InformStepsString, &o.InformSteps)
	}
	
	if VarType, ok := LearningmodulerequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if AssessmentForm, ok := LearningmodulerequestMap["assessmentForm"].(map[string]interface{}); ok {
		AssessmentFormString, _ := json.Marshal(AssessmentForm)
		json.Unmarshal(AssessmentFormString, &o.AssessmentForm)
	}
	
	if CoverArt, ok := LearningmodulerequestMap["coverArt"].(map[string]interface{}); ok {
		CoverArtString, _ := json.Marshal(CoverArt)
		json.Unmarshal(CoverArtString, &o.CoverArt)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningmodulerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
