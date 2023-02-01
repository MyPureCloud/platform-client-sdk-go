package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsevaluation
type Analyticsevaluation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CalibrationId - The calibration ID used for the purpose of training evaluators
	CalibrationId *string `json:"calibrationId,omitempty"`

	// ContextId - A unique identifier for an evaluation form, regardless of version
	ContextId *string `json:"contextId,omitempty"`

	// Deleted - Whether the evaluation has been deleted
	Deleted *bool `json:"deleted,omitempty"`

	// EvaluationId - Unique identifier for the evaluation
	EvaluationId *string `json:"evaluationId,omitempty"`

	// EvaluationStatus - Status of evaluation
	EvaluationStatus *string `json:"evaluationStatus,omitempty"`

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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsevaluation) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Analyticsevaluation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		EvaluationStatus *string `json:"evaluationStatus,omitempty"`
		
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
		Alias
	}{ 
		CalibrationId: o.CalibrationId,
		
		ContextId: o.ContextId,
		
		Deleted: o.Deleted,
		
		EvaluationId: o.EvaluationId,
		
		EvaluationStatus: o.EvaluationStatus,
		
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
		Alias:    (Alias)(o),
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
    
	if EvaluationStatus, ok := AnalyticsevaluationMap["evaluationStatus"].(string); ok {
		o.EvaluationStatus = &EvaluationStatus
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
