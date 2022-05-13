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

func (o *Sipsearchpublicrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Sipsearchpublicrequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		CallId: o.CallId,
		
		ToUser: o.ToUser,
		
		FromUser: o.FromUser,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Sipsearchpublicrequest) UnmarshalJSON(b []byte) error {
	var SipsearchpublicrequestMap map[string]interface{}
	err := json.Unmarshal(b, &SipsearchpublicrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SipsearchpublicrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SipsearchpublicrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if CallId, ok := SipsearchpublicrequestMap["callId"].(string); ok {
		o.CallId = &CallId
	}
    
	if ToUser, ok := SipsearchpublicrequestMap["toUser"].(string); ok {
		o.ToUser = &ToUser
	}
    
	if FromUser, ok := SipsearchpublicrequestMap["fromUser"].(string); ok {
		o.FromUser = &FromUser
	}
    
	if ConversationId, ok := SipsearchpublicrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := SipsearchpublicrequestMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if dateStartString, ok := SipsearchpublicrequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := SipsearchpublicrequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	
	if SelfUri, ok := SipsearchpublicrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Sipsearchpublicrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
