package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassessment
type Learningassessment struct { 
	// AssessmentId - The Id of the assessment
	AssessmentId *string `json:"assessmentId,omitempty"`


	// ContextId - The context Id of the related assessment form
	ContextId *string `json:"contextId,omitempty"`


	// AssessmentFormId - The Id of the related assessment form
	AssessmentFormId *string `json:"assessmentFormId,omitempty"`


	// Status - Status of the assessment
	Status *string `json:"status,omitempty"`


	// Answers - Answers for the assessment
	Answers *Assessmentscoringset `json:"answers,omitempty"`


	// DateCreated - Date the assessment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date the assessment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// DateSubmitted - Date the assessment was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSubmitted *time.Time `json:"dateSubmitted,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
