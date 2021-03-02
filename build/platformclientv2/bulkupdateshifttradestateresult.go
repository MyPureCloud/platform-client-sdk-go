package platformclientv2
import (
	"encoding/json"
)

// Bulkupdateshifttradestateresult
type Bulkupdateshifttradestateresult struct { 
	// Entities
	Entities *[]Bulkupdateshifttradestateresultitem `json:"entities,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateshifttradestateresult) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
