package platformclientv2
import (
	"encoding/json"
)

// Createsecuresession
type Createsecuresession struct { 
	// SourceParticipantId - requesting participant
	SourceParticipantId *string `json:"sourceParticipantId,omitempty"`


	// FlowId - the flow id to execute in the secure session
	FlowId *string `json:"flowId,omitempty"`


	// UserData - user data for the secure session
	UserData *string `json:"userData,omitempty"`


	// Disconnect - if true, disconnect the agent after creating the session
	Disconnect *bool `json:"disconnect,omitempty"`

}

// String returns a JSON representation of the model
func (o *Createsecuresession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
