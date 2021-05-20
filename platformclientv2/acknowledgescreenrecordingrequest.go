package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Acknowledgescreenrecordingrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
