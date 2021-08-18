package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Analyticssurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticssurvey

	
	EventTime := new(string)
	if u.EventTime != nil {
		
		*EventTime = timeutil.Strftime(u.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	SurveyCompletedDate := new(string)
	if u.SurveyCompletedDate != nil {
		
		*SurveyCompletedDate = timeutil.Strftime(u.SurveyCompletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		SurveyCompletedDate = nil
	}
	

	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		SurveyCompletedDate *string `json:"surveyCompletedDate,omitempty"`
		
		SurveyFormContextId *string `json:"surveyFormContextId,omitempty"`
		
		SurveyFormId *string `json:"surveyFormId,omitempty"`
		
		SurveyFormName *string `json:"surveyFormName,omitempty"`
		
		SurveyId *string `json:"surveyId,omitempty"`
		
		SurveyPromoterScore *int `json:"surveyPromoterScore,omitempty"`
		
		SurveyStatus *string `json:"surveyStatus,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		OSurveyTotalScore *int `json:"oSurveyTotalScore,omitempty"`
		*Alias
	}{ 
		EventTime: EventTime,
		
		QueueId: u.QueueId,
		
		SurveyCompletedDate: SurveyCompletedDate,
		
		SurveyFormContextId: u.SurveyFormContextId,
		
		SurveyFormId: u.SurveyFormId,
		
		SurveyFormName: u.SurveyFormName,
		
		SurveyId: u.SurveyId,
		
		SurveyPromoterScore: u.SurveyPromoterScore,
		
		SurveyStatus: u.SurveyStatus,
		
		UserId: u.UserId,
		
		OSurveyTotalScore: u.OSurveyTotalScore,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticssurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
