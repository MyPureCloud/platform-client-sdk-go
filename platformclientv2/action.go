package platformclientv2
import (
	"encoding/json"
)

// Action
type Action struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// IntegrationId - The ID of the integration for which this action is associated
	IntegrationId *string `json:"integrationId,omitempty"`


	// Category - Category of Action
	Category *string `json:"category,omitempty"`


	// Contract - Action contract
	Contract *Actioncontract `json:"contract,omitempty"`


	// Version - Version of this action
	Version *int32 `json:"version,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Action) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
