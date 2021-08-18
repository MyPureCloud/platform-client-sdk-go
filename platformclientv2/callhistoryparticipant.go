package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Callhistoryparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callhistoryparticipant

	
	StartTime := new(string)
	if u.StartTime != nil {
		
		*StartTime = timeutil.Strftime(u.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ExternalContact *Externalcontact `json:"externalContact,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		DidInteract *bool `json:"didInteract,omitempty"`
		
		SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		OutboundCampaign *Campaign `json:"outboundCampaign,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Address: u.Address,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		Purpose: u.Purpose,
		
		Direction: u.Direction,
		
		Ani: u.Ani,
		
		Dnis: u.Dnis,
		
		User: u.User,
		
		Queue: u.Queue,
		
		Group: u.Group,
		
		DisconnectType: u.DisconnectType,
		
		ExternalContact: u.ExternalContact,
		
		ExternalOrganization: u.ExternalOrganization,
		
		DidInteract: u.DidInteract,
		
		SipResponseCodes: u.SipResponseCodes,
		
		FlaggedReason: u.FlaggedReason,
		
		OutboundCampaign: u.OutboundCampaign,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callhistoryparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
