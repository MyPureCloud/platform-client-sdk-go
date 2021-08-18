package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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


	// StartTime - The time the user joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

func (u *Callhistoryconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callhistoryconversation

	
	StartTime := new(string)
	if u.StartTime != nil {
		
		*StartTime = timeutil.Strftime(u.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Callhistoryparticipant `json:"participants,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		WentToVoicemail *bool `json:"wentToVoicemail,omitempty"`
		
		MissedCall *bool `json:"missedCall,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		WasConference *bool `json:"wasConference,omitempty"`
		
		WasCallback *bool `json:"wasCallback,omitempty"`
		
		HadScreenShare *bool `json:"hadScreenShare,omitempty"`
		
		HadCobrowse *bool `json:"hadCobrowse,omitempty"`
		
		WasOutboundCampaign *bool `json:"wasOutboundCampaign,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Participants: u.Participants,
		
		Direction: u.Direction,
		
		WentToVoicemail: u.WentToVoicemail,
		
		MissedCall: u.MissedCall,
		
		StartTime: StartTime,
		
		WasConference: u.WasConference,
		
		WasCallback: u.WasCallback,
		
		HadScreenShare: u.HadScreenShare,
		
		HadCobrowse: u.HadCobrowse,
		
		WasOutboundCampaign: u.WasOutboundCampaign,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Callhistoryconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
