package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Assessmentjoblisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Assessmentjoblisting

	

	return json.Marshal(&struct { 
		Total *int `json:"total,omitempty"`
		
		Entities *[]Benefitassessmentjob `json:"entities,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Total: u.Total,
		
		Entities: u.Entities,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Assessmentjoblisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
