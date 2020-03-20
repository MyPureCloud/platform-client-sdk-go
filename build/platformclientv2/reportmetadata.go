package platformclientv2
import (
	"encoding/json"
)

// Reportmetadata
type Reportmetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Keywords
	Keywords *[]string `json:"keywords,omitempty"`


	// AvailableLocales
	AvailableLocales *[]string `json:"availableLocales,omitempty"`


	// Parameters
	Parameters *[]Parameter `json:"parameters,omitempty"`


	// ExampleUrl
	ExampleUrl *string `json:"exampleUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Reportmetadata) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
