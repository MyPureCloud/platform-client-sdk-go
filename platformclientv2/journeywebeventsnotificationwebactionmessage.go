package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebactionmessage
type Journeywebeventsnotificationwebactionmessage struct { 
	// Action
	Action *Journeywebeventsnotificationeventaction `json:"action,omitempty"`


	// ActionTarget
	ActionTarget *Journeywebeventsnotificationactiontarget `json:"actionTarget,omitempty"`


	// ActionMap
	ActionMap *Journeywebeventsnotificationactionmap `json:"actionMap,omitempty"`


	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`


	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`


	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`


	// Browser
	Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`


	// Device
	Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`


	// MktCampaign
	MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// VisitReferrer
	VisitReferrer *Journeywebeventsnotificationreferrer `json:"visitReferrer,omitempty"`

}

func (o *Journeywebeventsnotificationwebactionmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationwebactionmessage
	
	return json.Marshal(&struct { 
		Action *Journeywebeventsnotificationeventaction `json:"action,omitempty"`
		
		ActionTarget *Journeywebeventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		ActionMap *Journeywebeventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeywebeventsnotificationreferrer `json:"visitReferrer,omitempty"`
		*Alias
	}{ 
		Action: o.Action,
		
		ActionTarget: o.ActionTarget,
		
		ActionMap: o.ActionMap,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		
		UserAgentString: o.UserAgentString,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		MktCampaign: o.MktCampaign,
		
		VisitReferrer: o.VisitReferrer,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationwebactionmessage) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationwebactionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationwebactionmessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneywebeventsnotificationwebactionmessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionTarget, ok := JourneywebeventsnotificationwebactionmessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if ActionMap, ok := JourneywebeventsnotificationwebactionmessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ErrorCode, ok := JourneywebeventsnotificationwebactionmessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
	
	if ErrorMessage, ok := JourneywebeventsnotificationwebactionmessageMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
	
	if UserAgentString, ok := JourneywebeventsnotificationwebactionmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
	
	if IpAddress, ok := JourneywebeventsnotificationwebactionmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
	
	if IpOrganization, ok := JourneywebeventsnotificationwebactionmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
	
	if Browser, ok := JourneywebeventsnotificationwebactionmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneywebeventsnotificationwebactionmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneywebeventsnotificationwebactionmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneywebeventsnotificationwebactionmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneywebeventsnotificationwebactionmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationwebactionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
