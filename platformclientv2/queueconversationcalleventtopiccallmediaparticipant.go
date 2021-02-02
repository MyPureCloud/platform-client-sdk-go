package platformclientv2
import (
	"time"
	"encoding/json"
)

// Queueconversationcalleventtopiccallmediaparticipant
type Queueconversationcalleventtopiccallmediaparticipant struct { 
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
	User *Queueconversationcalleventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationcalleventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationcalleventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationcalleventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationcalleventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationcalleventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationcalleventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationcalleventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcalleventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationcalleventtopicjourneycontext `json:"journeyContext,omitempty"`


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
	Group *Queueconversationcalleventtopicurireference `json:"group,omitempty"`


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
	FaxStatus *Queueconversationcalleventtopicfaxstatus `json:"faxStatus,omitempty"`

}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopiccallmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
