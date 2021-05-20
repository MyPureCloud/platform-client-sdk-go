package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Screenrecordingmetadatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
