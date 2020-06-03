package platformclientv2
import (
	"encoding/json"
)

// Nluutterancesegment
type Nluutterancesegment struct { 
	// Text - The text of the segment.
	Text *string `json:"text,omitempty"`


	// Entity - The entity annotation of the segment.
	Entity *Namedentityannotation `json:"entity,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nluutterancesegment) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
