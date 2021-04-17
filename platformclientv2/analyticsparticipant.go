package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsparticipant
type Analyticsparticipant struct { 
	// ExternalContactId - External contact identifier
	ExternalContactId *string `json:"externalContactId,omitempty"`


	// ExternalOrganizationId - External organization identifier
	ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`


	// FlaggedReason - Reason for which participant flagged conversation
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// ParticipantId - Unique identifier for the participant
	ParticipantId *string `json:"participantId,omitempty"`


	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`


	// Purpose - The participant's purpose
	Purpose *string `json:"purpose,omitempty"`


	// TeamId - The team ID the user is a member of
	TeamId *string `json:"teamId,omitempty"`


	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`


	// Sessions - List of sessions associated to this participant
	Sessions *[]Analyticssession `json:"sessions,omitempty"`


	// Attributes - List of attributes associated to this participant
	Attributes *map[string]string `json:"attributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
