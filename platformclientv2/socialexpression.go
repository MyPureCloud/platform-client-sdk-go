package platformclientv2
import (
	"time"
	"encoding/json"
)

// Socialexpression
type Socialexpression struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// SocialMediaId - A globally unique identifier for the social media.
	SocialMediaId *string `json:"socialMediaId,omitempty"`


	// SocialMediaHub - The social network of the communication
	SocialMediaHub *string `json:"socialMediaHub,omitempty"`


	// SocialUserName - The user name for the communication.
	SocialUserName *string `json:"socialUserName,omitempty"`


	// PreviewText - The text preview of the communication contents
	PreviewText *string `json:"previewText,omitempty"`


	// RecordingId - A globally unique identifier for the recording associated with this chat.
	RecordingId *string `json:"recordingId,omitempty"`


	// Segments - The time line of the participant's chat, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`


	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the chat was placed on hold in the cloud clock if the chat is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Provider - The source provider for the social expression.
	Provider *string `json:"provider,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

// String returns a JSON representation of the model
func (o *Socialexpression) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
