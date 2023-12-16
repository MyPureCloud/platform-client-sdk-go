package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsocialexpressioneventtopicqueuemediasettings
type Conversationsocialexpressioneventtopicqueuemediasettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AlertingTimeoutSeconds - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`

	// AutoAnswerAlertToneSeconds - Specifies the duration of the alerting sound to be played for auto answered interactions.
	AutoAnswerAlertToneSeconds *float32 `json:"autoAnswerAlertToneSeconds,omitempty"`

	// ManualAnswerAlertToneSeconds - Specifies the duration of the alerting sound to be played for manually answered interactions
	ManualAnswerAlertToneSeconds *float32 `json:"manualAnswerAlertToneSeconds,omitempty"`

	// EnableAutoAnswer - Flag to indicate if auto answer is enabled for the given media type or media subtype.
	EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationsocialexpressioneventtopicqueuemediasettings) SetField(field string, fieldValue interface{}) {
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

func (o Conversationsocialexpressioneventtopicqueuemediasettings) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
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
	type Alias Conversationsocialexpressioneventtopicqueuemediasettings
	
	return json.Marshal(&struct { 
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		AutoAnswerAlertToneSeconds *float32 `json:"autoAnswerAlertToneSeconds,omitempty"`
		
		ManualAnswerAlertToneSeconds *float32 `json:"manualAnswerAlertToneSeconds,omitempty"`
		
		EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
		Alias
	}{ 
		AlertingTimeoutSeconds: o.AlertingTimeoutSeconds,
		
		AutoAnswerAlertToneSeconds: o.AutoAnswerAlertToneSeconds,
		
		ManualAnswerAlertToneSeconds: o.ManualAnswerAlertToneSeconds,
		
		EnableAutoAnswer: o.EnableAutoAnswer,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationsocialexpressioneventtopicqueuemediasettings) UnmarshalJSON(b []byte) error {
	var ConversationsocialexpressioneventtopicqueuemediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsocialexpressioneventtopicqueuemediasettingsMap)
	if err != nil {
		return err
	}
	
	if AlertingTimeoutSeconds, ok := ConversationsocialexpressioneventtopicqueuemediasettingsMap["alertingTimeoutSeconds"].(float64); ok {
		AlertingTimeoutSecondsInt := int(AlertingTimeoutSeconds)
		o.AlertingTimeoutSeconds = &AlertingTimeoutSecondsInt
	}
	
	if AutoAnswerAlertToneSeconds, ok := ConversationsocialexpressioneventtopicqueuemediasettingsMap["autoAnswerAlertToneSeconds"].(float64); ok {
		AutoAnswerAlertToneSecondsFloat32 := float32(AutoAnswerAlertToneSeconds)
		o.AutoAnswerAlertToneSeconds = &AutoAnswerAlertToneSecondsFloat32
	}
    
	if ManualAnswerAlertToneSeconds, ok := ConversationsocialexpressioneventtopicqueuemediasettingsMap["manualAnswerAlertToneSeconds"].(float64); ok {
		ManualAnswerAlertToneSecondsFloat32 := float32(ManualAnswerAlertToneSeconds)
		o.ManualAnswerAlertToneSeconds = &ManualAnswerAlertToneSecondsFloat32
	}
    
	if EnableAutoAnswer, ok := ConversationsocialexpressioneventtopicqueuemediasettingsMap["enableAutoAnswer"].(bool); ok {
		o.EnableAutoAnswer = &EnableAutoAnswer
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsocialexpressioneventtopicqueuemediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
