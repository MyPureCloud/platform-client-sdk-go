package platformclientv2
import (
	"encoding/json"
)

// Bulkcontactsrequest
type Bulkcontactsrequest struct { 
	// Entities
	Entities *[]Externalcontact `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkcontactsrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
