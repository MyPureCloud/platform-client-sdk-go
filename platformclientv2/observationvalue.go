package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Observationvalue
type Observationvalue struct { 
	// ObservationDate - The time at which the observation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ObservationDate *time.Time `json:"observationDate,omitempty"`


	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`


	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`


	// RequestedRoutingSkillIds - Unique identifier for a skill requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`


	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`


	// RoutingPriority - Routing priority for the current interaction
	RoutingPriority *int `json:"routingPriority,omitempty"`


	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`


	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`


	// Direction - The direction of the communication
	Direction *string `json:"direction,omitempty"`


	// ConvertedFrom - Session media type that was converted from in case of a media type conversion
	ConvertedFrom *string `json:"convertedFrom,omitempty"`


	// ConvertedTo - Session media type that was converted to in case of a media type conversion
	ConvertedTo *string `json:"convertedTo,omitempty"`


	// AddressFrom - The address that initiated an action
	AddressFrom *string `json:"addressFrom,omitempty"`


	// AddressTo - The address receiving an action
	AddressTo *string `json:"addressTo,omitempty"`


	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`


	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`


	// TeamId - The team id the user is a member of
	TeamId *string `json:"teamId,omitempty"`


	// RequestedRoutings - All routing types for requested/attempted routing methods
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`


	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`


	// ScoredAgents
	ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`

}

func (u *Observationvalue) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Observationvalue

	
	ObservationDate := new(string)
	if u.ObservationDate != nil {
		
		*ObservationDate = timeutil.Strftime(u.ObservationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ObservationDate = nil
	}
	

	return json.Marshal(&struct { 
		ObservationDate *string `json:"observationDate,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RoutingPriority *int `json:"routingPriority,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		ConvertedFrom *string `json:"convertedFrom,omitempty"`
		
		ConvertedTo *string `json:"convertedTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`
		*Alias
	}{ 
		ObservationDate: ObservationDate,
		
		ConversationId: u.ConversationId,
		
		SessionId: u.SessionId,
		
		RequestedRoutingSkillIds: u.RequestedRoutingSkillIds,
		
		RequestedLanguageId: u.RequestedLanguageId,
		
		RoutingPriority: u.RoutingPriority,
		
		ParticipantName: u.ParticipantName,
		
		UserId: u.UserId,
		
		Direction: u.Direction,
		
		ConvertedFrom: u.ConvertedFrom,
		
		ConvertedTo: u.ConvertedTo,
		
		AddressFrom: u.AddressFrom,
		
		AddressTo: u.AddressTo,
		
		Ani: u.Ani,
		
		Dnis: u.Dnis,
		
		TeamId: u.TeamId,
		
		RequestedRoutings: u.RequestedRoutings,
		
		UsedRouting: u.UsedRouting,
		
		ScoredAgents: u.ScoredAgents,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Observationvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
