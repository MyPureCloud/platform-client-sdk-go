package platformclientv2
import (
	"encoding/json"
)

// Intradayqueue
type Intradayqueue struct { 
	// Id - Queue ID
	Id *string `json:"id,omitempty"`


	// Name - Queue name
	Name *string `json:"name,omitempty"`


	// MediaTypes - The media types valid for this queue as defined by the service goal groups in this management unit
	MediaTypes *[]string `json:"mediaTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Intradayqueue) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
