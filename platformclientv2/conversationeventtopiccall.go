package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopiccall
type Conversationeventtopiccall struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

	// Recording - True if this call is being recorded.
	Recording *bool `json:"recording,omitempty"`

	// RecordingState - State of recording on this call.
	RecordingState *string `json:"recordingState,omitempty"`

	// RecordersState
	RecordersState *Conversationeventtopicrecordersstate `json:"recordersState,omitempty"`

	// Muted - True if this call is muted so that remote participants can't hear any audio from this end.
	Muted *bool `json:"muted,omitempty"`

	// Confined - True if this call is held and the person on this side hears hold music.
	Confined *bool `json:"confined,omitempty"`

	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`

	// SecurePause - True when the recording of this call is in secure pause status.
	SecurePause *bool `json:"securePause,omitempty"`

	// ErrorInfo
	ErrorInfo *Conversationeventtopicerrordetails `json:"errorInfo,omitempty"`

	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// StartHoldTime - The timestamp the call was placed on hold in the cloud clock if the call is currently on hold.
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`

	// Direction - Whether a call is inbound or outbound.
	Direction *string `json:"direction,omitempty"`

	// DocumentId - If call is a fax of a document in content management, the id of the document in content management.
	DocumentId *string `json:"documentId,omitempty"`

	// Self
	Self *Conversationeventtopicaddress `json:"self,omitempty"`

	// Other - Address and name data for a call endpoint.
	Other *Conversationeventtopicaddress `json:"other,omitempty"`

	// Provider - The source provider of the call.
	Provider *string `json:"provider,omitempty"`

	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`

	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`

	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`

	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`

	// DisconnectReasons - List of reasons that this call was disconnected. This will be set once the call disconnects.
	DisconnectReasons *[]Conversationeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`

	// FaxStatus
	FaxStatus *Conversationeventtopicfaxstatus `json:"faxStatus,omitempty"`

	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`

	// BargedTime - The timestamp when this participant was connected to the barge conference in the provider clock.
	BargedTime *time.Time `json:"bargedTime,omitempty"`

	// Wrapup - Call wrap up or disposition data.
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`

	// AfterCallWork
	AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`

	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

	// AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

	// QueueMediaSettings
	QueueMediaSettings *Conversationeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationeventtopiccall) SetField(field string, fieldValue interface{}) {
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

func (o Conversationeventtopiccall) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartHoldTime","ConnectedTime","DisconnectedTime","BargedTime", }
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
	type Alias Conversationeventtopiccall
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
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
	
	BargedTime := new(string)
	if o.BargedTime != nil {
		
		*BargedTime = timeutil.Strftime(o.BargedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BargedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		RecordersState *Conversationeventtopicrecordersstate `json:"recordersState,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		SecurePause *bool `json:"securePause,omitempty"`
		
		ErrorInfo *Conversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		Self *Conversationeventtopicaddress `json:"self,omitempty"`
		
		Other *Conversationeventtopicaddress `json:"other,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		DisconnectReasons *[]Conversationeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`
		
		FaxStatus *Conversationeventtopicfaxstatus `json:"faxStatus,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		BargedTime *string `json:"bargedTime,omitempty"`
		
		Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		QueueMediaSettings *Conversationeventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Recording: o.Recording,
		
		RecordingState: o.RecordingState,
		
		RecordersState: o.RecordersState,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Held: o.Held,
		
		SecurePause: o.SecurePause,
		
		ErrorInfo: o.ErrorInfo,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		Direction: o.Direction,
		
		DocumentId: o.DocumentId,
		
		Self: o.Self,
		
		Other: o.Other,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		DisconnectReasons: o.DisconnectReasons,
		
		FaxStatus: o.FaxStatus,
		
		UuiData: o.UuiData,
		
		BargedTime: BargedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AgentAssistantId: o.AgentAssistantId,
		
		QueueMediaSettings: o.QueueMediaSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationeventtopiccall) UnmarshalJSON(b []byte) error {
	var ConversationeventtopiccallMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopiccallMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationeventtopiccallMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := ConversationeventtopiccallMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationeventtopiccallMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Recording, ok := ConversationeventtopiccallMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if RecordingState, ok := ConversationeventtopiccallMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if RecordersState, ok := ConversationeventtopiccallMap["recordersState"].(map[string]interface{}); ok {
		RecordersStateString, _ := json.Marshal(RecordersState)
		json.Unmarshal(RecordersStateString, &o.RecordersState)
	}
	
	if Muted, ok := ConversationeventtopiccallMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Confined, ok := ConversationeventtopiccallMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
    
	if Held, ok := ConversationeventtopiccallMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if SecurePause, ok := ConversationeventtopiccallMap["securePause"].(bool); ok {
		o.SecurePause = &SecurePause
	}
    
	if ErrorInfo, ok := ConversationeventtopiccallMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := ConversationeventtopiccallMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startHoldTimeString, ok := ConversationeventtopiccallMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Direction, ok := ConversationeventtopiccallMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DocumentId, ok := ConversationeventtopiccallMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    
	if Self, ok := ConversationeventtopiccallMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Other, ok := ConversationeventtopiccallMap["other"].(map[string]interface{}); ok {
		OtherString, _ := json.Marshal(Other)
		json.Unmarshal(OtherString, &o.Other)
	}
	
	if Provider, ok := ConversationeventtopiccallMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := ConversationeventtopiccallMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := ConversationeventtopiccallMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if connectedTimeString, ok := ConversationeventtopiccallMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := ConversationeventtopiccallMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if DisconnectReasons, ok := ConversationeventtopiccallMap["disconnectReasons"].([]interface{}); ok {
		DisconnectReasonsString, _ := json.Marshal(DisconnectReasons)
		json.Unmarshal(DisconnectReasonsString, &o.DisconnectReasons)
	}
	
	if FaxStatus, ok := ConversationeventtopiccallMap["faxStatus"].(map[string]interface{}); ok {
		FaxStatusString, _ := json.Marshal(FaxStatus)
		json.Unmarshal(FaxStatusString, &o.FaxStatus)
	}
	
	if UuiData, ok := ConversationeventtopiccallMap["uuiData"].(string); ok {
		o.UuiData = &UuiData
	}
    
	if bargedTimeString, ok := ConversationeventtopiccallMap["bargedTime"].(string); ok {
		BargedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", bargedTimeString)
		o.BargedTime = &BargedTime
	}
	
	if Wrapup, ok := ConversationeventtopiccallMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := ConversationeventtopiccallMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := ConversationeventtopiccallMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if AgentAssistantId, ok := ConversationeventtopiccallMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if QueueMediaSettings, ok := ConversationeventtopiccallMap["queueMediaSettings"].(map[string]interface{}); ok {
		QueueMediaSettingsString, _ := json.Marshal(QueueMediaSettings)
		json.Unmarshal(QueueMediaSettingsString, &o.QueueMediaSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopiccall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
