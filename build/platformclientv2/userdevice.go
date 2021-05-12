package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userdevice
type Userdevice struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DeviceToken - device token sent by mobile clients.
	DeviceToken *string `json:"deviceToken,omitempty"`


	// NotificationId - notification id of the device.
	NotificationId *string `json:"notificationId,omitempty"`


	// Make - make of the device.
	Make *string `json:"make,omitempty"`


	// Model - Device model
	Model *string `json:"model,omitempty"`


	// AcceptNotifications - if the device accepts notifications
	AcceptNotifications *bool `json:"acceptNotifications,omitempty"`


	// VarType - type of the device; ios or android
	VarType *string `json:"type,omitempty"`


	// SessionHash
	SessionHash *string `json:"sessionHash,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
