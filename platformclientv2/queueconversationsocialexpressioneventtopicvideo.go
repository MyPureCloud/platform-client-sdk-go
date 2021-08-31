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


	// Self
	Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Context
	Context *string `json:"context,omitempty"`


	// AudioMuted
	AudioMuted *bool `json:"audioMuted,omitempty"`


	// VideoMuted
	VideoMuted *bool `json:"videoMuted,omitempty"`


	// SharingScreen
	SharingScreen *bool `json:"sharingScreen,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Msids
	Msids *[]string `json:"msids,omitempty"`


	// Wrapup
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

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
		
		Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		AudioMuted *bool `json:"audioMuted,omitempty"`
		
		VideoMuted *bool `json:"videoMuted,omitempty"`
		
		SharingScreen *bool `json:"sharingScreen,omitempty"`
		
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
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Self: o.Self,
		
		Id: o.Id,
		
		Context: o.Context,
		
		AudioMuted: o.AudioMuted,
		
		VideoMuted: o.VideoMuted,
		
		SharingScreen: o.SharingScreen,
		
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
