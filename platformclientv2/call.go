package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Call
type Call struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`

	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// Direction - The direction of the call
	Direction *string `json:"direction,omitempty"`

	// Recording - True if this call is being recorded.
	Recording *bool `json:"recording,omitempty"`

	// RecordingState - State of recording on this call.
	RecordingState *string `json:"recordingState,omitempty"`

	// Muted - True if this call is muted so that remote participants can't hear any audio from this end.
	Muted *bool `json:"muted,omitempty"`

	// Confined - True if this call is held and the person on this side hears hold music.
	Confined *bool `json:"confined,omitempty"`

	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`

	// SecurePause - True when the recording of this call is in secure pause status.
	SecurePause *bool `json:"securePause,omitempty"`

	// RecordingId - A globally unique identifier for the recording associated with this call.
	RecordingId *string `json:"recordingId,omitempty"`

	// Segments - The time line of the participant's call, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`

	// ErrorInfo
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartHoldTime - The timestamp the call was placed on hold in the cloud clock if the call is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// DocumentId - If call is an outbound fax of a document from content management, then this is the id in content management.
	DocumentId *string `json:"documentId,omitempty"`

	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// DisconnectReasons - List of reasons that this call was disconnected. This will be set once the call disconnects.
	DisconnectReasons *[]Disconnectreason `json:"disconnectReasons,omitempty"`

	// FaxStatus - Extra information on fax transmission.
	FaxStatus *Faxstatus `json:"faxStatus,omitempty"`

	// Provider - The source provider for the call.
	Provider *string `json:"provider,omitempty"`

	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`

	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`

	// Other - Address and name data for a call endpoint.
	Other *Address `json:"other,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`

	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

	// AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

	// Disposition - Call resolution data for Dialer bulk make calls commands.
	Disposition *Disposition `json:"disposition,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Call) SetField(field string, fieldValue interface{}) {
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

func (o Call) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartHoldTime","StartAlertingTime","ConnectedTime","DisconnectedTime", }
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
	type Alias Call
	
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
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		SecurePause *bool `json:"securePause,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		DisconnectReasons *[]Disconnectreason `json:"disconnectReasons,omitempty"`
		
		FaxStatus *Faxstatus `json:"faxStatus,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		Self *Address `json:"self,omitempty"`
		
		Other *Address `json:"other,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		Disposition *Disposition `json:"disposition,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		InitialState: o.InitialState,
		
		Id: o.Id,
		
		Direction: o.Direction,
		
		Recording: o.Recording,
		
		RecordingState: o.RecordingState,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Held: o.Held,
		
		SecurePause: o.SecurePause,
		
		RecordingId: o.RecordingId,
		
		Segments: o.Segments,
		
		ErrorInfo: o.ErrorInfo,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		DocumentId: o.DocumentId,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		DisconnectReasons: o.DisconnectReasons,
		
		FaxStatus: o.FaxStatus,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		UuiData: o.UuiData,
		
		Self: o.Self,
		
		Other: o.Other,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AgentAssistantId: o.AgentAssistantId,
		
		Disposition: o.Disposition,
		Alias:    (Alias)(o),
	})
}

func (o *Call) UnmarshalJSON(b []byte) error {
	var CallMap map[string]interface{}
	err := json.Unmarshal(b, &CallMap)
	if err != nil {
		return err
	}
	
	if State, ok := CallMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := CallMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Id, ok := CallMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Direction, ok := CallMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Recording, ok := CallMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if RecordingState, ok := CallMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if Muted, ok := CallMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Confined, ok := CallMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
    
	if Held, ok := CallMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if SecurePause, ok := CallMap["securePause"].(bool); ok {
		o.SecurePause = &SecurePause
	}
    
	if RecordingId, ok := CallMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
    
	if Segments, ok := CallMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	
	if ErrorInfo, ok := CallMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := CallMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startHoldTimeString, ok := CallMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if DocumentId, ok := CallMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    
	if startAlertingTimeString, ok := CallMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}
	
	if connectedTimeString, ok := CallMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := CallMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if DisconnectReasons, ok := CallMap["disconnectReasons"].([]interface{}); ok {
		DisconnectReasonsString, _ := json.Marshal(DisconnectReasons)
		json.Unmarshal(DisconnectReasonsString, &o.DisconnectReasons)
	}
	
	if FaxStatus, ok := CallMap["faxStatus"].(map[string]interface{}); ok {
		FaxStatusString, _ := json.Marshal(FaxStatus)
		json.Unmarshal(FaxStatusString, &o.FaxStatus)
	}
	
	if Provider, ok := CallMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := CallMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := CallMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if UuiData, ok := CallMap["uuiData"].(string); ok {
		o.UuiData = &UuiData
	}
    
	if Self, ok := CallMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Other, ok := CallMap["other"].(map[string]interface{}); ok {
		OtherString, _ := json.Marshal(Other)
		json.Unmarshal(OtherString, &o.Other)
	}
	
	if Wrapup, ok := CallMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := CallMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := CallMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if AgentAssistantId, ok := CallMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if Disposition, ok := CallMap["disposition"].(map[string]interface{}); ok {
		DispositionString, _ := json.Marshal(Disposition)
		json.Unmarshal(DispositionString, &o.Disposition)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Call) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
