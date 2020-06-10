package platformclientv2
import (
	"encoding/json"
)

// Coachingannotationcreaterequest
type Coachingannotationcreaterequest struct { 
	// Text - The text of the annotation.
	Text *string `json:"text,omitempty"`


	// AccessType - Determines the permissions required to view this item.
	AccessType *string `json:"accessType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Coachingannotationcreaterequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
