package platformclientv2
import (
	"encoding/json"
)

// Sipdownloadresponse
type Sipdownloadresponse struct { 
	// DownloadId - unique id of the downloaded file
	DownloadId *string `json:"downloadId,omitempty"`


	// DocumentId - Document id of pcap file
	DocumentId *string `json:"documentId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Sipdownloadresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
