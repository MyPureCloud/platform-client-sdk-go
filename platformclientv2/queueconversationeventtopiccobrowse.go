package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopiccobrowse
type Queueconversationeventtopiccobrowse struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Queueconversationeventtopicaddress `json:"self,omitempty"`


	// RoomId - The room id for the chat.
	RoomId *string `json:"roomId,omitempty"`


	// CobrowseSessionId - The co-browse session ID.
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`


	// CobrowseRole - This value identifies the role of the co-browse client within the co-browse session (a client is a sharer or a viewer).
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// Controlling - ID of co-browse participants for which this client has been granted control (list is empty if this client cannot control any shared pages).
	Controlling *[]string `json:"controlling,omitempty"`


	// ViewerUrl - The URL that can be used to open co-browse session in web browser.
	ViewerUrl *string `json:"viewerUrl,omitempty"`


	// Provider - The source provider of the co-browse communication.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// ProviderEventTime - The time when the provider event which triggered this conversation update happened in the corrected provider clock (milliseconds since 1970-01-01 00:00:00 UTC).
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

func (o *Queueconversationeventtopiccobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopiccobrowse
	
	ProviderEventTime := new(string)
	if o.ProviderEventTime != nil {
		
		*ProviderEventTime = timeutil.Strftime(o.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ProviderEventTime = nil
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
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Self *Queueconversationeventtopicaddress `json:"self,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		Controlling *[]string `json:"controlling,omitempty"`
		
		ViewerUrl *string `json:"viewerUrl,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ProviderEventTime *string `json:"providerEventTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		DisconnectType: o.DisconnectType,
		
		Id: o.Id,
		
		Self: o.Self,
		
		RoomId: o.RoomId,
		
		CobrowseSessionId: o.CobrowseSessionId,
		
		CobrowseRole: o.CobrowseRole,
		
		Controlling: o.Controlling,
		
		ViewerUrl: o.ViewerUrl,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		ProviderEventTime: ProviderEventTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopiccobrowse) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopiccobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopiccobrowseMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationeventtopiccobrowseMap["state"].(string); ok {
		o.State = &State
	}
	
	if DisconnectType, ok := QueueconversationeventtopiccobrowseMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if Id, ok := QueueconversationeventtopiccobrowseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Self, ok := QueueconversationeventtopiccobrowseMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if RoomId, ok := QueueconversationeventtopiccobrowseMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
	
	if CobrowseSessionId, ok := QueueconversationeventtopiccobrowseMap["cobrowseSessionId"].(string); ok {
		o.CobrowseSessionId = &CobrowseSessionId
	}
	
	if CobrowseRole, ok := QueueconversationeventtopiccobrowseMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}
	
	if Controlling, ok := QueueconversationeventtopiccobrowseMap["controlling"].([]interface{}); ok {
		ControllingString, _ := json.Marshal(Controlling)
		json.Unmarshal(ControllingString, &o.Controlling)
	}
	
	if ViewerUrl, ok := QueueconversationeventtopiccobrowseMap["viewerUrl"].(string); ok {
		o.ViewerUrl = &ViewerUrl
	}
	
	if Provider, ok := QueueconversationeventtopiccobrowseMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationeventtopiccobrowseMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationeventtopiccobrowseMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if providerEventTimeString, ok := QueueconversationeventtopiccobrowseMap["providerEventTime"].(string); ok {
		ProviderEventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", providerEventTimeString)
		o.ProviderEventTime = &ProviderEventTime
	}
	
	if connectedTimeString, ok := QueueconversationeventtopiccobrowseMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationeventtopiccobrowseMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationeventtopiccobrowseMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationeventtopiccobrowseMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationeventtopiccobrowseMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopiccobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
