package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsevaluation
type Analyticsevaluation struct { 
	// CalibrationId - The calibration ID used for the purpose of training evaluators
	CalibrationId *string `json:"calibrationId,omitempty"`


	// ContextId - A unique identifier for an evaluation form, regardless of version
	ContextId *string `json:"contextId,omitempty"`


	// Deleted - Whether the evaluation has been deleted
	Deleted *bool `json:"deleted,omitempty"`


	// EvaluationId - Unique identifier for the evaluation
	EvaluationId *string `json:"evaluationId,omitempty"`


	// EvaluatorId - A unique identifier of the user who evaluated the interaction
	EvaluatorId *string `json:"evaluatorId,omitempty"`


	// EventTime - Specifies when an evaluation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventTime *time.Time `json:"eventTime,omitempty"`


	// FormId - ID of the evaluation form used
	FormId *string `json:"formId,omitempty"`


	// FormName - Name of the evaluation form used
	FormName *string `json:"formName,omitempty"`


	// QueueId - The ID of the associated queue
	QueueId *string `json:"queueId,omitempty"`


	// Released - Whether the evaluation has been released
	Released *bool `json:"released,omitempty"`


	// Rescored - Whether the evaluation has been rescored at least once
	Rescored *bool `json:"rescored,omitempty"`


	// UserId - ID of the agent the evaluation was performed against
	UserId *string `json:"userId,omitempty"`


	// OTotalCriticalScore
	OTotalCriticalScore *int `json:"oTotalCriticalScore,omitempty"`


	// OTotalScore
	OTotalScore *int `json:"oTotalScore,omitempty"`

}

func (o *Analyticsevaluation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsevaluation
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		CalibrationId *string `json:"calibrationId,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		EvaluationId *string `json:"evaluationId,omitempty"`
		
		EvaluatorId *string `json:"evaluatorId,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		FormId *string `json:"formId,omitempty"`
		
		FormName *string `json:"formName,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		Released *bool `json:"released,omitempty"`
		
		Rescored *bool `json:"rescored,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		OTotalCriticalScore *int `json:"oTotalCriticalScore,omitempty"`
		
		OTotalScore *int `json:"oTotalScore,omitempty"`
		*Alias
	}{ 
		CalibrationId: o.CalibrationId,
		
		ContextId: o.ContextId,
		
		Deleted: o.Deleted,
		
		EvaluationId: o.EvaluationId,
		
		EvaluatorId: o.EvaluatorId,
		
		EventTime: EventTime,
		
		FormId: o.FormId,
		
		FormName: o.FormName,
		
		QueueId: o.QueueId,
		
		Released: o.Released,
		
		Rescored: o.Rescored,
		
		UserId: o.UserId,
		
		OTotalCriticalScore: o.OTotalCriticalScore,
		
		OTotalScore: o.OTotalScore,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsevaluation) UnmarshalJSON(b []byte) error {
	var AnalyticsevaluationMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsevaluationMap)
	if err != nil {
		return err
	}
	
	if CalibrationId, ok := AnalyticsevaluationMap["calibrationId"].(string); ok {
		o.CalibrationId = &CalibrationId
	}
    
	if ContextId, ok := AnalyticsevaluationMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if Deleted, ok := AnalyticsevaluationMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if EvaluationId, ok := AnalyticsevaluationMap["evaluationId"].(string); ok {
		o.EvaluationId = &EvaluationId
	}
    
	if EvaluatorId, ok := AnalyticsevaluationMap["evaluatorId"].(string); ok {
		o.EvaluatorId = &EvaluatorId
	}
    
	if eventTimeString, ok := AnalyticsevaluationMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if FormId, ok := AnalyticsevaluationMap["formId"].(string); ok {
		o.FormId = &FormId
	}
    
	if FormName, ok := AnalyticsevaluationMap["formName"].(string); ok {
		o.FormName = &FormName
	}
    
	if QueueId, ok := AnalyticsevaluationMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Released, ok := AnalyticsevaluationMap["released"].(bool); ok {
		o.Released = &Released
	}
    
	if Rescored, ok := AnalyticsevaluationMap["rescored"].(bool); ok {
		o.Rescored = &Rescored
	}
    
	if UserId, ok := AnalyticsevaluationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if OTotalCriticalScore, ok := AnalyticsevaluationMap["oTotalCriticalScore"].(float64); ok {
		OTotalCriticalScoreInt := int(OTotalCriticalScore)
		o.OTotalCriticalScore = &OTotalCriticalScoreInt
	}
	
	if OTotalScore, ok := AnalyticsevaluationMap["oTotalScore"].(float64); ok {
		OTotalScoreInt := int(OTotalScore)
		o.OTotalScore = &OTotalScoreInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsevaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
