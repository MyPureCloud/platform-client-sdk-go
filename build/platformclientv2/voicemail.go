package platformclientv2
import (
	"encoding/json"
)

// Voicemail
type Voicemail struct { 
	// Id - The voicemail id
	Id *string `json:"id,omitempty"`


	// UploadStatus - current state of the voicemail upload
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Voicemail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
