package platformclientv2
import (
	"time"
	"encoding/json"
)

// Analyticssession
type Analyticssession struct { 
	// MediaType - The session media type
	MediaType *string `json:"mediaType,omitempty"`


	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`


	// AddressOther
	AddressOther *string `json:"addressOther,omitempty"`


	// AddressSelf
	AddressSelf *string `json:"addressSelf,omitempty"`


	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`


	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`


	// MessageType - Message type for messaging services such as sms
	MessageType *string `json:"messageType,omitempty"`


	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`


	// Direction - Direction
	Direction *string `json:"direction,omitempty"`


	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`


	// SessionDnis - Dialed number for the current session; this can be different from dnis, e.g. if the call was transferred
	SessionDnis *string `json:"sessionDnis,omitempty"`


	// OutboundCampaignId - (Dialer) Unique identifier of the outbound campaign
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`


	// OutboundContactId - (Dialer) Unique identifier of the contact
	OutboundContactId *string `json:"outboundContactId,omitempty"`


	// OutboundContactListId - (Dialer) Unique identifier of the contact list that this contact belongs to
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`


	// DispositionAnalyzer - (Dialer) Unique identifier of the contact list that this contact belongs to
	DispositionAnalyzer *string `json:"dispositionAnalyzer,omitempty"`


	// DispositionName - (Dialer) Result of the analysis
	DispositionName *string `json:"dispositionName,omitempty"`


	// EdgeId - Unique identifier of the edge device
	EdgeId *string `json:"edgeId,omitempty"`


	// RemoteNameDisplayable
	RemoteNameDisplayable *string `json:"remoteNameDisplayable,omitempty"`


	// RoomId - Unique identifier for the room
	RoomId *string `json:"roomId,omitempty"`


	// MonitoredSessionId - The sessionID being monitored
	MonitoredSessionId *string `json:"monitoredSessionId,omitempty"`


	// MonitoredParticipantId
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// CallbackUserName - The name of the user requesting a call back
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// CallbackNumbers - List of numbers to callback
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackScheduledTime - Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// ScriptId - A unique identifier for a script
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId - A unique identifier for a peer
	PeerId *string `json:"peerId,omitempty"`


	// SkipEnabled - (Dialer) Whether the agent can skip the dialer contact
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// TimeoutSeconds - The number of seconds before PureCloud begins the call for a call back. 0 disables automatic calling
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`


	// CobrowseRole - Describe side of the cobrowse (sharer or viewer)
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// CobrowseRoomId - A unique identifier for a PureCloud Cobrowse room.
	CobrowseRoomId *string `json:"cobrowseRoomId,omitempty"`


	// MediaBridgeId
	MediaBridgeId *string `json:"mediaBridgeId,omitempty"`


	// ScreenShareAddressSelf - Direct ScreenShare address
	ScreenShareAddressSelf *string `json:"screenShareAddressSelf,omitempty"`


	// SharingScreen - Flag determining if screenShare is started or not (true/false)
	SharingScreen *bool `json:"sharingScreen,omitempty"`


	// ScreenShareRoomId - A unique identifier for a PureCloud ScreenShare room.
	ScreenShareRoomId *string `json:"screenShareRoomId,omitempty"`


	// VideoRoomId - A unique identifier for a PureCloud video room.
	VideoRoomId *string `json:"videoRoomId,omitempty"`


	// VideoAddressSelf - Direct Video address
	VideoAddressSelf *string `json:"videoAddressSelf,omitempty"`


	// Segments - List of segments for this session
	Segments *[]Analyticsconversationsegment `json:"segments,omitempty"`


	// Metrics - List of metrics for this session
	Metrics *[]Analyticssessionmetric `json:"metrics,omitempty"`


	// Flow - IVR flow execution associated with this session
	Flow *Analyticsflow `json:"flow,omitempty"`


	// MediaEndpointStats - Media endpoint stats associated with this session
	MediaEndpointStats *[]Analyticsmediaendpointstat `json:"mediaEndpointStats,omitempty"`


	// Recording - Flag determining if an audio recording was started or not
	Recording *bool `json:"recording,omitempty"`


	// JourneyCustomerId - ID of the journey customer
	JourneyCustomerId *string `json:"journeyCustomerId,omitempty"`


	// JourneyCustomerIdType - Type of the journey customer ID
	JourneyCustomerIdType *string `json:"journeyCustomerIdType,omitempty"`


	// JourneyCustomerSessionId - ID of the journey customer session
	JourneyCustomerSessionId *string `json:"journeyCustomerSessionId,omitempty"`


	// JourneyCustomerSessionIdType - Type of the journey customer session ID
	JourneyCustomerSessionIdType *string `json:"journeyCustomerSessionIdType,omitempty"`


	// JourneyActionId - Journey action ID
	JourneyActionId *string `json:"journeyActionId,omitempty"`


	// JourneyActionMapId - Journey action map ID
	JourneyActionMapId *string `json:"journeyActionMapId,omitempty"`


	// JourneyActionMapVersion - Journey action map version
	JourneyActionMapVersion *string `json:"journeyActionMapVersion,omitempty"`


	// ProtocolCallId - The original voice protocol call ID, e.g. a SIP call ID
	ProtocolCallId *string `json:"protocolCallId,omitempty"`


	// Provider - The source provider for the communication
	Provider *string `json:"provider,omitempty"`


	// Remote - Name, phone number, or email address of the remote party.
	Remote *string `json:"remote,omitempty"`


	// MediaCount - Count of any media (images, files, etc) included in this session
	MediaCount *int32 `json:"mediaCount,omitempty"`


	// FlowOutType - Type of flow out that occurred, e.g. voicemail, callback, or acd
	FlowOutType *string `json:"flowOutType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Analyticssession) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
