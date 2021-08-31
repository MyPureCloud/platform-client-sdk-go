package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeysurveyquestion
type Journeysurveyquestion struct { 
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

func (o *Journeysurveyquestion) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeysurveyquestion
	
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

func (o *Journeysurveyquestion) UnmarshalJSON(b []byte) error {
	var JourneysurveyquestionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneysurveyquestionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := JourneysurveyquestionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Label, ok := JourneysurveyquestionMap["label"].(string); ok {
		o.Label = &Label
	}
	
	if CustomerProperty, ok := JourneysurveyquestionMap["customerProperty"].(string); ok {
		o.CustomerProperty = &CustomerProperty
	}
	
	if Choices, ok := JourneysurveyquestionMap["choices"].([]interface{}); ok {
		ChoicesString, _ := json.Marshal(Choices)
		json.Unmarshal(ChoicesString, &o.Choices)
	}
	
	if IsMandatory, ok := JourneysurveyquestionMap["isMandatory"].(bool); ok {
		o.IsMandatory = &IsMandatory
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeysurveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
