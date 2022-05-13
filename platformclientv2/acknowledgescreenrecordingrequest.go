package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Acknowledgescreenrecordingrequest
type Acknowledgescreenrecordingrequest struct { 
	// ParticipantJid
	ParticipantJid *string `json:"participantJid,omitempty"`


	// RoomId
	RoomId *string `json:"roomId,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

}

func (o *Acknowledgescreenrecordingrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acknowledgescreenrecordingrequest
	
	return json.Marshal(&struct { 
		ParticipantJid *string `json:"participantJid,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		*Alias
	}{ 
		ParticipantJid: o.ParticipantJid,
		
		RoomId: o.RoomId,
		
		ConversationId: o.ConversationId,
		Alias:    (*Alias)(o),
	})
}

func (o *Acknowledgescreenrecordingrequest) UnmarshalJSON(b []byte) error {
	var AcknowledgescreenrecordingrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AcknowledgescreenrecordingrequestMap)
	if err != nil {
		return err
	}
	
	if ParticipantJid, ok := AcknowledgescreenrecordingrequestMap["participantJid"].(string); ok {
		o.ParticipantJid = &ParticipantJid
	}
    
	if RoomId, ok := AcknowledgescreenrecordingrequestMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
    
	if ConversationId, ok := AcknowledgescreenrecordingrequestMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Acknowledgescreenrecordingrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
