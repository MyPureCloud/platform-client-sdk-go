package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenrecordingmetadatarequest
type Screenrecordingmetadatarequest struct { 
	// ParticipantJid
	ParticipantJid *string `json:"participantJid,omitempty"`


	// RoomId
	RoomId *string `json:"roomId,omitempty"`


	// MetaData
	MetaData *[]Screenrecordingmetadata `json:"metaData,omitempty"`

}

func (u *Screenrecordingmetadatarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Screenrecordingmetadatarequest

	

	return json.Marshal(&struct { 
		ParticipantJid *string `json:"participantJid,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		MetaData *[]Screenrecordingmetadata `json:"metaData,omitempty"`
		*Alias
	}{ 
		ParticipantJid: u.ParticipantJid,
		
		RoomId: u.RoomId,
		
		MetaData: u.MetaData,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
