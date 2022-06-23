package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicvideo
type Queueconversationvideoeventtopicvideo struct { 
	// State
	State *string `json:"state,omitempty"`


	// InitialState
	InitialState *string `json:"initialState,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Context - The room id context (xmpp jid) for the conference session.
	Context *string `json:"context,omitempty"`


	// AudioMuted - Indicates whether this participant has muted their outgoing audio.
	AudioMuted *bool `json:"audioMuted,omitempty"`


	// VideoMuted - Indicates whether this participant has muted/paused their outgoing video.
	VideoMuted *bool `json:"videoMuted,omitempty"`


	// SharingScreen - Indicates whether this participant is sharing their screen to the session.
	SharingScreen *bool `json:"sharingScreen,omitempty"`


	// PeerCount - The number of peer participants from the perspective of the participant in the conference.
	PeerCount *interface{} `json:"peerCount,omitempty"`


	// Provider - The media provider controlling the video.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock.
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock.
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Msids - List of media stream ids
	Msids *[]string `json:"msids,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

func (o *Queueconversationvideoeventtopicvideo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicvideo
	
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
		
		Self *Queueconversationvideoeventtopicaddress `json:"self,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		AudioMuted *bool `json:"audioMuted,omitempty"`
		
		VideoMuted *bool `json:"videoMuted,omitempty"`
		
		SharingScreen *bool `json:"sharingScreen,omitempty"`
		
		PeerCount *interface{} `json:"peerCount,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Msids *[]string `json:"msids,omitempty"`
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		InitialState: o.InitialState,
		
		Self: o.Self,
		
		Id: o.Id,
		
		Context: o.Context,
		
		AudioMuted: o.AudioMuted,
		
		VideoMuted: o.VideoMuted,
		
		SharingScreen: o.SharingScreen,
		
		PeerCount: o.PeerCount,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Msids: o.Msids,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicvideo) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicvideoMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicvideoMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationvideoeventtopicvideoMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationvideoeventtopicvideoMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Self, ok := QueueconversationvideoeventtopicvideoMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Id, ok := QueueconversationvideoeventtopicvideoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Context, ok := QueueconversationvideoeventtopicvideoMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if AudioMuted, ok := QueueconversationvideoeventtopicvideoMap["audioMuted"].(bool); ok {
		o.AudioMuted = &AudioMuted
	}
    
	if VideoMuted, ok := QueueconversationvideoeventtopicvideoMap["videoMuted"].(bool); ok {
		o.VideoMuted = &VideoMuted
	}
    
	if SharingScreen, ok := QueueconversationvideoeventtopicvideoMap["sharingScreen"].(bool); ok {
		o.SharingScreen = &SharingScreen
	}
    
	if PeerCount, ok := QueueconversationvideoeventtopicvideoMap["peerCount"].(map[string]interface{}); ok {
		PeerCountString, _ := json.Marshal(PeerCount)
		json.Unmarshal(PeerCountString, &o.PeerCount)
	}
	
	if Provider, ok := QueueconversationvideoeventtopicvideoMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := QueueconversationvideoeventtopicvideoMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := QueueconversationvideoeventtopicvideoMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if DisconnectType, ok := QueueconversationvideoeventtopicvideoMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if connectedTimeString, ok := QueueconversationvideoeventtopicvideoMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopicvideoMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Msids, ok := QueueconversationvideoeventtopicvideoMap["msids"].([]interface{}); ok {
		MsidsString, _ := json.Marshal(Msids)
		json.Unmarshal(MsidsString, &o.Msids)
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopicvideoMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopicvideoMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopicvideoMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicvideo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
