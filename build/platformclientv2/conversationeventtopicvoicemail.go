package platformclientv2
import (
	"encoding/json"
)

// Conversationeventtopicvoicemail
type Conversationeventtopicvoicemail struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// UploadStatus
	UploadStatus *string `json:"uploadStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopicvoicemail) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
