package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicvideo
type Queueconversationsocialexpressioneventtopicvideo struct { 
	// State
	State *string `json:"state,omitempty"`


	// InitialState
	InitialState *string `json:"initialState,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`


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
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork - A communication's after-call work data.
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicvideo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicvideo
	
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
		
		Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`
		
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
		
		Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
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
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicvideo) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicvideoMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicvideoMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationsocialexpressioneventtopicvideoMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := QueueconversationsocialexpressioneventtopicvideoMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Self, ok := QueueconversationsocialexpressioneventtopicvideoMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicvideoMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Context, ok := QueueconversationsocialexpressioneventtopicvideoMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if AudioMuted, ok := QueueconversationsocialexpressioneventtopicvideoMap["audioMuted"].(bool); ok {
		o.AudioMuted = &AudioMuted
	}
    
	if VideoMuted, ok := QueueconversationsocialexpressioneventtopicvideoMap["videoMuted"].(bool); ok {
		o.VideoMuted = &VideoMuted
	}
    
	if SharingScreen, ok := QueueconversationsocialexpressioneventtopicvideoMap["sharingScreen"].(bool); ok {
		o.SharingScreen = &SharingScreen
	}
    
	if PeerCount, ok := QueueconversationsocialexpressioneventtopicvideoMap["peerCount"].(map[string]interface{}); ok {
		PeerCountString, _ := json.Marshal(PeerCount)
		json.Unmarshal(PeerCountString, &o.PeerCount)
	}
	
	if Provider, ok := QueueconversationsocialexpressioneventtopicvideoMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ScriptId, ok := QueueconversationsocialexpressioneventtopicvideoMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if PeerId, ok := QueueconversationsocialexpressioneventtopicvideoMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if DisconnectType, ok := QueueconversationsocialexpressioneventtopicvideoMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if connectedTimeString, ok := QueueconversationsocialexpressioneventtopicvideoMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationsocialexpressioneventtopicvideoMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Msids, ok := QueueconversationsocialexpressioneventtopicvideoMap["msids"].([]interface{}); ok {
		MsidsString, _ := json.Marshal(Msids)
		json.Unmarshal(MsidsString, &o.Msids)
	}
	
	if Wrapup, ok := QueueconversationsocialexpressioneventtopicvideoMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationsocialexpressioneventtopicvideoMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationsocialexpressioneventtopicvideoMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicvideoMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicvideo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
