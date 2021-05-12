package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssession
type Analyticssession struct { 
	// ActiveSkillIds - ID(s) of Skill(s) that are active on the conversation
	ActiveSkillIds *[]string `json:"activeSkillIds,omitempty"`


	// AcwSkipped - Marker for an agent that skipped after call work
	AcwSkipped *bool `json:"acwSkipped,omitempty"`


	// AddressFrom - The address that initiated an action
	AddressFrom *string `json:"addressFrom,omitempty"`


	// AddressOther - The email address for the participant on the other side of the email conversation
	AddressOther *string `json:"addressOther,omitempty"`


	// AddressSelf - The email address for the participant on this side of the email conversation
	AddressSelf *string `json:"addressSelf,omitempty"`


	// AddressTo - The address receiving an action
	AddressTo *string `json:"addressTo,omitempty"`


	// AgentAssistantId - Unique identifier of the active virtual agent assistant
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// AgentBullseyeRing - Bullseye ring of the targeted agent
	AgentBullseyeRing *int `json:"agentBullseyeRing,omitempty"`


	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`


	// AssignerId - ID of the user that manually assigned a conversation
	AssignerId *string `json:"assignerId,omitempty"`


	// Authenticated - Flag that indicates that the identity of the customer has been asserted as verified by the provider.
	Authenticated *bool `json:"authenticated,omitempty"`


	// CallbackNumbers - Callback phone number(s)
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackScheduledTime - Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// CallbackUserName - The name of the user requesting a call back
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// CobrowseRole - Describes side of the cobrowse (sharer or viewer)
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// CobrowseRoomId - A unique identifier for a PureCloud cobrowse room
	CobrowseRoomId *string `json:"cobrowseRoomId,omitempty"`


	// Direction - The direction of the communication
	Direction *string `json:"direction,omitempty"`


	// DispositionAnalyzer - (Dialer) Analyzer (for example speech.person)
	DispositionAnalyzer *string `json:"dispositionAnalyzer,omitempty"`


	// DispositionName - (Dialer) Result of the analysis (for example disposition.classification.callable.machine)
	DispositionName *string `json:"dispositionName,omitempty"`


	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`


	// EdgeId - Unique identifier of the edge device
	EdgeId *string `json:"edgeId,omitempty"`


	// FlowInType - Type of flow in that occurred when entering ACD.
	FlowInType *string `json:"flowInType,omitempty"`


	// FlowOutType - Type of flow out that occurred when emitting tFlowOut.
	FlowOutType *string `json:"flowOutType,omitempty"`


	// JourneyActionId - Identifier of the journey action.
	JourneyActionId *string `json:"journeyActionId,omitempty"`


	// JourneyActionMapId - Identifier of the journey action map that triggered the action.
	JourneyActionMapId *string `json:"journeyActionMapId,omitempty"`


	// JourneyActionMapVersion - Version of the journey action map that triggered the action.
	JourneyActionMapVersion *int `json:"journeyActionMapVersion,omitempty"`


	// JourneyCustomerId - Primary identifier of the journey customer in the source where the activities originate from.
	JourneyCustomerId *string `json:"journeyCustomerId,omitempty"`


	// JourneyCustomerIdType - Type of primary identifier of the journey customer (e.g. cookie).
	JourneyCustomerIdType *string `json:"journeyCustomerIdType,omitempty"`


	// JourneyCustomerSessionId - Unique identifier of the journey session.
	JourneyCustomerSessionId *string `json:"journeyCustomerSessionId,omitempty"`


	// JourneyCustomerSessionIdType - Type or category of journey sessions (e.g. web, ticket, delivery, atm).
	JourneyCustomerSessionIdType *string `json:"journeyCustomerSessionIdType,omitempty"`


	// MediaBridgeId - Media bridge ID for the conference session consistent across all participants
	MediaBridgeId *string `json:"mediaBridgeId,omitempty"`


	// MediaCount - Count of any media (images, files, etc) included in this session
	MediaCount *int `json:"mediaCount,omitempty"`


	// MediaType - The session media type
	MediaType *string `json:"mediaType,omitempty"`


	// MessageType - Message type for messaging services. E.g.: sms, facebook, twitter, line
	MessageType *string `json:"messageType,omitempty"`


	// MonitoredParticipantId - The participantId being monitored (if someone (e.g. an agent) is being monitored, this would be the ID of the participant that was monitored that would correspond to other participantIds present in the conversation)
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// OutboundCampaignId - (Dialer) Unique identifier of the outbound campaign
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`


	// OutboundContactId - (Dialer) Unique identifier of the contact
	OutboundContactId *string `json:"outboundContactId,omitempty"`


	// OutboundContactListId - (Dialer) Unique identifier of the contact list that this contact belongs to
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`


	// PeerId - This identifies pairs of related sessions on a conversation. E.g. an external session’s peerId will be the session that the call originally connected to, e.g. if an IVR was dialed, the IVR session, which will also have the external session’s ID as its peer. After that point, any transfers of that session to other internal components (acd, agent, etc.) will all spawn new sessions whose peerIds point back to that original external session.
	PeerId *string `json:"peerId,omitempty"`


	// ProtocolCallId - The original voice protocol call ID, e.g. a SIP call ID
	ProtocolCallId *string `json:"protocolCallId,omitempty"`


	// Provider - The source provider for the communication.
	Provider *string `json:"provider,omitempty"`


	// Recording - Flag determining if an audio recording was started or not
	Recording *bool `json:"recording,omitempty"`


	// Remote - Name, phone number, or email address of the remote party.
	Remote *string `json:"remote,omitempty"`


	// RemoteNameDisplayable - Unique identifier for the remote party
	RemoteNameDisplayable *string `json:"remoteNameDisplayable,omitempty"`


	// RemovedSkillIds - ID(s) of Skill(s) that have been removed by bullseye routing
	RemovedSkillIds *[]string `json:"removedSkillIds,omitempty"`


	// RequestedRoutings - Routing type(s) for requested/attempted routing methods.
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`


	// RoomId - Unique identifier for the room
	RoomId *string `json:"roomId,omitempty"`


	// RoutingRing - Routing ring for bullseye or preferred agent routing
	RoutingRing *int `json:"routingRing,omitempty"`


	// ScreenShareAddressSelf - Direct ScreenShare address
	ScreenShareAddressSelf *string `json:"screenShareAddressSelf,omitempty"`


	// ScreenShareRoomId - A unique identifier for a PureCloud ScreenShare room
	ScreenShareRoomId *string `json:"screenShareRoomId,omitempty"`


	// ScriptId - A unique identifier for a script
	ScriptId *string `json:"scriptId,omitempty"`


	// SelectedAgentId - Selected agent ID
	SelectedAgentId *string `json:"selectedAgentId,omitempty"`


	// SelectedAgentRank - Selected agent GPR rank
	SelectedAgentRank *int `json:"selectedAgentRank,omitempty"`


	// SessionDnis - Dialed number for the current session; this can be different from dnis, e.g. if the call was transferred
	SessionDnis *string `json:"sessionDnis,omitempty"`


	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`


	// SharingScreen - Flag determining if screenShare is started or not (true/false)
	SharingScreen *bool `json:"sharingScreen,omitempty"`


	// SkipEnabled - (Dialer) Whether the agent can skip the dialer contact
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`


	// VideoAddressSelf - Direct Video address
	VideoAddressSelf *string `json:"videoAddressSelf,omitempty"`


	// VideoRoomId - A unique identifier for a PureCloud video room
	VideoRoomId *string `json:"videoRoomId,omitempty"`


	// ProposedAgents - Proposed agents
	ProposedAgents *[]Analyticsproposedagent `json:"proposedAgents,omitempty"`


	// MediaEndpointStats - MediaEndpointStats associated with this session
	MediaEndpointStats *[]Analyticsmediaendpointstat `json:"mediaEndpointStats,omitempty"`


	// Flow - IVR flow execution associated with this session
	Flow *Analyticsflow `json:"flow,omitempty"`


	// Metrics - List of metrics for this session
	Metrics *[]Analyticssessionmetric `json:"metrics,omitempty"`


	// Segments - List of segments for this session
	Segments *[]Analyticsconversationsegment `json:"segments,omitempty"`


	// TimeoutSeconds - The number of seconds before PureCloud begins the call for a call back (0 disables automatic calling)
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticssession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
