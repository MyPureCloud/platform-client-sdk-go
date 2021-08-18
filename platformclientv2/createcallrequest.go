package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallrequest
type Createcallrequest struct { 
	// PhoneNumber - The phone number to dial.
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// CallerId - The caller id phone number for this outbound call.
	CallerId *string `json:"callerId,omitempty"`


	// CallerIdName - The caller id name for this outbound call.
	CallerIdName *string `json:"callerIdName,omitempty"`


	// CallFromQueueId - The queue ID to call on behalf of.
	CallFromQueueId *string `json:"callFromQueueId,omitempty"`


	// CallQueueId - The queue ID to call.
	CallQueueId *string `json:"callQueueId,omitempty"`


	// CallUserId - The user ID to call.
	CallUserId *string `json:"callUserId,omitempty"`


	// Priority - The priority to assign to this call (if calling a queue).
	Priority *int `json:"priority,omitempty"`


	// LanguageId - The language skill ID to use for routing this call (if calling a queue).
	LanguageId *string `json:"languageId,omitempty"`


	// RoutingSkillsIds - The skill ID's to use for routing this call (if calling a queue).
	RoutingSkillsIds *[]string `json:"routingSkillsIds,omitempty"`


	// ConversationIds - The list of existing call conversations to merge into a new ad-hoc conference.
	ConversationIds *[]string `json:"conversationIds,omitempty"`


	// Participants - The list of participants to call to create a new ad-hoc conference.
	Participants *[]Destination `json:"participants,omitempty"`


	// UuiData - User to User Information (UUI) data managed by SIP session application.
	UuiData *string `json:"uuiData,omitempty"`

}

func (u *Createcallrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcallrequest

	

	return json.Marshal(&struct { 
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		
		CallFromQueueId *string `json:"callFromQueueId,omitempty"`
		
		CallQueueId *string `json:"callQueueId,omitempty"`
		
		CallUserId *string `json:"callUserId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		RoutingSkillsIds *[]string `json:"routingSkillsIds,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		Participants *[]Destination `json:"participants,omitempty"`
		
		UuiData *string `json:"uuiData,omitempty"`
		*Alias
	}{ 
		PhoneNumber: u.PhoneNumber,
		
		CallerId: u.CallerId,
		
		CallerIdName: u.CallerIdName,
		
		CallFromQueueId: u.CallFromQueueId,
		
		CallQueueId: u.CallQueueId,
		
		CallUserId: u.CallUserId,
		
		Priority: u.Priority,
		
		LanguageId: u.LanguageId,
		
		RoutingSkillsIds: u.RoutingSkillsIds,
		
		ConversationIds: u.ConversationIds,
		
		Participants: u.Participants,
		
		UuiData: u.UuiData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createcallrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
