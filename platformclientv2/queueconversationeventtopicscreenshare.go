package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicscreenshare
type Queueconversationeventtopicscreenshare struct { 
	// State
	State *string `json:"state,omitempty"`


	// Self
	Self *Queueconversationeventtopicaddress `json:"self,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Context
	Context *string `json:"context,omitempty"`


	// Sharing
	Sharing *bool `json:"sharing,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


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

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicscreenshare) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
