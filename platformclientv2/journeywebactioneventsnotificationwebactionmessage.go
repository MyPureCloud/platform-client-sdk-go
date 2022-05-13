package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationwebactionmessage
type Journeywebactioneventsnotificationwebactionmessage struct { 
	// Action
	Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`


	// ActionTarget
	ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`


	// ActionMap
	ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`


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
	Browser *Journeywebactioneventsnotificationbrowser `json:"browser,omitempty"`


	// Device
	Device *Journeywebactioneventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeywebactioneventsnotificationgeolocation `json:"geolocation,omitempty"`


	// MktCampaign
	MktCampaign *Journeywebactioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// VisitReferrer
	VisitReferrer *Journeywebactioneventsnotificationreferrer `json:"visitReferrer,omitempty"`

}

func (o *Journeywebactioneventsnotificationwebactionmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationwebactionmessage
	
	return json.Marshal(&struct { 
		Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`
		
		ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		Browser *Journeywebactioneventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeywebactioneventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeywebactioneventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeywebactioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeywebactioneventsnotificationreferrer `json:"visitReferrer,omitempty"`
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

func (o *Journeywebactioneventsnotificationwebactionmessage) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationwebactionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationwebactionmessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneywebactioneventsnotificationwebactionmessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionTarget, ok := JourneywebactioneventsnotificationwebactionmessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if ActionMap, ok := JourneywebactioneventsnotificationwebactionmessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ErrorCode, ok := JourneywebactioneventsnotificationwebactionmessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := JourneywebactioneventsnotificationwebactionmessageMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if UserAgentString, ok := JourneywebactioneventsnotificationwebactionmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if IpAddress, ok := JourneywebactioneventsnotificationwebactionmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneywebactioneventsnotificationwebactionmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if Browser, ok := JourneywebactioneventsnotificationwebactionmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneywebactioneventsnotificationwebactionmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneywebactioneventsnotificationwebactionmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneywebactioneventsnotificationwebactionmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneywebactioneventsnotificationwebactionmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationwebactionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
