package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopiccall
type Queueconversationeventtopiccall struct { 
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
	ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DocumentId
	DocumentId *string `json:"documentId,omitempty"`


	// Self
	Self *Queueconversationeventtopicaddress `json:"self,omitempty"`


	// Other
	Other *Queueconversationeventtopicaddress `json:"other,omitempty"`


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
	DisconnectReasons *[]Queueconversationeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`


	// FaxStatus
	FaxStatus *Queueconversationeventtopicfaxstatus `json:"faxStatus,omitempty"`


	// UuiData
	UuiData *string `json:"uuiData,omitempty"`


	// Wrapup
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Queueconversationeventtopiccall) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopiccall

	
	StartHoldTime := new(string)
	if u.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(u.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if u.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(u.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		Self *Queueconversationeventtopicaddress `json:"self,omitempty"`
		
		Other *Queueconversationeventtopicaddress `json:"other,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		DisconnectReasons *[]Queueconversationeventtopicdisconnectreason `json:"disconnectReasons,omitempty"`
		
		FaxStatus *Queueconversationeventtopicfaxstatus `json:"faxStatus,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		
		Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		State: u.State,
		
		Recording: u.Recording,
		
		RecordingState: u.RecordingState,
		
		Muted: u.Muted,
		
		Confined: u.Confined,
		
		Held: u.Held,
		
		ErrorInfo: u.ErrorInfo,
		
		DisconnectType: u.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		Direction: u.Direction,
		
		DocumentId: u.DocumentId,
		
		Self: u.Self,
		
		Other: u.Other,
		
		Provider: u.Provider,
		
		ScriptId: u.ScriptId,
		
		PeerId: u.PeerId,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		DisconnectReasons: u.DisconnectReasons,
		
		FaxStatus: u.FaxStatus,
		
		UuiData: u.UuiData,
		
		Wrapup: u.Wrapup,
		
		AfterCallWork: u.AfterCallWork,
		
		AfterCallWorkRequired: u.AfterCallWorkRequired,
		
		AgentAssistantId: u.AgentAssistantId,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopiccall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
