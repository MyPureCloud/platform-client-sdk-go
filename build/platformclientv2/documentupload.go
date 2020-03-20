package platformclientv2
import (
	"encoding/json"
)

// Documentupload
type Documentupload struct { 
	// Name - The name of the document
	Name *string `json:"name,omitempty"`


	// Workspace - The workspace the document will be uploaded to
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// TagIds
	TagIds *[]string `json:"tagIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentupload) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
