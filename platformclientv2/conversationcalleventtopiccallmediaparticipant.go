package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcalleventtopiccallmediaparticipant
type Conversationcalleventtopiccallmediaparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

	// InitialState
	InitialState *string `json:"initialState,omitempty"`

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

	// ExternalContactInitialDivisionId
	ExternalContactInitialDivisionId *string `json:"externalContactInitialDivisionId,omitempty"`

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

	// ResumeTime
	ResumeTime *time.Time `json:"resumeTime,omitempty"`

	// ParkTime
	ParkTime *time.Time `json:"parkTime,omitempty"`

	// MediaRoles
	MediaRoles *[]string `json:"mediaRoles,omitempty"`

	// QueueMediaSettings
	QueueMediaSettings *Conversationcalleventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`

	// Muted
	Muted *bool `json:"muted,omitempty"`

	// Confined
	Confined *bool `json:"confined,omitempty"`

	// Recording
	Recording *bool `json:"recording,omitempty"`

	// RecordingState
	RecordingState *string `json:"recordingState,omitempty"`

	// RecordersState
	RecordersState *Conversationcalleventtopicrecordersstate `json:"recordersState,omitempty"`

	// Disposition
	Disposition *Conversationcalleventtopicdisposition `json:"disposition,omitempty"`

	// SecurePause
	SecurePause *bool `json:"securePause,omitempty"`

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

	// CoachedParticipantId
	CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`

	// BargedParticipantId
	BargedParticipantId *string `json:"bargedParticipantId,omitempty"`

	// BargedTime
	BargedTime *time.Time `json:"bargedTime,omitempty"`

	// ConsultParticipantId
	ConsultParticipantId *string `json:"consultParticipantId,omitempty"`

	// FaxStatus
	FaxStatus *Conversationcalleventtopicfaxstatus `json:"faxStatus,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcalleventtopiccallmediaparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcalleventtopiccallmediaparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime","ConnectedTime","EndTime","StartHoldTime","StartAcwTime","EndAcwTime","ResumeTime","ParkTime","BargedTime", }
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
	type Alias Conversationcalleventtopiccallmediaparticipant
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAcwTime := new(string)
	if o.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(o.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if o.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(o.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndAcwTime = nil
	}
	
	ResumeTime := new(string)
	if o.ResumeTime != nil {
		
		*ResumeTime = timeutil.Strftime(o.ResumeTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ResumeTime = nil
	}
	
	ParkTime := new(string)
	if o.ParkTime != nil {
		
		*ParkTime = timeutil.Strftime(o.ParkTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ParkTime = nil
	}
	
	BargedTime := new(string)
	if o.BargedTime != nil {
		
		*BargedTime = timeutil.Strftime(o.BargedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		BargedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Conversationcalleventtopicurireference `json:"user,omitempty"`
		
		Queue *Conversationcalleventtopicurireference `json:"queue,omitempty"`
		
		Team *Conversationcalleventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Conversationcalleventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Conversationcalleventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Conversationcalleventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalContactInitialDivisionId *string `json:"externalContactInitialDivisionId,omitempty"`
		
		ExternalOrganization *Conversationcalleventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Conversationcalleventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Conversationcalleventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Conversationcalleventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		ResumeTime *string `json:"resumeTime,omitempty"`
		
		ParkTime *string `json:"parkTime,omitempty"`
		
		MediaRoles *[]string `json:"mediaRoles,omitempty"`
		
		QueueMediaSettings *Conversationcalleventtopicqueuemediasettings `json:"queueMediaSettings,omitempty"`
		
		Muted *bool `json:"muted,omitempty"`
		
		Confined *bool `json:"confined,omitempty"`
		
		Recording *bool `json:"recording,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		RecordersState *Conversationcalleventtopicrecordersstate `json:"recordersState,omitempty"`
		
		Disposition *Conversationcalleventtopicdisposition `json:"disposition,omitempty"`
		
		SecurePause *bool `json:"securePause,omitempty"`
		
		Group *Conversationcalleventtopicurireference `json:"group,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		DocumentId *string `json:"documentId,omitempty"`
		
		MonitoredParticipantId *string `json:"monitoredParticipantId,omitempty"`
		
		CoachedParticipantId *string `json:"coachedParticipantId,omitempty"`
		
		BargedParticipantId *string `json:"bargedParticipantId,omitempty"`
		
		BargedTime *string `json:"bargedTime,omitempty"`
		
		ConsultParticipantId *string `json:"consultParticipantId,omitempty"`
		
		FaxStatus *Conversationcalleventtopicfaxstatus `json:"faxStatus,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Address: o.Address,
		
		StartTime: StartTime,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		StartHoldTime: StartHoldTime,
		
		Purpose: o.Purpose,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Direction: o.Direction,
		
		DisconnectType: o.DisconnectType,
		
		Held: o.Held,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupPrompt: o.WrapupPrompt,
		
		User: o.User,
		
		Queue: o.Queue,
		
		Team: o.Team,
		
		Attributes: o.Attributes,
		
		ErrorInfo: o.ErrorInfo,
		
		Script: o.Script,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		WrapupSkipped: o.WrapupSkipped,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		Provider: o.Provider,
		
		ExternalContact: o.ExternalContact,
		
		ExternalContactInitialDivisionId: o.ExternalContactInitialDivisionId,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Wrapup: o.Wrapup,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		Peer: o.Peer,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ResumeTime: ResumeTime,
		
		ParkTime: ParkTime,
		
		MediaRoles: o.MediaRoles,
		
		QueueMediaSettings: o.QueueMediaSettings,
		
		Muted: o.Muted,
		
		Confined: o.Confined,
		
		Recording: o.Recording,
		
		RecordingState: o.RecordingState,
		
		RecordersState: o.RecordersState,
		
		Disposition: o.Disposition,
		
		SecurePause: o.SecurePause,
		
		Group: o.Group,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		DocumentId: o.DocumentId,
		
		MonitoredParticipantId: o.MonitoredParticipantId,
		
		CoachedParticipantId: o.CoachedParticipantId,
		
		BargedParticipantId: o.BargedParticipantId,
		
		BargedTime: BargedTime,
		
		ConsultParticipantId: o.ConsultParticipantId,
		
		FaxStatus: o.FaxStatus,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcalleventtopiccallmediaparticipant) UnmarshalJSON(b []byte) error {
	var ConversationcalleventtopiccallmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcalleventtopiccallmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcalleventtopiccallmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationcalleventtopiccallmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := ConversationcalleventtopiccallmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := ConversationcalleventtopiccallmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := ConversationcalleventtopiccallmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationcalleventtopiccallmediaparticipantMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := ConversationcalleventtopiccallmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := ConversationcalleventtopiccallmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := ConversationcalleventtopiccallmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := ConversationcalleventtopiccallmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := ConversationcalleventtopiccallmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if User, ok := ConversationcalleventtopiccallmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := ConversationcalleventtopiccallmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := ConversationcalleventtopiccallmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := ConversationcalleventtopiccallmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := ConversationcalleventtopiccallmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := ConversationcalleventtopiccallmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := ConversationcalleventtopiccallmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ConversationcalleventtopiccallmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := ConversationcalleventtopiccallmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := ConversationcalleventtopiccallmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := ConversationcalleventtopiccallmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalContactInitialDivisionId, ok := ConversationcalleventtopiccallmediaparticipantMap["externalContactInitialDivisionId"].(string); ok {
		o.ExternalContactInitialDivisionId = &ExternalContactInitialDivisionId
	}
    
	if ExternalOrganization, ok := ConversationcalleventtopiccallmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := ConversationcalleventtopiccallmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := ConversationcalleventtopiccallmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := ConversationcalleventtopiccallmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if ScreenRecordingState, ok := ConversationcalleventtopiccallmediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := ConversationcalleventtopiccallmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := ConversationcalleventtopiccallmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if resumeTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["resumeTime"].(string); ok {
		ResumeTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", resumeTimeString)
		o.ResumeTime = &ResumeTime
	}
	
	if parkTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["parkTime"].(string); ok {
		ParkTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", parkTimeString)
		o.ParkTime = &ParkTime
	}
	
	if MediaRoles, ok := ConversationcalleventtopiccallmediaparticipantMap["mediaRoles"].([]interface{}); ok {
		MediaRolesString, _ := json.Marshal(MediaRoles)
		json.Unmarshal(MediaRolesString, &o.MediaRoles)
	}
	
	if QueueMediaSettings, ok := ConversationcalleventtopiccallmediaparticipantMap["queueMediaSettings"].(map[string]interface{}); ok {
		QueueMediaSettingsString, _ := json.Marshal(QueueMediaSettings)
		json.Unmarshal(QueueMediaSettingsString, &o.QueueMediaSettings)
	}
	
	if Muted, ok := ConversationcalleventtopiccallmediaparticipantMap["muted"].(bool); ok {
		o.Muted = &Muted
	}
    
	if Confined, ok := ConversationcalleventtopiccallmediaparticipantMap["confined"].(bool); ok {
		o.Confined = &Confined
	}
    
	if Recording, ok := ConversationcalleventtopiccallmediaparticipantMap["recording"].(bool); ok {
		o.Recording = &Recording
	}
    
	if RecordingState, ok := ConversationcalleventtopiccallmediaparticipantMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if RecordersState, ok := ConversationcalleventtopiccallmediaparticipantMap["recordersState"].(map[string]interface{}); ok {
		RecordersStateString, _ := json.Marshal(RecordersState)
		json.Unmarshal(RecordersStateString, &o.RecordersState)
	}
	
	if Disposition, ok := ConversationcalleventtopiccallmediaparticipantMap["disposition"].(map[string]interface{}); ok {
		DispositionString, _ := json.Marshal(Disposition)
		json.Unmarshal(DispositionString, &o.Disposition)
	}
	
	if SecurePause, ok := ConversationcalleventtopiccallmediaparticipantMap["securePause"].(bool); ok {
		o.SecurePause = &SecurePause
	}
    
	if Group, ok := ConversationcalleventtopiccallmediaparticipantMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if Ani, ok := ConversationcalleventtopiccallmediaparticipantMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := ConversationcalleventtopiccallmediaparticipantMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if DocumentId, ok := ConversationcalleventtopiccallmediaparticipantMap["documentId"].(string); ok {
		o.DocumentId = &DocumentId
	}
    
	if MonitoredParticipantId, ok := ConversationcalleventtopiccallmediaparticipantMap["monitoredParticipantId"].(string); ok {
		o.MonitoredParticipantId = &MonitoredParticipantId
	}
    
	if CoachedParticipantId, ok := ConversationcalleventtopiccallmediaparticipantMap["coachedParticipantId"].(string); ok {
		o.CoachedParticipantId = &CoachedParticipantId
	}
    
	if BargedParticipantId, ok := ConversationcalleventtopiccallmediaparticipantMap["bargedParticipantId"].(string); ok {
		o.BargedParticipantId = &BargedParticipantId
	}
    
	if bargedTimeString, ok := ConversationcalleventtopiccallmediaparticipantMap["bargedTime"].(string); ok {
		BargedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", bargedTimeString)
		o.BargedTime = &BargedTime
	}
	
	if ConsultParticipantId, ok := ConversationcalleventtopiccallmediaparticipantMap["consultParticipantId"].(string); ok {
		o.ConsultParticipantId = &ConsultParticipantId
	}
    
	if FaxStatus, ok := ConversationcalleventtopiccallmediaparticipantMap["faxStatus"].(map[string]interface{}); ok {
		FaxStatusString, _ := json.Marshal(FaxStatus)
		json.Unmarshal(FaxStatusString, &o.FaxStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcalleventtopiccallmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
