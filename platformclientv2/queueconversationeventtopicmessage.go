package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationeventtopicmessage
type Queueconversationeventtopicmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`


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


	// ToAddress
	ToAddress *Queueconversationeventtopicaddress `json:"toAddress,omitempty"`


	// FromAddress
	FromAddress *Queueconversationeventtopicaddress `json:"fromAddress,omitempty"`


	// Messages
	Messages *[]Queueconversationeventtopicmessagedetails `json:"messages,omitempty"`


	// MessagesTranscriptUri
	MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// RecipientCountry
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType
	RecipientType *string `json:"recipientType,omitempty"`


	// Wrapup
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessage) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
