package platformclientv2
import (
	"encoding/json"
)

// Documentupdate
type Documentupdate struct { 
	// ChangeNumber
	ChangeNumber *int32 `json:"changeNumber,omitempty"`


	// Name - The name of the document
	Name *string `json:"name,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// AddTags
	AddTags *[]string `json:"addTags,omitempty"`


	// RemoveTags
	RemoveTags *[]string `json:"removeTags,omitempty"`


	// AddTagIds
	AddTagIds *[]string `json:"addTagIds,omitempty"`


	// RemoveTagIds
	RemoveTagIds *[]string `json:"removeTagIds,omitempty"`


	// UpdateAttributes
	UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`


	// RemoveAttributes
	RemoveAttributes *[]string `json:"removeAttributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Documentupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
