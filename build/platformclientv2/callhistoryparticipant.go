package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Callhistoryparticipant
type Callhistoryparticipant struct { 
	// Id - The unique participant ID.
	Id *string `json:"id,omitempty"`


	// Name - The display friendly name of the participant.
	Name *string `json:"name,omitempty"`


	// Address - The participant address.
	Address *string `json:"address,omitempty"`


	// StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose *string `json:"purpose,omitempty"`


	// Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
	Direction *string `json:"direction,omitempty"`


	// Ani - The call ANI.
	Ani *string `json:"ani,omitempty"`


	// Dnis - The call DNIS.
	Dnis *string `json:"dnis,omitempty"`


	// User - The PureCloud user for this participant.
	User *User `json:"user,omitempty"`


	// Queue - The PureCloud queue for this participant.
	Queue *Queue `json:"queue,omitempty"`


	// Group - The group involved in the group ring call.
	Group *Group `json:"group,omitempty"`


	// DisconnectType - The reason the participant was disconnected from the conversation.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ExternalContact - The PureCloud external contact
	ExternalContact *Externalcontact `json:"externalContact,omitempty"`


	// ExternalOrganization - The PureCloud external organization
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`


	// DidInteract - Indicates whether the contact ever connected
	DidInteract *bool `json:"didInteract,omitempty"`


	// SipResponseCodes - Indicates SIP Response codes associated with the participant
	SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`


	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// OutboundCampaign - The outbound campaign associated with the participant
	OutboundCampaign *Campaign `json:"outboundCampaign,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callhistoryparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
