package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatconversationresponse
type Createwebchatconversationresponse struct { 
	// Id - Chat Conversation identifier
	Id *string `json:"id,omitempty"`


	// Jwt - The JWT that you can use to identify subsequent calls on this conversation
	Jwt *string `json:"jwt,omitempty"`


	// EventStreamUri - The URI which provides the conversation event stream.
	EventStreamUri *string `json:"eventStreamUri,omitempty"`


	// Member - Chat Member
	Member *Webchatmemberinfo `json:"member,omitempty"`

}

func (o *Createwebchatconversationresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createwebchatconversationresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Jwt *string `json:"jwt,omitempty"`
		
		EventStreamUri *string `json:"eventStreamUri,omitempty"`
		
		Member *Webchatmemberinfo `json:"member,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Jwt: o.Jwt,
		
		EventStreamUri: o.EventStreamUri,
		
		Member: o.Member,
		Alias:    (*Alias)(o),
	})
}

func (o *Createwebchatconversationresponse) UnmarshalJSON(b []byte) error {
	var CreatewebchatconversationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatconversationresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreatewebchatconversationresponseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Jwt, ok := CreatewebchatconversationresponseMap["jwt"].(string); ok {
		o.Jwt = &Jwt
	}
	
	if EventStreamUri, ok := CreatewebchatconversationresponseMap["eventStreamUri"].(string); ok {
		o.EventStreamUri = &EventStreamUri
	}
	
	if Member, ok := CreatewebchatconversationresponseMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatconversationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
