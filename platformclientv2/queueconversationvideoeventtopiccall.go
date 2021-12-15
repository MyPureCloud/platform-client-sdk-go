package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopiccall
type Queueconversationvideoeventtopiccall struct { 
	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


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


	// ErrorInfo
	ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the call was placed on hold in the cloud clock if the call is currently on hold.
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Direction - Whether a call is inbound or outbound.
	Direction *string `json:"direction,omitempty"`


	// DocumentId - If call is a fax of a document in content management, the id of the document in content management.
	DocumentId *string `json:"documentId,omitempty"`


	// Self
	Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`


	// Other - Address and name data for a call endpoint.
	Other *Queueconversationvideoeventtopicaddress `json:"other,omitempty"`


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
	DisconnectReasons *[]Queueconversationvideoeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`


	// FaxStatus
	FaxStatus *Queueconversationvideoeventtopicfaxstatus `json:"faxStatus,omitempty"`


	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`

}

func (o *Queueconversationvideoeventtopiccall) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopiccall
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`
		
		Other *Queueconversationvideoeventtopicaddress `json:"other,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		DisconnectReasons *[]Queueconversationvideoeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`
		
		FaxStatus *Queueconversationvideoeventtopicfaxstatus `json:"faxStatus,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Recording: o.Recording,
		
		RecordingState: o.RecordingState,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Held: o.Held,
		
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
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AgentAssistantId: o.AgentAssistantId,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopiccall) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopiccallMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopiccallMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopiccallMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := QueueconversationvideoeventtopiccallMap["state"].(string); ok {
		o.State = &State
	}
	
	if Recording, ok := QueueconversationvideoeventtopiccallMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
	
	if RecordingState, ok := QueueconversationvideoeventtopiccallMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
	
	if Muted, ok := QueueconversationvideoeventtopiccallMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
	
	if Confined, ok := QueueconversationvideoeventtopiccallMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
	
	if Held, ok := QueueconversationvideoeventtopiccallMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if ErrorInfo, ok := QueueconversationvideoeventtopiccallMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if DisconnectType, ok := QueueconversationvideoeventtopiccallMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationvideoeventtopiccallMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Direction, ok := QueueconversationvideoeventtopiccallMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DocumentId, ok := QueueconversationvideoeventtopiccallMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
	
	if Self, ok := QueueconversationvideoeventtopiccallMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Other, ok := QueueconversationvideoeventtopiccallMap["other"].(map[string]interface{}); ok {
		OtherString, _ := json.Marshal(Other)
		json.Unmarshal(OtherString, &o.Other)
	}
	
	if Provider, ok := QueueconversationvideoeventtopiccallMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationvideoeventtopiccallMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationvideoeventtopiccallMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if connectedTimeString, ok := QueueconversationvideoeventtopiccallMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopiccallMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if DisconnectReasons, ok := QueueconversationvideoeventtopiccallMap["disconnectReasons"].([]interface{}); ok {
		DisconnectReasonsString, _ := json.Marshal(DisconnectReasons)
		json.Unmarshal(DisconnectReasonsString, &o.DisconnectReasons)
	}
	
	if FaxStatus, ok := QueueconversationvideoeventtopiccallMap["faxStatus"].(map[string]interface{}); ok {
		FaxStatusString, _ := json.Marshal(FaxStatus)
		json.Unmarshal(FaxStatusString, &o.FaxStatus)
	}
	
	if UuiData, ok := QueueconversationvideoeventtopiccallMap["uuiData"].(string); ok {
		o.UuiData = &UuiData
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopiccallMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopiccallMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopiccallMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AgentAssistantId, ok := QueueconversationvideoeventtopiccallMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopiccall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
