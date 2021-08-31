package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopiccobrowse
type Queueconversationeventtopiccobrowse struct { 
	// State
	State *string `json:"state,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Self
	Self *Queueconversationeventtopicaddress `json:"self,omitempty"`


	// RoomId
	RoomId *string `json:"roomId,omitempty"`


	// CobrowseSessionId
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`


	// CobrowseRole
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// Controlling
	Controlling *[]string `json:"controlling,omitempty"`


	// ViewerUrl
	ViewerUrl *string `json:"viewerUrl,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// ProviderEventTime
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Wrapup
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopiccobrowse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopiccobrowse
	
	ProviderEventTime := new(string)
	if o.ProviderEventTime != nil {
		
		*ProviderEventTime = timeutil.Strftime(o.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ProviderEventTime = nil
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
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Self *Queueconversationeventtopicaddress `json:"self,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		Controlling *[]string `json:"controlling,omitempty"`
		
		ViewerUrl *string `json:"viewerUrl,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ProviderEventTime *string `json:"providerEventTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		DisconnectType: o.DisconnectType,
		
		Id: o.Id,
		
		Self: o.Self,
		
		RoomId: o.RoomId,
		
		CobrowseSessionId: o.CobrowseSessionId,
		
		CobrowseRole: o.CobrowseRole,
		
		Controlling: o.Controlling,
		
		ViewerUrl: o.ViewerUrl,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		ProviderEventTime: ProviderEventTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopiccobrowse) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopiccobrowseMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopiccobrowseMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationeventtopiccobrowseMap["state"].(string); ok {
		o.State = &State
	}
	
	if DisconnectType, ok := QueueconversationeventtopiccobrowseMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if Id, ok := QueueconversationeventtopiccobrowseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Self, ok := QueueconversationeventtopiccobrowseMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if RoomId, ok := QueueconversationeventtopiccobrowseMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
	
	if CobrowseSessionId, ok := QueueconversationeventtopiccobrowseMap["cobrowseSessionId"].(string); ok {
		o.CobrowseSessionId = &CobrowseSessionId
	}
	
	if CobrowseRole, ok := QueueconversationeventtopiccobrowseMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}
	
	if Controlling, ok := QueueconversationeventtopiccobrowseMap["controlling"].([]interface{}); ok {
		ControllingString, _ := json.Marshal(Controlling)
		json.Unmarshal(ControllingString, &o.Controlling)
	}
	
	if ViewerUrl, ok := QueueconversationeventtopiccobrowseMap["viewerUrl"].(string); ok {
		o.ViewerUrl = &ViewerUrl
	}
	
	if Provider, ok := QueueconversationeventtopiccobrowseMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationeventtopiccobrowseMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationeventtopiccobrowseMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if providerEventTimeString, ok := QueueconversationeventtopiccobrowseMap["providerEventTime"].(string); ok {
		ProviderEventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", providerEventTimeString)
		o.ProviderEventTime = &ProviderEventTime
	}
	
	if connectedTimeString, ok := QueueconversationeventtopiccobrowseMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationeventtopiccobrowseMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationeventtopiccobrowseMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationeventtopiccobrowseMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationeventtopiccobrowseMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationeventtopiccobrowseMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopiccobrowse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
