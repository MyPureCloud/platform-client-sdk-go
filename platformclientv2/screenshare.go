package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Screenshare
type Screenshare struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Context - The room id context (xmpp jid) for the conference session.
	Context *string `json:"context,omitempty"`


	// Sharing - Indicates whether this participant is sharing their screen.
	Sharing *bool `json:"sharing,omitempty"`


	// PeerCount - The number of peer participants from the perspective of the participant in the conference.
	PeerCount *int `json:"peerCount,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Provider - The source provider for the screen share.
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


	// InitialState - The initial connection state of this communication.
	InitialState *string `json:"initialState,omitempty"`

}

func (o *Screenshare) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Screenshare
	
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
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		Sharing *bool `json:"sharing,omitempty"`
		
		PeerCount *int `json:"peerCount,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartAlertingTime *string `json:"startAlertingTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		Segments *[]Segment `json:"segments,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		InitialState *string `json:"initialState,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Id: o.Id,
		
		Context: o.Context,
		
		Sharing: o.Sharing,
		
		PeerCount: o.PeerCount,
		
		DisconnectType: o.DisconnectType,
		
		StartAlertingTime: StartAlertingTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Provider: o.Provider,
		
		PeerId: o.PeerId,
		
		Segments: o.Segments,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		InitialState: o.InitialState,
		Alias:    (*Alias)(o),
	})
}

func (o *Screenshare) UnmarshalJSON(b []byte) error {
	var ScreenshareMap map[string]interface{}
	err := json.Unmarshal(b, &ScreenshareMap)
	if err != nil {
		return err
	}
	
	if State, ok := ScreenshareMap["state"].(string); ok {
		o.State = &State
	}
    
	if Id, ok := ScreenshareMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Context, ok := ScreenshareMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if Sharing, ok := ScreenshareMap["sharing"].(bool); ok {
		o.Sharing = &Sharing
	}
    
	if PeerCount, ok := ScreenshareMap["peerCount"].(float64); ok {
		PeerCountInt := int(PeerCount)
		o.PeerCount = &PeerCountInt
	}
	
	if DisconnectType, ok := ScreenshareMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if startAlertingTimeString, ok := ScreenshareMap["startAlertingTime"].(string); ok {
		StartAlertingTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAlertingTimeString)
		o.StartAlertingTime = &StartAlertingTime
	}
	
	if connectedTimeString, ok := ScreenshareMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := ScreenshareMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Provider, ok := ScreenshareMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if PeerId, ok := ScreenshareMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
    
	if Segments, ok := ScreenshareMap["segments"].([]interface{}); ok {
		SegmentsString, _ := json.Marshal(Segments)
		json.Unmarshal(SegmentsString, &o.Segments)
	}
	
	if Wrapup, ok := ScreenshareMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := ScreenshareMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := ScreenshareMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
    
	if InitialState, ok := ScreenshareMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Screenshare) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
