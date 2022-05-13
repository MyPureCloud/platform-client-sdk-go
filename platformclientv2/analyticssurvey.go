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

func (o *Analyticssurvey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticssurvey
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	SurveyCompletedDate := new(string)
	if o.SurveyCompletedDate != nil {
		
		*SurveyCompletedDate = timeutil.Strftime(o.SurveyCompletedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		QueueId: o.QueueId,
		
		SurveyCompletedDate: SurveyCompletedDate,
		
		SurveyFormContextId: o.SurveyFormContextId,
		
		SurveyFormId: o.SurveyFormId,
		
		SurveyFormName: o.SurveyFormName,
		
		SurveyId: o.SurveyId,
		
		SurveyPromoterScore: o.SurveyPromoterScore,
		
		SurveyStatus: o.SurveyStatus,
		
		UserId: o.UserId,
		
		OSurveyTotalScore: o.OSurveyTotalScore,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticssurvey) UnmarshalJSON(b []byte) error {
	var AnalyticssurveyMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticssurveyMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := AnalyticssurveyMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if QueueId, ok := AnalyticssurveyMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if surveyCompletedDateString, ok := AnalyticssurveyMap["surveyCompletedDate"].(string); ok {
		SurveyCompletedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", surveyCompletedDateString)
		o.SurveyCompletedDate = &SurveyCompletedDate
	}
	
	if SurveyFormContextId, ok := AnalyticssurveyMap["surveyFormContextId"].(string); ok {
		o.SurveyFormContextId = &SurveyFormContextId
	}
    
	if SurveyFormId, ok := AnalyticssurveyMap["surveyFormId"].(string); ok {
		o.SurveyFormId = &SurveyFormId
	}
    
	if SurveyFormName, ok := AnalyticssurveyMap["surveyFormName"].(string); ok {
		o.SurveyFormName = &SurveyFormName
	}
    
	if SurveyId, ok := AnalyticssurveyMap["surveyId"].(string); ok {
		o.SurveyId = &SurveyId
	}
    
	if SurveyPromoterScore, ok := AnalyticssurveyMap["surveyPromoterScore"].(float64); ok {
		SurveyPromoterScoreInt := int(SurveyPromoterScore)
		o.SurveyPromoterScore = &SurveyPromoterScoreInt
	}
	
	if SurveyStatus, ok := AnalyticssurveyMap["surveyStatus"].(string); ok {
		o.SurveyStatus = &SurveyStatus
	}
    
	if UserId, ok := AnalyticssurveyMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if OSurveyTotalScore, ok := AnalyticssurveyMap["oSurveyTotalScore"].(float64); ok {
		OSurveyTotalScoreInt := int(OSurveyTotalScore)
		o.OSurveyTotalScore = &OSurveyTotalScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticssurvey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
