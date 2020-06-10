package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationvideoeventtopiccallback
type Queueconversationvideoeventtopiccallback struct { 
	// State
	State *string `json:"state,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// DialerPreview
	DialerPreview *Queueconversationvideoeventtopicdialerpreview `json:"dialerPreview,omitempty"`


	// Voicemail
	Voicemail *Queueconversationvideoeventtopicvoicemail `json:"voicemail,omitempty"`


	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// ExternalCampaign
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// SkipEnabled
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// TimeoutSeconds
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// CallbackScheduledTime
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// AutomatedCallbackConfigId
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`


	// Wrapup
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopiccallback) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
