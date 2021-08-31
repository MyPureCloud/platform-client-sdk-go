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

func (o *Userdevice) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		DeviceToken: o.DeviceToken,
		
		NotificationId: o.NotificationId,
		
		Make: o.Make,
		
		Model: o.Model,
		
		AcceptNotifications: o.AcceptNotifications,
		
		VarType: o.VarType,
		
		SessionHash: o.SessionHash,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userdevice) UnmarshalJSON(b []byte) error {
	var UserdeviceMap map[string]interface{}
	err := json.Unmarshal(b, &UserdeviceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserdeviceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := UserdeviceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if DeviceToken, ok := UserdeviceMap["deviceToken"].(string); ok {
		o.DeviceToken = &DeviceToken
	}
	
	if NotificationId, ok := UserdeviceMap["notificationId"].(string); ok {
		o.NotificationId = &NotificationId
	}
	
	if Make, ok := UserdeviceMap["make"].(string); ok {
		o.Make = &Make
	}
	
	if Model, ok := UserdeviceMap["model"].(string); ok {
		o.Model = &Model
	}
	
	if AcceptNotifications, ok := UserdeviceMap["acceptNotifications"].(bool); ok {
		o.AcceptNotifications = &AcceptNotifications
	}
	
	if VarType, ok := UserdeviceMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if SessionHash, ok := UserdeviceMap["sessionHash"].(string); ok {
		o.SessionHash = &SessionHash
	}
	
	if SelfUri, ok := UserdeviceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userdevice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
