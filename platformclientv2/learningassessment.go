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

func (o *Learningassessment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassessment
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateSubmitted := new(string)
	if o.DateSubmitted != nil {
		
		*DateSubmitted = timeutil.Strftime(o.DateSubmitted, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		AssessmentId: o.AssessmentId,
		
		ContextId: o.ContextId,
		
		AssessmentFormId: o.AssessmentFormId,
		
		Status: o.Status,
		
		Answers: o.Answers,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateSubmitted: DateSubmitted,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassessment) UnmarshalJSON(b []byte) error {
	var LearningassessmentMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassessmentMap)
	if err != nil {
		return err
	}
	
	if AssessmentId, ok := LearningassessmentMap["assessmentId"].(string); ok {
		o.AssessmentId = &AssessmentId
	}
	
	if ContextId, ok := LearningassessmentMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
	
	if AssessmentFormId, ok := LearningassessmentMap["assessmentFormId"].(string); ok {
		o.AssessmentFormId = &AssessmentFormId
	}
	
	if Status, ok := LearningassessmentMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Answers, ok := LearningassessmentMap["answers"].(map[string]interface{}); ok {
		AnswersString, _ := json.Marshal(Answers)
		json.Unmarshal(AnswersString, &o.Answers)
	}
	
	if dateCreatedString, ok := LearningassessmentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := LearningassessmentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateSubmittedString, ok := LearningassessmentMap["dateSubmitted"].(string); ok {
		DateSubmitted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSubmittedString)
		o.DateSubmitted = &DateSubmitted
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassessment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
