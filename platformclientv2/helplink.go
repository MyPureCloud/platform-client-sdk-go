package platformclientv2
import (
	"encoding/json"
)

// Helplink - Link to a help or support resource
type Helplink struct { 
	// Uri - URI of the help resource
	Uri *string `json:"uri,omitempty"`


	// Title - Link text of the resource
	Title *string `json:"title,omitempty"`


	// Description - Description of the document or resource
	Description *string `json:"description,omitempty"`

}

// String returns a JSON representation of the model
func (o *Helplink) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
