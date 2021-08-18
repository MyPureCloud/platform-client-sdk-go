package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsparticipantwithoutattributes
type Analyticsparticipantwithoutattributes struct { 
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

}

func (u *Analyticsparticipantwithoutattributes) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsparticipantwithoutattributes

	

	return json.Marshal(&struct { 
		ExternalContactId *string `json:"externalContactId,omitempty"`
		
		ExternalOrganizationId *string `json:"externalOrganizationId,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Sessions *[]Analyticssession `json:"sessions,omitempty"`
		*Alias
	}{ 
		ExternalContactId: u.ExternalContactId,
		
		ExternalOrganizationId: u.ExternalOrganizationId,
		
		FlaggedReason: u.FlaggedReason,
		
		ParticipantId: u.ParticipantId,
		
		ParticipantName: u.ParticipantName,
		
		Purpose: u.Purpose,
		
		TeamId: u.TeamId,
		
		UserId: u.UserId,
		
		Sessions: u.Sessions,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Analyticsparticipantwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
