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

func (o *Analyticsparticipantwithoutattributes) MarshalJSON() ([]byte, error) {
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
		ExternalContactId: o.ExternalContactId,
		
		ExternalOrganizationId: o.ExternalOrganizationId,
		
		FlaggedReason: o.FlaggedReason,
		
		ParticipantId: o.ParticipantId,
		
		ParticipantName: o.ParticipantName,
		
		Purpose: o.Purpose,
		
		TeamId: o.TeamId,
		
		UserId: o.UserId,
		
		Sessions: o.Sessions,
		Alias:    (*Alias)(o),
	})
}

func (o *Analyticsparticipantwithoutattributes) UnmarshalJSON(b []byte) error {
	var AnalyticsparticipantwithoutattributesMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsparticipantwithoutattributesMap)
	if err != nil {
		return err
	}
	
	if ExternalContactId, ok := AnalyticsparticipantwithoutattributesMap["externalContactId"].(string); ok {
		o.ExternalContactId = &ExternalContactId
	}
    
	if ExternalOrganizationId, ok := AnalyticsparticipantwithoutattributesMap["externalOrganizationId"].(string); ok {
		o.ExternalOrganizationId = &ExternalOrganizationId
	}
    
	if FlaggedReason, ok := AnalyticsparticipantwithoutattributesMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if ParticipantId, ok := AnalyticsparticipantwithoutattributesMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if ParticipantName, ok := AnalyticsparticipantwithoutattributesMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if Purpose, ok := AnalyticsparticipantwithoutattributesMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if TeamId, ok := AnalyticsparticipantwithoutattributesMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if UserId, ok := AnalyticsparticipantwithoutattributesMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Sessions, ok := AnalyticsparticipantwithoutattributesMap["sessions"].([]interface{}); ok {
		SessionsString, _ := json.Marshal(Sessions)
		json.Unmarshal(SessionsString, &o.Sessions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsparticipantwithoutattributes) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
