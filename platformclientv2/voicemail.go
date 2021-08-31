package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Voicemail
type Voicemail struct { 
	// Id - The voicemail id
	Id *string `json:"id,omitempty"`


	// UploadStatus - current state of the voicemail upload
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

func (o *Voicemail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Voicemail
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		UploadStatus *string `json:"uploadStatus,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		UploadStatus: o.UploadStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Voicemail) UnmarshalJSON(b []byte) error {
	var VoicemailMap map[string]interface{}
	err := json.Unmarshal(b, &VoicemailMap)
	if err != nil {
		return err
	}
	
	if Id, ok := VoicemailMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if UploadStatus, ok := VoicemailMap["uploadStatus"].(string); ok {
		o.UploadStatus = &UploadStatus
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Voicemail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
