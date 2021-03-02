package platformclientv2
import (
	"encoding/json"
)

// Queueconversationcallbackeventtopicvoicemail
type Queueconversationcallbackeventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
