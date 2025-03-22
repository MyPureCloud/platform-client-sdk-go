package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssession
type Analyticssession struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// AgentOwned - Flag indicating an agent-owned callback
	AgentOwned *bool `json:"agentOwned,omitempty"`

	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`

	// AssignerId - ID of the user that manually assigned a conversation
	AssignerId *string `json:"assignerId,omitempty"`

	// Authenticated - Flag that indicates that the identity of the customer has been asserted as verified by the provider.
	Authenticated *bool `json:"authenticated,omitempty"`

	// BargedParticipantId - The participantId being barged in on (if someone (e.g. an agent) is being barged in on, this would correspond to one of the other participantIds present in the conversation)
	BargedParticipantId *string `json:"bargedParticipantId,omitempty"`

	// Bcc - Blind carbon copy email address(es)
	Bcc *[]string `json:"bcc,omitempty"`

	// CallbackNumbers - Callback phone number(s)
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`

	// CallbackScheduledTime - Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`

	// CallbackUserName - The name of the user requesting a call back
	CallbackUserName *string `json:"callbackUserName,omitempty"`

	// Cc - Carbon copy email address(es)
	Cc *[]string `json:"cc,omitempty"`

	// Cleared - Flag that indicates that the conversation has been cleared by the customer
	Cleared *bool `json:"cleared,omitempty"`

	// CoachedParticipantId - The participantId being coached (if someone (e.g. an agent) is being coached, this would correspond to one of the other participantIds present in the conversation)
	CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`

	// CobrowseRole - Describes side of the cobrowse (sharer or viewer)
	CobrowseRole *string `json:"cobrowseRole,omitempty"`

	// CobrowseRoomId - A unique identifier for a Genesys Cloud cobrowse room
	CobrowseRoomId *string `json:"cobrowseRoomId,omitempty"`

	// DeliveryPushed - Flag that indicates that the push delivery mechanism was used
	DeliveryPushed *bool `json:"deliveryPushed,omitempty"`

	// DeliveryStatus - The email or SMS delivery status
	DeliveryStatus *string `json:"deliveryStatus,omitempty"`

	// DeliveryStatusChangeDate - Date and time of the most recent delivery status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DeliveryStatusChangeDate *time.Time `json:"deliveryStatusChangeDate,omitempty"`

	// DestinationAddresses - Destination address(es) of transfers or consults
	DestinationAddresses *[]string `json:"destinationAddresses,omitempty"`

	// DetectedSpeechEnd - Absolute time when the speech ended. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DetectedSpeechEnd *time.Time `json:"detectedSpeechEnd,omitempty"`

	// DetectedSpeechStart - Absolute time when the speech started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DetectedSpeechStart *time.Time `json:"detectedSpeechStart,omitempty"`

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

	// EligibleAgentCounts - Number of eligible agents for each predictive routing attempt
	EligibleAgentCounts *[]int `json:"eligibleAgentCounts,omitempty"`

	// ExtendedDeliveryStatus - Extended delivery status
	ExtendedDeliveryStatus *string `json:"extendedDeliveryStatus,omitempty"`

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

	// MonitoredParticipantId - The participantId being monitored (if someone (e.g. an agent) is being monitored, this would correspond to one of the other participantIds present in the conversation)
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

	// RoutingRule - Routing rule for preferred, conditional and predictive routing type
	RoutingRule *string `json:"routingRule,omitempty"`

	// RoutingRuleType - Routing rule type
	RoutingRuleType *string `json:"routingRuleType,omitempty"`

	// ScreenShareAddressSelf - Direct screen share address
	ScreenShareAddressSelf *string `json:"screenShareAddressSelf,omitempty"`

	// ScreenShareRoomId - A unique identifier for a Genesys Cloud screen share room
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

	// SharingScreen - Flag determining if screen share is started or not (true/false)
	SharingScreen *bool `json:"sharingScreen,omitempty"`

	// SkipEnabled - (Dialer) Whether the agent can skip the dialer contact
	SkipEnabled *bool `json:"skipEnabled,omitempty"`

	// TimeoutSeconds - The number of seconds before Genesys Cloud begins the call for a call back (0 disables automatic calling)
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`

	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`

	// VideoAddressSelf - Direct Video address
	VideoAddressSelf *string `json:"videoAddressSelf,omitempty"`

	// VideoRoomId - A unique identifier for a Genesys Cloud video room
	VideoRoomId *string `json:"videoRoomId,omitempty"`

	// WaitingInteractionCounts - Number of waiting interactions for each predictive routing attempt
	WaitingInteractionCounts *[]int `json:"waitingInteractionCounts,omitempty"`

	// AgentGroups - Conditional group routing agent groups
	AgentGroups *[]Analyticsagentgroup `json:"agentGroups,omitempty"`

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
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticssession) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Analyticssession) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CallbackScheduledTime","DeliveryStatusChangeDate","DetectedSpeechEnd","DetectedSpeechStart", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Analyticssession
	
	CallbackScheduledTime := new(string)
	if o.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(o.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	
	DeliveryStatusChangeDate := new(string)
	if o.DeliveryStatusChangeDate != nil {
		
		*DeliveryStatusChangeDate = timeutil.Strftime(o.DeliveryStatusChangeDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DeliveryStatusChangeDate = nil
	}
	
	DetectedSpeechEnd := new(string)
	if o.DetectedSpeechEnd != nil {
		
		*DetectedSpeechEnd = timeutil.Strftime(o.DetectedSpeechEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DetectedSpeechEnd = nil
	}
	
	DetectedSpeechStart := new(string)
	if o.DetectedSpeechStart != nil {
		
		*DetectedSpeechStart = timeutil.Strftime(o.DetectedSpeechStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DetectedSpeechStart = nil
	}
	
	return json.Marshal(&struct { 
		ActiveSkillIds *[]string `json:"activeSkillIds,omitempty"`
		
		AcwSkipped *bool `json:"acwSkipped,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		AddressOther *string `json:"addressOther,omitempty"`
		
		AddressSelf *string `json:"addressSelf,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
		AgentBullseyeRing *int `json:"agentBullseyeRing,omitempty"`
		
		AgentOwned *bool `json:"agentOwned,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		AssignerId *string `json:"assignerId,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		
		BargedParticipantId *string `json:"bargedParticipantId,omitempty"`
		
		Bcc *[]string `json:"bcc,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		Cc *[]string `json:"cc,omitempty"`
		
		Cleared *bool `json:"cleared,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		CobrowseRoomId *string `json:"cobrowseRoomId,omitempty"`
		
		DeliveryPushed *bool `json:"deliveryPushed,omitempty"`
		
		DeliveryStatus *string `json:"deliveryStatus,omitempty"`
		
		DeliveryStatusChangeDate *string `json:"deliveryStatusChangeDate,omitempty"`
		
		DestinationAddresses *[]string `json:"destinationAddresses,omitempty"`
		
		DetectedSpeechEnd *string `json:"detectedSpeechEnd,omitempty"`
		
		DetectedSpeechStart *string `json:"detectedSpeechStart,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DispositionAnalyzer *string `json:"dispositionAnalyzer,omitempty"`
		
		DispositionName *string `json:"dispositionName,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		EdgeId *string `json:"edgeId,omitempty"`
		
		EligibleAgentCounts *[]int `json:"eligibleAgentCounts,omitempty"`
		
		ExtendedDeliveryStatus *string `json:"extendedDeliveryStatus,omitempty"`
		
		FlowInType *string `json:"flowInType,omitempty"`
		
		FlowOutType *string `json:"flowOutType,omitempty"`
		
		JourneyActionId *string `json:"journeyActionId,omitempty"`
		
		JourneyActionMapId *string `json:"journeyActionMapId,omitempty"`
		
		JourneyActionMapVersion *int `json:"journeyActionMapVersion,omitempty"`
		
		JourneyCustomerId *string `json:"journeyCustomerId,omitempty"`
		
		JourneyCustomerIdType *string `json:"journeyCustomerIdType,omitempty"`
		
		JourneyCustomerSessionId *string `json:"journeyCustomerSessionId,omitempty"`
		
		JourneyCustomerSessionIdType *string `json:"journeyCustomerSessionIdType,omitempty"`
		
		MediaBridgeId *string `json:"mediaBridgeId,omitempty"`
		
		MediaCount *int `json:"mediaCount,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ProtocolCallId *string `json:"protocolCallId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		Remote *string `json:"remote,omitempty"`
		
		RemoteNameDisplayable *string `json:"remoteNameDisplayable,omitempty"`
		
		RemovedSkillIds *[]string `json:"removedSkillIds,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		RoutingRing *int `json:"routingRing,omitempty"`
		
		RoutingRule *string `json:"routingRule,omitempty"`
		
		RoutingRuleType *string `json:"routingRuleType,omitempty"`
		
		ScreenShareAddressSelf *string `json:"screenShareAddressSelf,omitempty"`
		
		ScreenShareRoomId *string `json:"screenShareRoomId,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		SelectedAgentId *string `json:"selectedAgentId,omitempty"`
		
		SelectedAgentRank *int `json:"selectedAgentRank,omitempty"`
		
		SessionDnis *string `json:"sessionDnis,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		SharingScreen *bool `json:"sharingScreen,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		VideoAddressSelf *string `json:"videoAddressSelf,omitempty"`
		
		VideoRoomId *string `json:"videoRoomId,omitempty"`
		
		WaitingInteractionCounts *[]int `json:"waitingInteractionCounts,omitempty"`
		
		AgentGroups *[]Analyticsagentgroup `json:"agentGroups,omitempty"`
		
		ProposedAgents *[]Analyticsproposedagent `json:"proposedAgents,omitempty"`
		
		MediaEndpointStats *[]Analyticsmediaendpointstat `json:"mediaEndpointStats,omitempty"`
		
		Flow *Analyticsflow `json:"flow,omitempty"`
		
		Metrics *[]Analyticssessionmetric `json:"metrics,omitempty"`
		
		Segments *[]Analyticsconversationsegment `json:"segments,omitempty"`
		Alias
	}{ 
		ActiveSkillIds: o.ActiveSkillIds,
		
		AcwSkipped: o.AcwSkipped,
		
		AddressFrom: o.AddressFrom,
		
		AddressOther: o.AddressOther,
		
		AddressSelf: o.AddressSelf,
		
		AddressTo: o.AddressTo,
		
		AgentAssistantId: o.AgentAssistantId,
		
		AgentBullseyeRing: o.AgentBullseyeRing,
		
		AgentOwned: o.AgentOwned,
		
		Ani: o.Ani,
		
		AssignerId: o.AssignerId,
		
		Authenticated: o.Authenticated,
		
		BargedParticipantId: o.BargedParticipantId,
		
		Bcc: o.Bcc,
		
		CallbackNumbers: o.CallbackNumbers,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		CallbackUserName: o.CallbackUserName,
		
		Cc: o.Cc,
		
		Cleared: o.Cleared,
		
		CoachedParticipantId: o.CoachedParticipantId,
		
		CobrowseRole: o.CobrowseRole,
		
		CobrowseRoomId: o.CobrowseRoomId,
		
		DeliveryPushed: o.DeliveryPushed,
		
		DeliveryStatus: o.DeliveryStatus,
		
		DeliveryStatusChangeDate: DeliveryStatusChangeDate,
		
		DestinationAddresses: o.DestinationAddresses,
		
		DetectedSpeechEnd: DetectedSpeechEnd,
		
		DetectedSpeechStart: DetectedSpeechStart,
		
		Direction: o.Direction,
		
		DispositionAnalyzer: o.DispositionAnalyzer,
		
		DispositionName: o.DispositionName,
		
		Dnis: o.Dnis,
		
		EdgeId: o.EdgeId,
		
		EligibleAgentCounts: o.EligibleAgentCounts,
		
		ExtendedDeliveryStatus: o.ExtendedDeliveryStatus,
		
		FlowInType: o.FlowInType,
		
		FlowOutType: o.FlowOutType,
		
		JourneyActionId: o.JourneyActionId,
		
		JourneyActionMapId: o.JourneyActionMapId,
		
		JourneyActionMapVersion: o.JourneyActionMapVersion,
		
		JourneyCustomerId: o.JourneyCustomerId,
		
		JourneyCustomerIdType: o.JourneyCustomerIdType,
		
		JourneyCustomerSessionId: o.JourneyCustomerSessionId,
		
		JourneyCustomerSessionIdType: o.JourneyCustomerSessionIdType,
		
		MediaBridgeId: o.MediaBridgeId,
		
		MediaCount: o.MediaCount,
		
		MediaType: o.MediaType,
		
		MessageType: o.MessageType,
		
		MonitoredParticipantId: o.MonitoredParticipantId,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		OutboundContactId: o.OutboundContactId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		PeerId: o.PeerId,
		
		ProtocolCallId: o.ProtocolCallId,
		
		Provider: o.Provider,
		
		Recording: o.Recording,
		
		Remote: o.Remote,
		
		RemoteNameDisplayable: o.RemoteNameDisplayable,
		
		RemovedSkillIds: o.RemovedSkillIds,
		
		RequestedRoutings: o.RequestedRoutings,
		
		RoomId: o.RoomId,
		
		RoutingRing: o.RoutingRing,
		
		RoutingRule: o.RoutingRule,
		
		RoutingRuleType: o.RoutingRuleType,
		
		ScreenShareAddressSelf: o.ScreenShareAddressSelf,
		
		ScreenShareRoomId: o.ScreenShareRoomId,
		
		ScriptId: o.ScriptId,
		
		SelectedAgentId: o.SelectedAgentId,
		
		SelectedAgentRank: o.SelectedAgentRank,
		
		SessionDnis: o.SessionDnis,
		
		SessionId: o.SessionId,
		
		SharingScreen: o.SharingScreen,
		
		SkipEnabled: o.SkipEnabled,
		
		TimeoutSeconds: o.TimeoutSeconds,
		
		UsedRouting: o.UsedRouting,
		
		VideoAddressSelf: o.VideoAddressSelf,
		
		VideoRoomId: o.VideoRoomId,
		
		WaitingInteractionCounts: o.WaitingInteractionCounts,
		
		AgentGroups: o.AgentGroups,
		
		ProposedAgents: o.ProposedAgents,
		
		MediaEndpointStats: o.MediaEndpointStats,
		
		Flow: o.Flow,
		
		Metrics: o.Metrics,
		
		Segments: o.Segments,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticssession) UnmarshalJSON(b []byte) error {
	var AnalyticssessionMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticssessionMap)
	if err != nil {
		return err
	}
	
	if ActiveSkillIds, ok := AnalyticssessionMap["activeSkillIds"].([]interface{}); ok {
		ActiveSkillIdsString, _ := json.Marshal(ActiveSkillIds)
		json.Unmarshal(ActiveSkillIdsString, &o.ActiveSkillIds)
	}
	
	if AcwSkipped, ok := AnalyticssessionMap["acwSkipped"].(bool); ok {
		o.AcwSkipped = &AcwSkipped
	}
    
	if AddressFrom, ok := AnalyticssessionMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if AddressOther, ok := AnalyticssessionMap["addressOther"].(string); ok {
		o.AddressOther = &AddressOther
	}
    
	if AddressSelf, ok := AnalyticssessionMap["addressSelf"].(string); ok {
		o.AddressSelf = &AddressSelf
	}
    
	if AddressTo, ok := AnalyticssessionMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if AgentAssistantId, ok := AnalyticssessionMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
    
	if AgentBullseyeRing, ok := AnalyticssessionMap["agentBullseyeRing"].(float64); ok {
		AgentBullseyeRingInt := int(AgentBullseyeRing)
		o.AgentBullseyeRing = &AgentBullseyeRingInt
	}
	
	if AgentOwned, ok := AnalyticssessionMap["agentOwned"].(bool); ok {
		o.AgentOwned = &AgentOwned
	}
    
	if Ani, ok := AnalyticssessionMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if AssignerId, ok := AnalyticssessionMap["assignerId"].(string); ok {
		o.AssignerId = &AssignerId
	}
    
	if Authenticated, ok := AnalyticssessionMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    
	if BargedParticipantId, ok := AnalyticssessionMap["bargedParticipantId"].(string); ok {
		o.BargedParticipantId = &BargedParticipantId
	}
    
	if Bcc, ok := AnalyticssessionMap["bcc"].([]interface{}); ok {
		BccString, _ := json.Marshal(Bcc)
		json.Unmarshal(BccString, &o.Bcc)
	}
	
	if CallbackNumbers, ok := AnalyticssessionMap["callbackNumbers"].([]interface{}); ok {
		CallbackNumbersString, _ := json.Marshal(CallbackNumbers)
		json.Unmarshal(CallbackNumbersString, &o.CallbackNumbers)
	}
	
	if callbackScheduledTimeString, ok := AnalyticssessionMap["callbackScheduledTime"].(string); ok {
		CallbackScheduledTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", callbackScheduledTimeString)
		o.CallbackScheduledTime = &CallbackScheduledTime
	}
	
	if CallbackUserName, ok := AnalyticssessionMap["callbackUserName"].(string); ok {
		o.CallbackUserName = &CallbackUserName
	}
    
	if Cc, ok := AnalyticssessionMap["cc"].([]interface{}); ok {
		CcString, _ := json.Marshal(Cc)
		json.Unmarshal(CcString, &o.Cc)
	}
	
	if Cleared, ok := AnalyticssessionMap["cleared"].(bool); ok {
		o.Cleared = &Cleared
	}
    
	if CoachedParticipantId, ok := AnalyticssessionMap["coachedParticipantId"].(string); ok {
		o.CoachedParticipantId = &CoachedParticipantId
	}
    
	if CobrowseRole, ok := AnalyticssessionMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}
    
	if CobrowseRoomId, ok := AnalyticssessionMap["cobrowseRoomId"].(string); ok {
		o.CobrowseRoomId = &CobrowseRoomId
	}
    
	if DeliveryPushed, ok := AnalyticssessionMap["deliveryPushed"].(bool); ok {
		o.DeliveryPushed = &DeliveryPushed
	}
    
	if DeliveryStatus, ok := AnalyticssessionMap["deliveryStatus"].(string); ok {
		o.DeliveryStatus = &DeliveryStatus
	}
    
	if deliveryStatusChangeDateString, ok := AnalyticssessionMap["deliveryStatusChangeDate"].(string); ok {
		DeliveryStatusChangeDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", deliveryStatusChangeDateString)
		o.DeliveryStatusChangeDate = &DeliveryStatusChangeDate
	}
	
	if DestinationAddresses, ok := AnalyticssessionMap["destinationAddresses"].([]interface{}); ok {
		DestinationAddressesString, _ := json.Marshal(DestinationAddresses)
		json.Unmarshal(DestinationAddressesString, &o.DestinationAddresses)
	}
	
	if detectedSpeechEndString, ok := AnalyticssessionMap["detectedSpeechEnd"].(string); ok {
		DetectedSpeechEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", detectedSpeechEndString)
		o.DetectedSpeechEnd = &DetectedSpeechEnd
	}
	
	if detectedSpeechStartString, ok := AnalyticssessionMap["detectedSpeechStart"].(string); ok {
		DetectedSpeechStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", detectedSpeechStartString)
		o.DetectedSpeechStart = &DetectedSpeechStart
	}
	
	if Direction, ok := AnalyticssessionMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DispositionAnalyzer, ok := AnalyticssessionMap["dispositionAnalyzer"].(string); ok {
		o.DispositionAnalyzer = &DispositionAnalyzer
	}
    
	if DispositionName, ok := AnalyticssessionMap["dispositionName"].(string); ok {
		o.DispositionName = &DispositionName
	}
    
	if Dnis, ok := AnalyticssessionMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if EdgeId, ok := AnalyticssessionMap["edgeId"].(string); ok {
		o.EdgeId = &EdgeId
	}
    
	if EligibleAgentCounts, ok := AnalyticssessionMap["eligibleAgentCounts"].([]interface{}); ok {
		EligibleAgentCountsString, _ := json.Marshal(EligibleAgentCounts)
		json.Unmarshal(EligibleAgentCountsString, &o.EligibleAgentCounts)
	}
	
	if ExtendedDeliveryStatus, ok := AnalyticssessionMap["extendedDeliveryStatus"].(string); ok {
		o.ExtendedDeliveryStatus = &ExtendedDeliveryStatus
	}
    
	if FlowInType, ok := AnalyticssessionMap["flowInType"].(string); ok {
		o.FlowInType = &FlowInType
	}
    
	if FlowOutType, ok := AnalyticssessionMap["flowOutType"].(string); ok {
		o.FlowOutType = &FlowOutType
	}
    
	if JourneyActionId, ok := AnalyticssessionMap["journeyActionId"].(string); ok {
		o.JourneyActionId = &JourneyActionId
	}
    
	if JourneyActionMapId, ok := AnalyticssessionMap["journeyActionMapId"].(string); ok {
		o.JourneyActionMapId = &JourneyActionMapId
	}
    
	if JourneyActionMapVersion, ok := AnalyticssessionMap["journeyActionMapVersion"].(float64); ok {
		JourneyActionMapVersionInt := int(JourneyActionMapVersion)
		o.JourneyActionMapVersion = &JourneyActionMapVersionInt
	}
	
	if JourneyCustomerId, ok := AnalyticssessionMap["journeyCustomerId"].(string); ok {
		o.JourneyCustomerId = &JourneyCustomerId
	}
    
	if JourneyCustomerIdType, ok := AnalyticssessionMap["journeyCustomerIdType"].(string); ok {
		o.JourneyCustomerIdType = &JourneyCustomerIdType
	}
    
	if JourneyCustomerSessionId, ok := AnalyticssessionMap["journeyCustomerSessionId"].(string); ok {
		o.JourneyCustomerSessionId = &JourneyCustomerSessionId
	}
    
	if JourneyCustomerSessionIdType, ok := AnalyticssessionMap["journeyCustomerSessionIdType"].(string); ok {
		o.JourneyCustomerSessionIdType = &JourneyCustomerSessionIdType
	}
    
	if MediaBridgeId, ok := AnalyticssessionMap["mediaBridgeId"].(string); ok {
		o.MediaBridgeId = &MediaBridgeId
	}
    
	if MediaCount, ok := AnalyticssessionMap["mediaCount"].(float64); ok {
		MediaCountInt := int(MediaCount)
		o.MediaCount = &MediaCountInt
	}
	
	if MediaType, ok := AnalyticssessionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if MessageType, ok := AnalyticssessionMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if MonitoredParticipantId, ok := AnalyticssessionMap["monitoredParticipantId"].(string); ok {
		o.MonitoredParticipantId = &MonitoredParticipantId
	}
    
	if OutboundCampaignId, ok := AnalyticssessionMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
    
	if OutboundContactId, ok := AnalyticssessionMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
    
	if OutboundContactListId, ok := AnalyticssessionMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
    
	if PeerId, ok := AnalyticssessionMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if ProtocolCallId, ok := AnalyticssessionMap["protocolCallId"].(string); ok {
		o.ProtocolCallId = &ProtocolCallId
	}
    
	if Provider, ok := AnalyticssessionMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Recording, ok := AnalyticssessionMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if Remote, ok := AnalyticssessionMap["remote"].(string); ok {
		o.Remote = &Remote
	}
    
	if RemoteNameDisplayable, ok := AnalyticssessionMap["remoteNameDisplayable"].(string); ok {
		o.RemoteNameDisplayable = &RemoteNameDisplayable
	}
    
	if RemovedSkillIds, ok := AnalyticssessionMap["removedSkillIds"].([]interface{}); ok {
		RemovedSkillIdsString, _ := json.Marshal(RemovedSkillIds)
		json.Unmarshal(RemovedSkillIdsString, &o.RemovedSkillIds)
	}
	
	if RequestedRoutings, ok := AnalyticssessionMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if RoomId, ok := AnalyticssessionMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
    
	if RoutingRing, ok := AnalyticssessionMap["routingRing"].(float64); ok {
		RoutingRingInt := int(RoutingRing)
		o.RoutingRing = &RoutingRingInt
	}
	
	if RoutingRule, ok := AnalyticssessionMap["routingRule"].(string); ok {
		o.RoutingRule = &RoutingRule
	}
    
	if RoutingRuleType, ok := AnalyticssessionMap["routingRuleType"].(string); ok {
		o.RoutingRuleType = &RoutingRuleType
	}
    
	if ScreenShareAddressSelf, ok := AnalyticssessionMap["screenShareAddressSelf"].(string); ok {
		o.ScreenShareAddressSelf = &ScreenShareAddressSelf
	}
    
	if ScreenShareRoomId, ok := AnalyticssessionMap["screenShareRoomId"].(string); ok {
		o.ScreenShareRoomId = &ScreenShareRoomId
	}
    
	if ScriptId, ok := AnalyticssessionMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
    
	if SelectedAgentId, ok := AnalyticssessionMap["selectedAgentId"].(string); ok {
		o.SelectedAgentId = &SelectedAgentId
	}
    
	if SelectedAgentRank, ok := AnalyticssessionMap["selectedAgentRank"].(float64); ok {
		SelectedAgentRankInt := int(SelectedAgentRank)
		o.SelectedAgentRank = &SelectedAgentRankInt
	}
	
	if SessionDnis, ok := AnalyticssessionMap["sessionDnis"].(string); ok {
		o.SessionDnis = &SessionDnis
	}
    
	if SessionId, ok := AnalyticssessionMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if SharingScreen, ok := AnalyticssessionMap["sharingScreen"].(bool); ok {
		o.SharingScreen = &SharingScreen
	}
    
	if SkipEnabled, ok := AnalyticssessionMap["skipEnabled"].(bool); ok {
		o.SkipEnabled = &SkipEnabled
	}
    
	if TimeoutSeconds, ok := AnalyticssessionMap["timeoutSeconds"].(float64); ok {
		TimeoutSecondsInt := int(TimeoutSeconds)
		o.TimeoutSeconds = &TimeoutSecondsInt
	}
	
	if UsedRouting, ok := AnalyticssessionMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
    
	if VideoAddressSelf, ok := AnalyticssessionMap["videoAddressSelf"].(string); ok {
		o.VideoAddressSelf = &VideoAddressSelf
	}
    
	if VideoRoomId, ok := AnalyticssessionMap["videoRoomId"].(string); ok {
		o.VideoRoomId = &VideoRoomId
	}
    
	if WaitingInteractionCounts, ok := AnalyticssessionMap["waitingInteractionCounts"].([]interface{}); ok {
		WaitingInteractionCountsString, _ := json.Marshal(WaitingInteractionCounts)
		json.Unmarshal(WaitingInteractionCountsString, &o.WaitingInteractionCounts)
	}
	
	if AgentGroups, ok := AnalyticssessionMap["agentGroups"].([]interface{}); ok {
		AgentGroupsString, _ := json.Marshal(AgentGroups)
		json.Unmarshal(AgentGroupsString, &o.AgentGroups)
	}
	
	if ProposedAgents, ok := AnalyticssessionMap["proposedAgents"].([]interface{}); ok {
		ProposedAgentsString, _ := json.Marshal(ProposedAgents)
		json.Unmarshal(ProposedAgentsString, &o.ProposedAgents)
	}
	
	if MediaEndpointStats, ok := AnalyticssessionMap["mediaEndpointStats"].([]interface{}); ok {
		MediaEndpointStatsString, _ := json.Marshal(MediaEndpointStats)
		json.Unmarshal(MediaEndpointStatsString, &o.MediaEndpointStats)
	}
	
	if Flow, ok := AnalyticssessionMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if Metrics, ok := AnalyticssessionMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Segments, ok := AnalyticssessionMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticssession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
