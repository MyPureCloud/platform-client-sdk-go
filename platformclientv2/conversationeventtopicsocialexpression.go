package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicsocialexpression
type Conversationeventtopicsocialexpression struct { 
	// State
	State *string `json:"state,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// SocialMediaId
	SocialMediaId *string `json:"socialMediaId,omitempty"`


	// SocialMediaHub
	SocialMediaHub *string `json:"socialMediaHub,omitempty"`


	// SocialUserName
	SocialUserName *string `json:"socialUserName,omitempty"`


	// PreviewText
	PreviewText *string `json:"previewText,omitempty"`


	// RecordingId
	RecordingId *string `json:"recordingId,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Wrapup
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicsocialexpression) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
