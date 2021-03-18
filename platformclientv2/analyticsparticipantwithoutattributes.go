package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsparticipantwithoutattributes
type Analyticsparticipantwithoutattributes struct { 
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


	// TeamId - The team id the user is a member of
	TeamId *string `json:"teamId,omitempty"`


	// Sessions - List of sessions associated to this participant
	Sessions *[]Analyticssession `json:"sessions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsparticipantwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
