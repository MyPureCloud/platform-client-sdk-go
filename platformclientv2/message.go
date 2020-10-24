package platformclientv2
import (
	"time"
	"encoding/json"
)

// Message
type Message struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`


	// Segments - The time line of the participant's message, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`


	// Direction - The direction of the message.
	Direction *string `json:"direction,omitempty"`


	// RecordingId - A globally unique identifier for the recording associated with this message.
	RecordingId *string `json:"recordingId,omitempty"`


	// ErrorInfo
	ErrorInfo *Errorbody `json:"errorInfo,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the message was placed on hold in the cloud clock if the message is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Provider - The source provider for the message.
	Provider *string `json:"provider,omitempty"`


	// VarType - Indicates the type of message platform from which the message originated.
	VarType *string `json:"type,omitempty"`


	// RecipientCountry - Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType - The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
	RecipientType *string `json:"recipientType,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// ToAddress - Address and name data for a call endpoint.
	ToAddress *Address `json:"toAddress,omitempty"`


	// FromAddress - Address and name data for a call endpoint.
	FromAddress *Address `json:"fromAddress,omitempty"`


	// Messages - The messages sent on this communication channel.
	Messages *[]Messagedetails `json:"messages,omitempty"`


	// JourneyContext - A subset of the Journey System's data relevant to a part of a conversation (for external linkage and internal usage/context).
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

// String returns a JSON representation of the model
func (o *Message) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
