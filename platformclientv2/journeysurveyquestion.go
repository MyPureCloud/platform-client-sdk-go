package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Journeysurveyquestion) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
