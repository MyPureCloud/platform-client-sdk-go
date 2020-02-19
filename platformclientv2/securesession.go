package platformclientv2
import (
	"encoding/json"
)

// Securesession
type Securesession struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Flow - The flow to execute securely
	Flow *Domainentityref `json:"flow,omitempty"`


	// UserData - Customer-provided data
	UserData *string `json:"userData,omitempty"`


	// State - The current state of a secure session
	State *string `json:"state,omitempty"`


	// SourceParticipantId - Unique identifier for the participant initiating the secure session.
	SourceParticipantId *string `json:"sourceParticipantId,omitempty"`


	// Disconnect - If true, disconnect the agent after creating the session
	Disconnect *bool `json:"disconnect,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Securesession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
