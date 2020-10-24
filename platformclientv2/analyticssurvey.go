package platformclientv2
import (
	"time"
	"encoding/json"
)

// Analyticssurvey
type Analyticssurvey struct { 
	// SurveyId - Unique identifier for the survey
	SurveyId *string `json:"surveyId,omitempty"`


	// SurveyFormId - Unique identifier for the survey form
	SurveyFormId *string `json:"surveyFormId,omitempty"`


	// SurveyFormName - Name of the survey form
	SurveyFormName *string `json:"surveyFormName,omitempty"`


	// SurveyFormContextId - Unique identifier for the survey form, regardless of version
	SurveyFormContextId *string `json:"surveyFormContextId,omitempty"`


	// EventTime - Specifies when a survey occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// UserId - A unique identifier of the PureCloud user
	UserId *string `json:"userId,omitempty"`


	// QueueId - Unique identifier for the queue the conversation was on
	QueueId *string `json:"queueId,omitempty"`


	// SurveyStatus - Survey status
	SurveyStatus *string `json:"surveyStatus,omitempty"`


	// SurveyPromoterScore - Promoter score of the survey
	SurveyPromoterScore *int32 `json:"surveyPromoterScore,omitempty"`


	// SurveyCompletedDate - Completion date/time of the survey. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SurveyCompletedDate *time.Time `json:"surveyCompletedDate,omitempty"`


	// OSurveyTotalScore
	OSurveyTotalScore *int64 `json:"oSurveyTotalScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticssurvey) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
