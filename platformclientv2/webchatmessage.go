package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatmessage
type Webchatmessage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Conversation - The identifier of the conversation
	Conversation *Webchatconversation `json:"conversation,omitempty"`


	// Sender - The member who sent the message
	Sender *Webchatmemberinfo `json:"sender,omitempty"`


	// Body - The message body.
	Body *string `json:"body,omitempty"`


	// BodyType - The purpose of the message within the conversation, such as a standard text entry versus a greeting.
	BodyType *string `json:"bodyType,omitempty"`


	// Timestamp - The timestamp of the message, in ISO-8601 format
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Webchatmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatmessage

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Conversation *Webchatconversation `json:"conversation,omitempty"`
		
		Sender *Webchatmemberinfo `json:"sender,omitempty"`
		
		Body *string `json:"body,omitempty"`
		
		BodyType *string `json:"bodyType,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Conversation: u.Conversation,
		
		Sender: u.Sender,
		
		Body: u.Body,
		
		BodyType: u.BodyType,
		
		Timestamp: Timestamp,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
