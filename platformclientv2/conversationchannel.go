package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationchannel
type Conversationchannel struct { 
	// VarType - The type or category of this channel.
	VarType *string `json:"type,omitempty"`


	// MessageType - Message type for messaging conversations.
	MessageType *string `json:"messageType,omitempty"`


	// Platform - The source provider for the conversation (e.g. Edge, PureCloud Messaging, PureCloud Email).
	Platform *string `json:"platform,omitempty"`

}

func (o *Conversationchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationchannel
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		Platform *string `json:"platform,omitempty"`
		*Alias
	}{ 
		VarType: o.VarType,
		
		MessageType: o.MessageType,
		
		Platform: o.Platform,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationchannel) UnmarshalJSON(b []byte) error {
	var ConversationchannelMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationchannelMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConversationchannelMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if MessageType, ok := ConversationchannelMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if Platform, ok := ConversationchannelMap["platform"].(string); ok {
		o.Platform = &Platform
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
