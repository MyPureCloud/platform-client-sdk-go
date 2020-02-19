package platformclientv2
import (
	"encoding/json"
)

// Destination
type Destination struct { 
	// Address - Address or phone number.
	Address *string `json:"address,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// UserId
	UserId *string `json:"userId,omitempty"`


	// QueueId
	QueueId *string `json:"queueId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Destination) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
