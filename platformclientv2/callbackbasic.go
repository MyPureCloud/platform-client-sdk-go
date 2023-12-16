package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callbackbasic
type Callbackbasic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`

	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// Segments - The time line of the participant's callback, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`

	// Direction - The direction of the call
	Direction *string `json:"direction,omitempty"`

	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartHoldTime - The timestamp the callback was placed on hold in the cloud clock if the callback is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// DialerPreview - The preview data to be used when this callback is a Preview.
	DialerPreview *Dialerpreview `json:"dialerPreview,omitempty"`

	// Voicemail - The voicemail data to be used when this callback is an ACD voicemail.
	Voicemail *Voicemail `json:"voicemail,omitempty"`

	// CallbackNumbers - The phone number(s) to use to place the callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackUserName - The name of the user requesting a callback.
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`

	// ExternalCampaign - True if the call for the callback uses external dialing.
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`

	// SkipEnabled - True if the ability to skip a callback should be enabled.
	SkipEnabled *bool `json:"skipEnabled,omitempty"`

	// TimeoutSeconds - The number of seconds before the system automatically places a call for a callback.  0 means the automatic placement is disabled.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`

	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// CallbackScheduledTime - The timestamp when this communication is scheduled in the provider clock. If this value is missing it indicates the callback will be placed immediately. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

	// AutomatedCallbackConfigId - The id of the config for automatically placing the callback (and handling the disposition). If null, the callback will not be placed automatically but routed to an agent as per normal.
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`

	// Provider - The source provider for the callback.
	Provider *string `json:"provider,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

	// CallerId - The phone number displayed to recipients of the phone call. The value should conform to the E164 format.
	CallerId *string `json:"callerId,omitempty"`

	// CallerIdName - The name displayed to recipients of the phone call.
	CallerIdName *string `json:"callerIdName,omitempty"`

	// QueueMediaSettings - Represents the queue settings for this media type.
	QueueMediaSettings *Conversationqueuemediasettings `json:"queueMediaSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callbackbasic) SetField(field string, fieldValue interface{}) {
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

func (o Callbackbasic) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartHoldTime","StartAlertingTime","ConnectedTime","DisconnectedTime","CallbackScheduledTime", }
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
	type Alias Callbackbasic
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAlertingTime := new(string)
	if o.StartAlertingTime != nil {
		
		*StartAlertingTime = timeutil.Strftime(o.StartAlertingTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAlertingTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	CallbackScheduledTime := new(string)
	if o.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(o.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		DialerPreview *Dialerpreview `json:"dialerPreview,omitempty"`
		
		Voicemail *Voicemail `json:"voicemail,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		ExternalCampaign *bool `json:"externalCampaign,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		
		QueueMediaSettings *Conversationqueuemediasettings `json:"queueMediaSettings,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		InitialState: o.InitialState,
		
		Id: o.Id,
		
		Segments: o.Segments,
		
		Direction: o.Direction,
		
		Held: o.Held,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		DialerPreview: o.DialerPreview,
		
		Voicemail: o.Voicemail,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackUserName: o.CallbackUserName,
		
		ScriptId: o.ScriptId,
		
		ExternalCampaign: o.ExternalCampaign,
		
		SkipEnabled: o.SkipEnabled,
		
		TimeoutSeconds: o.TimeoutSeconds,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		AutomatedCallbackConfigId: o.AutomatedCallbackConfigId,
		
		Provider: o.Provider,
		
		PeerId: o.PeerId,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		CallerId: o.CallerId,
		
		CallerIdName: o.CallerIdName,
		
		QueueMediaSettings: o.QueueMediaSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Callbackbasic) UnmarshalJSON(b []byte) error {
	var CallbackbasicMap map[string]interface{}
	err := json.Unmarshal(b, &CallbackbasicMap)
	if err != nil {
		return err
	}
	
	if State, ok := CallbackbasicMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := CallbackbasicMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Id, ok := CallbackbasicMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Segments, ok := CallbackbasicMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	
	if Direction, ok := CallbackbasicMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Held, ok := CallbackbasicMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if DisconnectType, ok := CallbackbasicMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startHoldTimeString, ok := CallbackbasicMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if DialerPreview, ok := CallbackbasicMap["dialerPreview"].(map[string]interface{}); ok {
		DialerPreviewString, _ := json.Marshal(DialerPreview)
		json.Unmarshal(DialerPreviewString, &o.DialerPreview)
	}
	
	if Voicemail, ok := CallbackbasicMap["voicemail"].(map[string]interface{}); ok {
		VoicemailString, _ := json.Marshal(Voicemail)
		json.Unmarshal(VoicemailString, &o.Voicemail)
	}
	
	if CallbackNumbers, ok := CallbackbasicMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if CallbackUserName, ok := CallbackbasicMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if ScriptId, ok := CallbackbasicMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if ExternalCampaign, ok := CallbackbasicMap["externalCampaign"].(bool); ok {
		o.ExternalCampaign = &ExternalCampaign
	}
    
	if SkipEnabled, ok := CallbackbasicMap["skipEnabled"].(bool); ok {
		o.SkipEnabled = &SkipEnabled
	}
    
	if TimeoutSeconds, ok := CallbackbasicMap["timeoutSeconds"].(float64); ok {
		TimeoutSecondsInt := int(TimeoutSeconds)
		o.TimeoutSeconds = &TimeoutSecondsInt
	}
	
	if startAlertingTimeString, ok := CallbackbasicMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}
	
	if connectedTimeString, ok := CallbackbasicMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := CallbackbasicMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if callbackScheduledTimeString, ok := CallbackbasicMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	
	if AutomatedCallbackConfigId, ok := CallbackbasicMap["automatedCallbackConfigId"].(string); ok {
		o.AutomatedCallbackConfigId = &AutomatedCallbackConfigId
	}
    
	if Provider, ok := CallbackbasicMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if PeerId, ok := CallbackbasicMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if Wrapup, ok := CallbackbasicMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := CallbackbasicMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := CallbackbasicMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if CallerId, ok := CallbackbasicMap["callerId"].(string); ok {
		o.CallerId = &CallerId
	}
    
	if CallerIdName, ok := CallbackbasicMap["callerIdName"].(string); ok {
		o.CallerIdName = &CallerIdName
	}
    
	if QueueMediaSettings, ok := CallbackbasicMap["queueMediaSettings"].(map[string]interface{}); ok {
		QueueMediaSettingsString, _ := json.Marshal(QueueMediaSettings)
		json.Unmarshal(QueueMediaSettingsString, &o.QueueMediaSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callbackbasic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
