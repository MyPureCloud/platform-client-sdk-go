package platformclientv2
import (
	"encoding/json"
)

// Partialuploadresponse
type Partialuploadresponse struct { 
	// Id - The reference id for a partial import request
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Partialuploadresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
