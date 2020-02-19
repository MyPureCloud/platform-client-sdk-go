package platformclientv2
import (
	"encoding/json"
)

// Analyticsparticipant
type Analyticsparticipant struct { 
	// ParticipantId - Unique identifier for the participant
	ParticipantId *string `json:"participantId,omitempty"`


	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`


	// UserId - If a user, then this will be the unique identifier for the user
	UserId *string `json:"userId,omitempty"`


	// Purpose - The participant's purpose
	Purpose *string `json:"purpose,omitempty"`


	// ExternalContactId - External Contact Identifier
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ExternalOrganizationId - External Organization Identifier
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// FlaggedReason - Reason for which participant flagged conversation
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// Sessions - List of sessions associated to this participant
	Sessions *[]Analyticssession `json:"sessions,omitempty"`


	// Attributes - List of attributes associated to this participant
	Attributes *map[string]string `json:"attributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
