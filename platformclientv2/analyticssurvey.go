package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssurvey
type Analyticssurvey struct { 
	// EventTime - Specifies when an event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// QueueId - The ID of the associated queue
	QueueId *string `json:"queueId,omitempty"`


	// SurveyCompletedDate - Completion datetime of the survey in ISO 8601 format
	SurveyCompletedDate *time.Time `json:"surveyCompletedDate,omitempty"`


	// SurveyFormContextId - Unique identifier for the survey form, regardless of version
	SurveyFormContextId *string `json:"surveyFormContextId,omitempty"`


	// SurveyFormId - ID of the survey form used
	SurveyFormId *string `json:"surveyFormId,omitempty"`


	// SurveyFormName - Name of the survey form used
	SurveyFormName *string `json:"surveyFormName,omitempty"`


	// SurveyId - ID of the survey
	SurveyId *string `json:"surveyId,omitempty"`


	// SurveyPromoterScore - Score of the survey used with NPS
	SurveyPromoterScore *int `json:"surveyPromoterScore,omitempty"`


	// SurveyStatus - The status of the survey
	SurveyStatus *string `json:"surveyStatus,omitempty"`


	// UserId - ID of the agent the survey was performed against
	UserId *string `json:"userId,omitempty"`


	// OSurveyTotalScore
	OSurveyTotalScore *int `json:"oSurveyTotalScore,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticssurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
