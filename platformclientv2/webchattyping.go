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

func (o *Webchattyping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchattyping
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Conversation: o.Conversation,
		
		Sender: o.Sender,
		
		Timestamp: Timestamp,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchattyping) UnmarshalJSON(b []byte) error {
	var WebchattypingMap map[string]interface{}
	err := json.Unmarshal(b, &WebchattypingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchattypingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Conversation, ok := WebchattypingMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Sender, ok := WebchattypingMap["sender"].(map[string]interface{}); ok {
		SenderString, _ := json.Marshal(Sender)
		json.Unmarshal(SenderString, &o.Sender)
	}
	
	if timestampString, ok := WebchattypingMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchattyping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
