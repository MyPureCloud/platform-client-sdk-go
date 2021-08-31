package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (o *Analyticsparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticsparticipant
	
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
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		*Alias
	}{ 
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		FlaggedReason: o.FlaggedReason,
		
		ParticipantId: o.ParticipantId,
		
		ParticipantName: o.ParticipantName,
		
		Purpose: o.Purpose,
		
		TeamId: o.TeamId,
		
		UserId: o.UserId,
		
		Sessions: o.Sessions,
		
		Attributes: o.Attributes,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsparticipant) UnmarshalJSON(b []byte) error {
	var AnalyticsparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsparticipantMap)
	if err != nil {
		return err
	}
	
	if ExternalContactId, ok := AnalyticsparticipantMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
	
	if ExternalOrganizationId, ok := AnalyticsparticipantMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
	
	if FlaggedReason, ok := AnalyticsparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if ParticipantId, ok := AnalyticsparticipantMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if ParticipantName, ok := AnalyticsparticipantMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
	
	if Purpose, ok := AnalyticsparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if TeamId, ok := AnalyticsparticipantMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
	
	if UserId, ok := AnalyticsparticipantMap["userId"].(string); ok {
		o.UserId = &UserId
	}
	
	if Sessions, ok := AnalyticsparticipantMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	
	if Attributes, ok := AnalyticsparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
