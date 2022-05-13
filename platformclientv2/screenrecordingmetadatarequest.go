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

func (o *Screenrecordingmetadatarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Screenrecordingmetadatarequest
	
	return json.Marshal(&struct { 
		ParticipantJid *string `json:"participantJid,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		MetaData *[]Screenrecordingmetadata `json:"metaData,omitempty"`
		*Alias
	}{ 
		ParticipantJid: o.ParticipantJid,
		
		RoomId: o.RoomId,
		
		MetaData: o.MetaData,
		Alias:    (*Alias)(o),
	})
}

func (o *Screenrecordingmetadatarequest) UnmarshalJSON(b []byte) error {
	var ScreenrecordingmetadatarequestMap map[string]interface{}
	err := json.Unmarshal(b, &ScreenrecordingmetadatarequestMap)
	if err != nil {
		return err
	}
	
	if ParticipantJid, ok := ScreenrecordingmetadatarequestMap["participantJid"].(string); ok {
		o.ParticipantJid = &ParticipantJid
	}
    
	if RoomId, ok := ScreenrecordingmetadatarequestMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
    
	if MetaData, ok := ScreenrecordingmetadatarequestMap["metaData"].([]interface{}); ok {
		MetaDataString, _ := json.Marshal(MetaData)
		json.Unmarshal(MetaDataString, &o.MetaData)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Screenrecordingmetadatarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
