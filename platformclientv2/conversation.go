package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversation
type Conversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ExternalTag - The external tag associated with the conversation.
	ExternalTag *string `json:"externalTag,omitempty"`


	// StartTime - The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// EndTime - The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// Address - The address of the conversation as seen from an external participant. For phone calls this will be the DNIS for inbound calls and the ANI for outbound calls. For other media types this will be the address of the destination participant for inbound and the address of the initiating participant for outbound.
	Address *string `json:"address,omitempty"`


	// Participants - The list of all participants in the conversation.
	Participants *[]Participant `json:"participants,omitempty"`


	// ConversationIds - A list of conversations to merge into this conversation to create a conference. This field is null except when being used to create a conference.
	ConversationIds *[]string `json:"conversationIds,omitempty"`


	// MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants *int `json:"maxParticipants,omitempty"`


	// RecordingState - On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings; otherwise indicates state of conversation recording.
	RecordingState *string `json:"recordingState,omitempty"`


	// State - The conversation's state
	State *string `json:"state,omitempty"`


	// Divisions - Identifiers of divisions associated with this conversation
	Divisions *[]Conversationdivisionmembership `json:"divisions,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Conversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversation

	
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
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Participants *[]Participant `json:"participants,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Divisions *[]Conversationdivisionmembership `json:"divisions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ExternalTag: u.ExternalTag,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		Address: u.Address,
		
		Participants: u.Participants,
		
		ConversationIds: u.ConversationIds,
		
		MaxParticipants: u.MaxParticipants,
		
		RecordingState: u.RecordingState,
		
		State: u.State,
		
		Divisions: u.Divisions,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
