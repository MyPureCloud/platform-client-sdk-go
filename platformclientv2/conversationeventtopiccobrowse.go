package platformclientv2
import (
	"time"
	"encoding/json"
)

// Conversationeventtopiccobrowse
type Conversationeventtopiccobrowse struct { 
	// State
	State *string `json:"state,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Self
	Self *Conversationeventtopicaddress `json:"self,omitempty"`


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
	Wrapup *Conversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Conversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationeventtopiccobrowse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
