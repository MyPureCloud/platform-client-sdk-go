package platformclientv2
import (
	"time"
	"encoding/json"
)

// Callback
type Callback struct { 
	// State - The connection state of this communication.
	State *string `json:"state,omitempty"`


	// Id - A globally unique identifier for this communication.
	Id *string `json:"id,omitempty"`


	// Segments - The time line of the participant's callback, divided into activity segments.
	Segments *[]Segment `json:"segments,omitempty"`


	// Direction - The direction of the call
	Direction *string `json:"direction,omitempty"`


	// Held - True if this call is held and the person on this side hears silence.
	Held *bool `json:"held,omitempty"`


	// DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime - The timestamp the callback was placed on hold in the cloud clock if the callback is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// DialerPreview - The preview data to be used when this callback is a Preview.
	DialerPreview *Dialerpreview `json:"dialerPreview,omitempty"`


	// Voicemail - The voicemail data to be used when this callback is an ACD voicemail.
	Voicemail *Voicemail `json:"voicemail,omitempty"`


	// CallbackNumbers - The phone number(s) to use to place the callback.
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName - The name of the user requesting a callback.
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// ScriptId - The UUID of the script to use.
	ScriptId *string `json:"scriptId,omitempty"`


	// ExternalCampaign - True if the call for the callback uses external dialing.
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// SkipEnabled - True if the ability to skip a callback should be enabled.
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// TimeoutSeconds - The number of seconds before the system automatically places a call for a callback.  0 means the automatic placement is disabled.
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`


	// StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAlertingTime *time.Time `json:"startAlertingTime,omitempty"`


	// ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// CallbackScheduledTime - The timestamp when this communication is scheduled in the provider clock. If this value is missing it indicates the callback will be placed immediately. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// AutomatedCallbackConfigId - The id of the config for automatically placing the callback (and handling the disposition). If null, the callback will not be placed automatically but routed to an agent as per normal.
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`


	// Provider - The source provider for the callback.
	Provider *string `json:"provider,omitempty"`


	// PeerId - The id of the peer communication corresponding to a matching leg for this communication.
	PeerId *string `json:"peerId,omitempty"`


	// Wrapup - Call wrap up or disposition data.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// AfterCallWork - After-call work for the communication.
	AfterCallWork *Aftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`

}

// String returns a JSON representation of the model
func (o *Callback) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
