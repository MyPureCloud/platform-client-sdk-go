package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Learningassessment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassessment

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateSubmitted := new(string)
	if u.DateSubmitted != nil {
		
		*DateSubmitted = timeutil.Strftime(u.DateSubmitted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSubmitted = nil
	}
	

	return json.Marshal(&struct { 
		AssessmentId *string `json:"assessmentId,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		AssessmentFormId *string `json:"assessmentFormId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Answers *Assessmentscoringset `json:"answers,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateSubmitted *string `json:"dateSubmitted,omitempty"`
		*Alias
	}{ 
		AssessmentId: u.AssessmentId,
		
		ContextId: u.ContextId,
		
		AssessmentFormId: u.AssessmentFormId,
		
		Status: u.Status,
		
		Answers: u.Answers,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateSubmitted: DateSubmitted,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
