package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Userdevice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userdevice

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DeviceToken *string `json:"deviceToken,omitempty"`
		
		NotificationId *string `json:"notificationId,omitempty"`
		
		Make *string `json:"make,omitempty"`
		
		Model *string `json:"model,omitempty"`
		
		AcceptNotifications *bool `json:"acceptNotifications,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		SessionHash *string `json:"sessionHash,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DeviceToken: u.DeviceToken,
		
		NotificationId: u.NotificationId,
		
		Make: u.Make,
		
		Model: u.Model,
		
		AcceptNotifications: u.AcceptNotifications,
		
		VarType: u.VarType,
		
		SessionHash: u.SessionHash,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
