package platformclientv2
import (
	"encoding/json"
)

// Uploadurlresponse
type Uploadurlresponse struct { 
	// Url - Presigned URL to PUT the file to
	Url *string `json:"url,omitempty"`


	// UploadKey - Key that identifies the file in the storage including the file name
	UploadKey *string `json:"uploadKey,omitempty"`


	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`

}

// String returns a JSON representation of the model
func (o *Uploadurlresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
