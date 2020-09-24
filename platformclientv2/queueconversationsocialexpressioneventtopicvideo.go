package platformclientv2
import (
	"time"
	"encoding/json"
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
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicvideo) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
