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

func (o *Webchatmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatmessage
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Conversation: o.Conversation,
		
		Sender: o.Sender,
		
		Body: o.Body,
		
		BodyType: o.BodyType,
		
		Timestamp: Timestamp,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatmessage) UnmarshalJSON(b []byte) error {
	var WebchatmessageMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebchatmessageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Conversation, ok := WebchatmessageMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Sender, ok := WebchatmessageMap["sender"].(map[string]interface{}); ok {
		SenderString, _ := json.Marshal(Sender)
		json.Unmarshal(SenderString, &o.Sender)
	}
	
	if Body, ok := WebchatmessageMap["body"].(string); ok {
		o.Body = &Body
	}
    
	if BodyType, ok := WebchatmessageMap["bodyType"].(string); ok {
		o.BodyType = &BodyType
	}
    
	if timestampString, ok := WebchatmessageMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if SelfUri, ok := WebchatmessageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
