package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentjoblisting
type Assessmentjoblisting struct { 
	// Total
	Total *int `json:"total,omitempty"`


	// Entities
	Entities *[]Benefitassessmentjob `json:"entities,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentjoblisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
