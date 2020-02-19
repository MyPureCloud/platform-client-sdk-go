package platformclientv2
import (
	"encoding/json"
)

// Keyrotationschedule
type Keyrotationschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Period - Value to set schedule to
	Period *string `json:"period,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Keyrotationschedule) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
