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

func (u *Acknowledgescreenrecordingrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Acknowledgescreenrecordingrequest

	

	return json.Marshal(&struct { 
		ParticipantJid *string `json:"participantJid,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		*Alias
	}{ 
		ParticipantJid: u.ParticipantJid,
		
		RoomId: u.RoomId,
		
		ConversationId: u.ConversationId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Acknowledgescreenrecordingrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
