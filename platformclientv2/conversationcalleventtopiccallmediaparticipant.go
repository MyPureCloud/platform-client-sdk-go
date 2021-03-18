package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopiccallmediaparticipant
type Conversationcalleventtopiccallmediaparticipant struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose
	Purpose *string `json:"purpose,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User
	User *Conversationcalleventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Conversationcalleventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Conversationcalleventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationcalleventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Conversationcalleventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Conversationcalleventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Conversationcalleventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Conversationcalleventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Conversationcalleventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Conversationcalleventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// Muted
	Muted *bool `json:"muted,omitempty"`


	// Confined
	Confined *bool `json:"confined,omitempty"`


	// Recording
	Recording *bool `json:"recording,omitempty"`


	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`


	// Group
	Group *Conversationcalleventtopicurireference `json:"group,omitempty"`


	// Ani
	Ani *string `json:"ani,omitempty"`


	// Dnis
	Dnis *string `json:"dnis,omitempty"`


	// DocumentId
	DocumentId *string `json:"documentId,omitempty"`


	// MonitoredParticipantId
	MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`


	// ConsultParticipantId
	ConsultParticipantId *string `json:"consultParticipantId,omitempty"`


	// FaxStatus
	FaxStatus *Conversationcalleventtopicfaxstatus `json:"faxStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopiccallmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
