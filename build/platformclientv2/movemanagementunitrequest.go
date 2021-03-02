package platformclientv2
import (
	"encoding/json"
)

// Movemanagementunitrequest
type Movemanagementunitrequest struct { 
	// BusinessUnitId - The ID of the business unit to which to move the management unit
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Movemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
