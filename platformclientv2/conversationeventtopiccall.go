package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopiccall
type Conversationeventtopiccall struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Recording
	Recording *bool `json:"recording,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Muted
	Muted *bool `json:"muted,omitempty"`


	// Confined
	Confined *bool `json:"confined,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DocumentId
	DocumentId *string `json:"documentId,omitempty"`


	// Self
	Self *Conversationeventtopicaddress `json:"self,omitempty"`


	// Other
	Other *Conversationeventtopicaddress `json:"other,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// DisconnectReasons
	DisconnectReasons *[]Conversationeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`


	// FaxStatus
	FaxStatus *Conversationeventtopicfaxstatus `json:"faxStatus,omitempty"`


	// UuiData
	UuiData *string `json:"uuiData,omitempty"`


	// Wrapup
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Conversationeventtopiccall) MarshalJSON() ([]byte, error) {
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
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
		
		Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
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
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
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
	
	if Recording, ok := ConversationeventtopiccallMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
	
	if RecordingState, ok := ConversationeventtopiccallMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
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
	
	if AdditionalProperties, ok := ConversationeventtopiccallMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopiccall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
