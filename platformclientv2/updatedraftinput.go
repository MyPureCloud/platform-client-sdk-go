package platformclientv2
import (
	"encoding/json"
)

// Updatedraftinput - Definition of an Action Draft to be created or updated.
type Updatedraftinput struct { 
	// Category - Category of action
	Category *string `json:"category,omitempty"`


	// Name - Name of action
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Contract - Action contract
	Contract *Actioncontractinput `json:"contract,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Version - Version of current Draft
	Version *int `json:"version,omitempty"`

}

// String returns a JSON representation of the model
func (o *Updatedraftinput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
