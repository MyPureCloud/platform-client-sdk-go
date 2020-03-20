package platformclientv2
import (
	"encoding/json"
)

// Setuuidatarequest
type Setuuidatarequest struct { 
	// UuiData - The value of the uuiData to set.
	UuiData *string `json:"uuiData,omitempty"`

}

// String returns a JSON representation of the model
func (o *Setuuidatarequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
