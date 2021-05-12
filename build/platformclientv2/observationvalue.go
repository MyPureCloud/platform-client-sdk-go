package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Observationvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
