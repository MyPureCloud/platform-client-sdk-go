package platformclientv2
import (
	"encoding/json"
)

// Responsetext - Contains information about the text associated with a response.
type Responsetext struct { 
	// Content - Response text content.
	Content *string `json:"content,omitempty"`


	// ContentType - Response text content type.
	ContentType *string `json:"contentType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Responsetext) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
