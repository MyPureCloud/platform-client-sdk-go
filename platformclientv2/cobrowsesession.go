package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Cobrowsesession
type Cobrowsesession struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Self - Address and name data for a call endpoint.
	Self *Address `json:"self,omitempty"`


	// CobrowseSessionId - The co-browse session ID.
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`


	// CobrowseRole - This value identifies the role of the co-browse client within the co-browse session (a client is a sharer or a viewer).
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// Controlling - ID of co-browse participants for which this client has been granted control (list is empty if this client cannot control any shared pages).
	Controlling *[]string `json:"controlling,omitempty"`


	// ViewerUrl - The URL that can be used to open co-browse session in web browser.
	ViewerUrl *string `json:"viewerUrl,omitempty"`


	// ProviderEventTime - The time when the provider event which triggered this conversation update happened in the corrected provider clock (milliseconds since 1970-01-01 00:00:00 UTC). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Provider - The source provider for the co-browse session.
	Provider *string `json:"provider,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// Segments - The time line of the participant's call, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

func (o *Cobrowsesession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Cobrowsesession
	
	ProviderEventTime := new(string)
	if o.ProviderEventTime != nil {
		
		*ProviderEventTime = timeutil.Strftime(o.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ProviderEventTime = nil
	}
	
	StartAlertingTime := new(string)
	if o.StartAlertingTime != nil {
		
		*StartAlertingTime = timeutil.Strftime(o.StartAlertingTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAlertingTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Self *Address `json:"self,omitempty"`
		
		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		Controlling *[]string `json:"controlling,omitempty"`
		
		ViewerUrl *string `json:"viewerUrl,omitempty"`
		
		ProviderEventTime *string `json:"providerEventTime,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		InitialState: o.InitialState,
		
		Id: o.Id,
		
		DisconnectType: o.DisconnectType,
		
		Self: o.Self,
		
		CobrowseSessionId: o.CobrowseSessionId,
		
		CobrowseRole: o.CobrowseRole,
		
		Controlling: o.Controlling,
		
		ViewerUrl: o.ViewerUrl,
		
		ProviderEventTime: ProviderEventTime,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Provider: o.Provider,
		
		PeerId: o.PeerId,
		
		Segments: o.Segments,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		Alias:    (*Alias)(o),
	})
}

func (o *Cobrowsesession) UnmarshalJSON(b []byte) error {
	var CobrowsesessionMap map[string]interface{}
	err := json.Unmarshal(b, &CobrowsesessionMap)
	if err != nil {
		return err
	}
	
	if State, ok := CobrowsesessionMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := CobrowsesessionMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Id, ok := CobrowsesessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DisconnectType, ok := CobrowsesessionMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Self, ok := CobrowsesessionMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if CobrowseSessionId, ok := CobrowsesessionMap["cobrowseSessionId"].(string); ok {
		o.CobrowseSessionId = &CobrowseSessionId
	}
    
	if CobrowseRole, ok := CobrowsesessionMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}
    
	if Controlling, ok := CobrowsesessionMap["controlling"].([]interface{}); ok {
		ControllingString, _ := json.Marshal(Controlling)
		json.Unmarshal(ControllingString, &o.Controlling)
	}
	
	if ViewerUrl, ok := CobrowsesessionMap["viewerUrl"].(string); ok {
		o.ViewerUrl = &ViewerUrl
	}
    
	if providerEventTimeString, ok := CobrowsesessionMap["providerEventTime"].(string); ok {
		ProviderEventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", providerEventTimeString)
		o.ProviderEventTime = &ProviderEventTime
	}
	
	if startAlertingTimeString, ok := CobrowsesessionMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}
	
	if connectedTimeString, ok := CobrowsesessionMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := CobrowsesessionMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Provider, ok := CobrowsesessionMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if PeerId, ok := CobrowsesessionMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if Segments, ok := CobrowsesessionMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	
	if Wrapup, ok := CobrowsesessionMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := CobrowsesessionMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := CobrowsesessionMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Cobrowsesession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
