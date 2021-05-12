package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Assessmentform
type Assessmentform struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// DateModified - Last modified date of the assessment form. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContextId - The unique Id for all versions of this assessment form
	ContextId *string `json:"contextId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// Published - If true, assessment form is published
	Published *bool `json:"published,omitempty"`


	// PassPercent - The pass percent for the assessment form
	PassPercent *int `json:"passPercent,omitempty"`


	// QuestionGroups - A list of question groups
	QuestionGroups *[]Assessmentformquestiongroup `json:"questionGroups,omitempty"`

}

// String returns a JSON representation of the model
func (o *Assessmentform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
