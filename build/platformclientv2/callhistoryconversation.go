package platformclientv2
import (
	"time"
	"encoding/json"
)

// Callhistoryconversation
type Callhistoryconversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Participants - The list of participants involved in the conversation.
	Participants *[]Callhistoryparticipant `json:"participants,omitempty"`


	// Direction - The direction of the call relating to the current user
	Direction *string `json:"direction,omitempty"`


	// WentToVoicemail - Did the call end in the current user's voicemail
	WentToVoicemail *bool `json:"wentToVoicemail,omitempty"`


	// MissedCall - Did the user not answer this conversation
	MissedCall *bool `json:"missedCall,omitempty"`


	// StartTime - The time the user joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	StartTime *time.Time `json:"startTime,omitempty"`


	// WasConference - Was this conversation a conference
	WasConference *bool `json:"wasConference,omitempty"`


	// WasCallback - Was this conversation a callback
	WasCallback *bool `json:"wasCallback,omitempty"`


	// HadScreenShare - Did this conversation have a screen share session
	HadScreenShare *bool `json:"hadScreenShare,omitempty"`


	// HadCobrowse - Did this conversation have a cobrowse session
	HadCobrowse *bool `json:"hadCobrowse,omitempty"`


	// WasOutboundCampaign - Was this conversation associated with an outbound campaign
	WasOutboundCampaign *bool `json:"wasOutboundCampaign,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callhistoryconversation) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
