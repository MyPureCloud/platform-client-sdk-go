package platformclientv2
import (
	"encoding/json"
)

// Bulkorganizationsrequest
type Bulkorganizationsrequest struct { 
	// Entities
	Entities *[]Externalorganization `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkorganizationsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
