package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackmediasettings
type Callbackmediasettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
	EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`

	// AlertingTimeoutSeconds - The alerting timeout for the media type, in seconds
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`

	// ServiceLevel - The targeted service level for the media type
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`

	// AutoAnswerAlertToneSeconds - How long to play the alerting tone for an auto-answer interaction
	AutoAnswerAlertToneSeconds *float64 `json:"autoAnswerAlertToneSeconds,omitempty"`

	// ManualAnswerAlertToneSeconds - How long to play the alerting tone for a manual-answer interaction
	ManualAnswerAlertToneSeconds *float64 `json:"manualAnswerAlertToneSeconds,omitempty"`

	// SubTypeSettings - Map of media subtype to media subtype specific settings.
	SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`

	// Mode - The mode callbacks will use on this queue.
	Mode *string `json:"mode,omitempty"`

	// EnableAutoDialAndEnd - Flag to enable Auto-Dial and Auto-End automation for callbacks on this queue.
	EnableAutoDialAndEnd *bool `json:"enableAutoDialAndEnd,omitempty"`

	// AutoDialDelaySeconds - Time in seconds after agent connects to callback before outgoing call is auto-dialed. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
	AutoDialDelaySeconds *int `json:"autoDialDelaySeconds,omitempty"`

	// AutoEndDelaySeconds - Time in seconds after agent disconnects from the outgoing call before the encasing callback is auto-ended. Allowable values in range 0 - 1200 seconds. Defaults to 300 seconds.
	AutoEndDelaySeconds *int `json:"autoEndDelaySeconds,omitempty"`

	// PacingModifier - Controls the maximum number of outbound calls at one time when mode is CustomerFirst.
	PacingModifier *float64 `json:"pacingModifier,omitempty"`

	// MaxRetryCount - Maximum number of retries that should be attempted to try and connect a customer first callback to a customer when the initial callback attempt did not connect.
	MaxRetryCount *int `json:"maxRetryCount,omitempty"`

	// RetryDelaySeconds - Delay in seconds between each retry of a customer first callback.
	RetryDelaySeconds *int `json:"retryDelaySeconds,omitempty"`

	// LiveVoiceReactionType - The action to take if a live voice is detected during the outbound call of a customer first callback.
	LiveVoiceReactionType *string `json:"liveVoiceReactionType,omitempty"`

	// LiveVoiceFlow - The inbound flow to transfer to if a live voice is detected during the outbound call of a customer first callback.
	LiveVoiceFlow *Domainentityref `json:"liveVoiceFlow,omitempty"`

	// AnsweringMachineReactionType - The action to take if an answering machine is detected during the outbound call of a customer first callback.
	AnsweringMachineReactionType *string `json:"answeringMachineReactionType,omitempty"`

	// AnsweringMachineFlow - The inbound flow to transfer to if an answering machine is detected during the outbound call of a customer first callback when answeringMachineReactionType is set to TransferToFlow.
	AnsweringMachineFlow *Domainentityref `json:"answeringMachineFlow,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callbackmediasettings) SetField(field string, fieldValue interface{}) {
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

func (o Callbackmediasettings) MarshalJSON() ([]byte, error) {
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
	type Alias Callbackmediasettings
	
	return json.Marshal(&struct { 
		EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
		
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		AutoAnswerAlertToneSeconds *float64 `json:"autoAnswerAlertToneSeconds,omitempty"`
		
		ManualAnswerAlertToneSeconds *float64 `json:"manualAnswerAlertToneSeconds,omitempty"`
		
		SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`
		
		Mode *string `json:"mode,omitempty"`
		
		EnableAutoDialAndEnd *bool `json:"enableAutoDialAndEnd,omitempty"`
		
		AutoDialDelaySeconds *int `json:"autoDialDelaySeconds,omitempty"`
		
		AutoEndDelaySeconds *int `json:"autoEndDelaySeconds,omitempty"`
		
		PacingModifier *float64 `json:"pacingModifier,omitempty"`
		
		MaxRetryCount *int `json:"maxRetryCount,omitempty"`
		
		RetryDelaySeconds *int `json:"retryDelaySeconds,omitempty"`
		
		LiveVoiceReactionType *string `json:"liveVoiceReactionType,omitempty"`
		
		LiveVoiceFlow *Domainentityref `json:"liveVoiceFlow,omitempty"`
		
		AnsweringMachineReactionType *string `json:"answeringMachineReactionType,omitempty"`
		
		AnsweringMachineFlow *Domainentityref `json:"answeringMachineFlow,omitempty"`
		Alias
	}{ 
		EnableAutoAnswer: o.EnableAutoAnswer,
		
		AlertingTimeoutSeconds: o.AlertingTimeoutSeconds,
		
		ServiceLevel: o.ServiceLevel,
		
		AutoAnswerAlertToneSeconds: o.AutoAnswerAlertToneSeconds,
		
		ManualAnswerAlertToneSeconds: o.ManualAnswerAlertToneSeconds,
		
		SubTypeSettings: o.SubTypeSettings,
		
		Mode: o.Mode,
		
		EnableAutoDialAndEnd: o.EnableAutoDialAndEnd,
		
		AutoDialDelaySeconds: o.AutoDialDelaySeconds,
		
		AutoEndDelaySeconds: o.AutoEndDelaySeconds,
		
		PacingModifier: o.PacingModifier,
		
		MaxRetryCount: o.MaxRetryCount,
		
		RetryDelaySeconds: o.RetryDelaySeconds,
		
		LiveVoiceReactionType: o.LiveVoiceReactionType,
		
		LiveVoiceFlow: o.LiveVoiceFlow,
		
		AnsweringMachineReactionType: o.AnsweringMachineReactionType,
		
		AnsweringMachineFlow: o.AnsweringMachineFlow,
		Alias:    (Alias)(o),
	})
}

func (o *Callbackmediasettings) UnmarshalJSON(b []byte) error {
	var CallbackmediasettingsMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackmediasettingsMap)
	if err != nil {
		return err
	}
	
	if EnableAutoAnswer, ok := CallbackmediasettingsMap["enableAutoAnswer"].(bool); ok {
		o.EnableAutoAnswer = &EnableAutoAnswer
	}
    
	if AlertingTimeoutSeconds, ok := CallbackmediasettingsMap["alertingTimeoutSeconds"].(float64); ok {
		AlertingTimeoutSecondsInt := int(AlertingTimeoutSeconds)
		o.AlertingTimeoutSeconds = &AlertingTimeoutSecondsInt
	}
	
	if ServiceLevel, ok := CallbackmediasettingsMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if AutoAnswerAlertToneSeconds, ok := CallbackmediasettingsMap["autoAnswerAlertToneSeconds"].(float64); ok {
		o.AutoAnswerAlertToneSeconds = &AutoAnswerAlertToneSeconds
	}
    
	if ManualAnswerAlertToneSeconds, ok := CallbackmediasettingsMap["manualAnswerAlertToneSeconds"].(float64); ok {
		o.ManualAnswerAlertToneSeconds = &ManualAnswerAlertToneSeconds
	}
    
	if SubTypeSettings, ok := CallbackmediasettingsMap["subTypeSettings"].(map[string]interface{}); ok {
		SubTypeSettingsString, _ := json.Marshal(SubTypeSettings)
		json.Unmarshal(SubTypeSettingsString, &o.SubTypeSettings)
	}
	
	if Mode, ok := CallbackmediasettingsMap["mode"].(string); ok {
		o.Mode = &Mode
	}
    
	if EnableAutoDialAndEnd, ok := CallbackmediasettingsMap["enableAutoDialAndEnd"].(bool); ok {
		o.EnableAutoDialAndEnd = &EnableAutoDialAndEnd
	}
    
	if AutoDialDelaySeconds, ok := CallbackmediasettingsMap["autoDialDelaySeconds"].(float64); ok {
		AutoDialDelaySecondsInt := int(AutoDialDelaySeconds)
		o.AutoDialDelaySeconds = &AutoDialDelaySecondsInt
	}
	
	if AutoEndDelaySeconds, ok := CallbackmediasettingsMap["autoEndDelaySeconds"].(float64); ok {
		AutoEndDelaySecondsInt := int(AutoEndDelaySeconds)
		o.AutoEndDelaySeconds = &AutoEndDelaySecondsInt
	}
	
	if PacingModifier, ok := CallbackmediasettingsMap["pacingModifier"].(float64); ok {
		o.PacingModifier = &PacingModifier
	}
    
	if MaxRetryCount, ok := CallbackmediasettingsMap["maxRetryCount"].(float64); ok {
		MaxRetryCountInt := int(MaxRetryCount)
		o.MaxRetryCount = &MaxRetryCountInt
	}
	
	if RetryDelaySeconds, ok := CallbackmediasettingsMap["retryDelaySeconds"].(float64); ok {
		RetryDelaySecondsInt := int(RetryDelaySeconds)
		o.RetryDelaySeconds = &RetryDelaySecondsInt
	}
	
	if LiveVoiceReactionType, ok := CallbackmediasettingsMap["liveVoiceReactionType"].(string); ok {
		o.LiveVoiceReactionType = &LiveVoiceReactionType
	}
    
	if LiveVoiceFlow, ok := CallbackmediasettingsMap["liveVoiceFlow"].(map[string]interface{}); ok {
		LiveVoiceFlowString, _ := json.Marshal(LiveVoiceFlow)
		json.Unmarshal(LiveVoiceFlowString, &o.LiveVoiceFlow)
	}
	
	if AnsweringMachineReactionType, ok := CallbackmediasettingsMap["answeringMachineReactionType"].(string); ok {
		o.AnsweringMachineReactionType = &AnsweringMachineReactionType
	}
    
	if AnsweringMachineFlow, ok := CallbackmediasettingsMap["answeringMachineFlow"].(map[string]interface{}); ok {
		AnsweringMachineFlowString, _ := json.Marshal(AnsweringMachineFlow)
		json.Unmarshal(AnsweringMachineFlowString, &o.AnsweringMachineFlow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackmediasettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
