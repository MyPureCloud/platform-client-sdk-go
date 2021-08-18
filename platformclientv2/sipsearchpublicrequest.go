package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Sipsearchpublicrequest
type Sipsearchpublicrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// CallId - unique identification of the placed call
	CallId *string `json:"callId,omitempty"`


	// ToUser - SIP user to who the call was placed
	ToUser *string `json:"toUser,omitempty"`


	// FromUser - SIP user who placed the call
	FromUser *string `json:"fromUser,omitempty"`


	// ConversationId - Unique identification of the conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId - Unique identification of the participant
	ParticipantId *string `json:"participantId,omitempty"`


	// DateStart - Start date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStart *time.Time `json:"dateStart,omitempty"`


	// DateEnd - End date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnd *time.Time `json:"dateEnd,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Sipsearchpublicrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sipsearchpublicrequest

	
	DateStart := new(string)
	if u.DateStart != nil {
		
		*DateStart = timeutil.Strftime(u.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if u.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(u.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CallId *string `json:"callId,omitempty"`
		
		ToUser *string `json:"toUser,omitempty"`
		
		FromUser *string `json:"fromUser,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CallId: u.CallId,
		
		ToUser: u.ToUser,
		
		FromUser: u.FromUser,
		
		ConversationId: u.ConversationId,
		
		ParticipantId: u.ParticipantId,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Sipsearchpublicrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
