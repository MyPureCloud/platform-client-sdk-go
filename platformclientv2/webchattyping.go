package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchattyping
type Webchattyping struct { 
	// Id - The event identifier of this typing indicator event (useful to guard against event re-delivery
	Id *string `json:"id,omitempty"`


	// Conversation - The identifier of the conversation
	Conversation *Webchatconversation `json:"conversation,omitempty"`


	// Sender - The member who sent the message
	Sender *Webchatmemberinfo `json:"sender,omitempty"`


	// Timestamp - The timestamp of the message, in ISO-8601 format
	Timestamp *time.Time `json:"timestamp,omitempty"`

}

func (u *Webchattyping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchattyping

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Conversation *Webchatconversation `json:"conversation,omitempty"`
		
		Sender *Webchatmemberinfo `json:"sender,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Conversation: u.Conversation,
		
		Sender: u.Sender,
		
		Timestamp: Timestamp,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webchattyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
