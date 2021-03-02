package platformclientv2
import (
	"encoding/json"
)

// Architectflownotificationflownotification
type Architectflownotificationflownotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// CheckedInVersion
	CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`


	// SavedVersion
	SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`


	// PublishedVersion
	PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationflownotification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
