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

func (o *Callhistoryconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Callhistoryconversation
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		Direction: o.Direction,
		
		WentToVoicemail: o.WentToVoicemail,
		
		MissedCall: o.MissedCall,
		
		StartTime: StartTime,
		
		WasConference: o.WasConference,
		
		WasCallback: o.WasCallback,
		
		HadScreenShare: o.HadScreenShare,
		
		HadCobrowse: o.HadCobrowse,
		
		WasOutboundCampaign: o.WasOutboundCampaign,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Callhistoryconversation) UnmarshalJSON(b []byte) error {
	var CallhistoryconversationMap map[string]interface{}
	err := json.Unmarshal(b, &CallhistoryconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallhistoryconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CallhistoryconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := CallhistoryconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if Direction, ok := CallhistoryconversationMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if WentToVoicemail, ok := CallhistoryconversationMap["wentToVoicemail"].(bool); ok {
		o.WentToVoicemail = &WentToVoicemail
	}
    
	if MissedCall, ok := CallhistoryconversationMap["missedCall"].(bool); ok {
		o.MissedCall = &MissedCall
	}
    
	if startTimeString, ok := CallhistoryconversationMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if WasConference, ok := CallhistoryconversationMap["wasConference"].(bool); ok {
		o.WasConference = &WasConference
	}
    
	if WasCallback, ok := CallhistoryconversationMap["wasCallback"].(bool); ok {
		o.WasCallback = &WasCallback
	}
    
	if HadScreenShare, ok := CallhistoryconversationMap["hadScreenShare"].(bool); ok {
		o.HadScreenShare = &HadScreenShare
	}
    
	if HadCobrowse, ok := CallhistoryconversationMap["hadCobrowse"].(bool); ok {
		o.HadCobrowse = &HadCobrowse
	}
    
	if WasOutboundCampaign, ok := CallhistoryconversationMap["wasOutboundCampaign"].(bool); ok {
		o.WasOutboundCampaign = &WasOutboundCampaign
	}
    
	if SelfUri, ok := CallhistoryconversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Callhistoryconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
