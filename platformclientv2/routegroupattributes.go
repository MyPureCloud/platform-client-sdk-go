package platformclientv2
import (
	"encoding/json"
)

// Routegroupattributes - Attributes for the associated route group
type Routegroupattributes struct { 
	// Queue - The queue to which the associated route group applies
	Queue *Queuereference `json:"queue,omitempty"`


	// MediaType - The media type to which the associated route group applies
	MediaType *string `json:"mediaType,omitempty"`


	// Language - The language to which the associated route group applies
	Language *Languagereference `json:"language,omitempty"`


	// Skills - The skill set to which the associated route group applies
	Skills *[]Routingskillreference `json:"skills,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routegroupattributes) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
