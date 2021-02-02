package platformclientv2
import (
	"encoding/json"
)

// Replaceresponse
type Replaceresponse struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// UploadStatus
	UploadStatus *Domainentityref `json:"uploadStatus,omitempty"`


	// UploadDestinationUri
	UploadDestinationUri *string `json:"uploadDestinationUri,omitempty"`


	// UploadMethod
	UploadMethod *string `json:"uploadMethod,omitempty"`

}

// String returns a JSON representation of the model
func (o *Replaceresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
