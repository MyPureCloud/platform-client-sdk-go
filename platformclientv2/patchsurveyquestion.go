package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchsurveyquestion
type Patchsurveyquestion struct { 
	// VarType - Type of survey question.
	VarType *string `json:"type,omitempty"`


	// Label - Label of question.
	Label *string `json:"label,omitempty"`


	// CustomerProperty - The customer property that the answer maps to.
	CustomerProperty *string `json:"customerProperty,omitempty"`


	// Choices - Choices available to user.
	Choices *[]string `json:"choices,omitempty"`


	// IsMandatory - Whether answering this question is mandatory.
	IsMandatory *bool `json:"isMandatory,omitempty"`

}

func (o *Patchsurveyquestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchsurveyquestion
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		CustomerProperty *string `json:"customerProperty,omitempty"`
		
		Choices *[]string `json:"choices,omitempty"`
		
		IsMandatory *bool `json:"isMandatory,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		Label: o.Label,
		
		CustomerProperty: o.CustomerProperty,
		
		Choices: o.Choices,
		
		IsMandatory: o.IsMandatory,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchsurveyquestion) UnmarshalJSON(b []byte) error {
	var PatchsurveyquestionMap map[string]interface{}
	err := json.Unmarshal(b, &PatchsurveyquestionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := PatchsurveyquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Label, ok := PatchsurveyquestionMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if CustomerProperty, ok := PatchsurveyquestionMap["customerProperty"].(string); ok {
		o.CustomerProperty = &CustomerProperty
	}
    
	if Choices, ok := PatchsurveyquestionMap["choices"].([]interface{}); ok {
		ChoicesString, _ := json.Marshal(Choices)
		json.Unmarshal(ChoicesString, &o.Choices)
	}
	
	if IsMandatory, ok := PatchsurveyquestionMap["isMandatory"].(bool); ok {
		o.IsMandatory = &IsMandatory
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Patchsurveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
