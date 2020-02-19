package platformclientv2
import (
	"encoding/json"
)

// Postactioninput - Definition of an Action to be created or updated.
type Postactioninput struct { 
	// Category - Category of action
	Category *string `json:"category,omitempty"`


	// Name - Name of action
	Name *string `json:"name,omitempty"`


	// IntegrationId - The ID of the integration this action is associated to
	IntegrationId *string `json:"integrationId,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Contract - Action contract
	Contract *Actioncontractinput `json:"contract,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`

}

// String returns a JSON representation of the model
func (o *Postactioninput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
