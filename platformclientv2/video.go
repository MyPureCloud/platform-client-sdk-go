package platformclientv2
import (
	"time"
	"encoding/json"
)

// Video
type Video struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


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
	PeerCount *int32 `json:"peerCount,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Provider - The source provider for the video.
	Provider *string `json:"provider,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// Msids - List of media stream ids
	Msids *[]string `json:"msids,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`

}

// String returns a JSON representation of the model
func (o *Video) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
