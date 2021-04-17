package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsconversationsegment
type Analyticsconversationsegment struct { 
	// AudioMuted - Flag indicating if audio is muted or not (true/false)
	AudioMuted *bool `json:"audioMuted,omitempty"`


	// Conference - Indicates whether the segment was a conference
	Conference *bool `json:"conference,omitempty"`


	// DestinationConversationId - The unique identifier of a new conversation when a conversation is ended for a conference
	DestinationConversationId *string `json:"destinationConversationId,omitempty"`


	// DestinationSessionId - The unique identifier of a new session when a session is ended for a conference
	DestinationSessionId *string `json:"destinationSessionId,omitempty"`


	// DisconnectType - The session disconnect type
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ErrorCode - A code corresponding to the error that occurred
	ErrorCode *string `json:"errorCode,omitempty"`


	// GroupId - Unique identifier for a PureCloud group
	GroupId *string `json:"groupId,omitempty"`


	// Q850ResponseCodes - Q.850 response code(s)
	Q850ResponseCodes *[]int `json:"q850ResponseCodes,omitempty"`


	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`


	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`


	// RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`


	// RequestedRoutingUserIds - Unique identifier(s) for agent(s) requested for an interaction
	RequestedRoutingUserIds *[]string `json:"requestedRoutingUserIds,omitempty"`


	// SegmentEnd - The end time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentEnd *time.Time `json:"segmentEnd,omitempty"`


	// SegmentStart - The start time of a segment. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	SegmentStart *time.Time `json:"segmentStart,omitempty"`


	// SegmentType - The activity that takes place in the segment, such as hold or interact
	SegmentType *string `json:"segmentType,omitempty"`


	// SipResponseCodes - SIP response code(s)
	SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`


	// SourceConversationId - The unique identifier of the previous conversation when a new conversation is created for a conference
	SourceConversationId *string `json:"sourceConversationId,omitempty"`


	// SourceSessionId - The unique identifier of the previous session when a new session is created for a conference
	SourceSessionId *string `json:"sourceSessionId,omitempty"`


	// Subject - The subject for the initial email that started this conversation
	Subject *string `json:"subject,omitempty"`


	// VideoMuted - Flag indicating if video is muted/paused or not (true/false)
	VideoMuted *bool `json:"videoMuted,omitempty"`


	// WrapUpCode - Wrap up code
	WrapUpCode *string `json:"wrapUpCode,omitempty"`


	// WrapUpNote - Note entered by an agent during after-call work
	WrapUpNote *string `json:"wrapUpNote,omitempty"`


	// WrapUpTags - Tag(s) assigned during after-call work
	WrapUpTags *[]string `json:"wrapUpTags,omitempty"`


	// ScoredAgents - Scored agents
	ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`


	// Properties - Additional segment properties
	Properties *[]Analyticsproperty `json:"properties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
