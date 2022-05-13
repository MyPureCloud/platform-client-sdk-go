package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebmessage
type Journeywebeventsnotificationwebmessage struct { 
	// EventName
	EventName *string `json:"eventName,omitempty"`


	// TotalEventCount
	TotalEventCount *int `json:"totalEventCount,omitempty"`


	// TotalPageviewCount
	TotalPageviewCount *int `json:"totalPageviewCount,omitempty"`


	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`


	// SearchQuery
	SearchQuery *string `json:"searchQuery,omitempty"`


	// Authenticated
	Authenticated *bool `json:"authenticated,omitempty"`


	// Browser
	Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`


	// Device
	Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`


	// MktCampaign
	MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// Page
	Page *Journeywebeventsnotificationpage `json:"page,omitempty"`


	// Referrer
	Referrer *Journeywebeventsnotificationreferrer `json:"referrer,omitempty"`


	// Attributes
	Attributes *map[string]Journeywebeventsnotificationcustomeventattribute `json:"attributes,omitempty"`


	// Traits
	Traits *map[string]Journeywebeventsnotificationcustomeventattribute `json:"traits,omitempty"`

}

func (o *Journeywebeventsnotificationwebmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationwebmessage
	
	return json.Marshal(&struct { 
		EventName *string `json:"eventName,omitempty"`
		
		TotalEventCount *int `json:"totalEventCount,omitempty"`
		
		TotalPageviewCount *int `json:"totalPageviewCount,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		SearchQuery *string `json:"searchQuery,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		
		Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		Page *Journeywebeventsnotificationpage `json:"page,omitempty"`
		
		Referrer *Journeywebeventsnotificationreferrer `json:"referrer,omitempty"`
		
		Attributes *map[string]Journeywebeventsnotificationcustomeventattribute `json:"attributes,omitempty"`
		
		Traits *map[string]Journeywebeventsnotificationcustomeventattribute `json:"traits,omitempty"`
		*Alias
	}{ 
		EventName: o.EventName,
		
		TotalEventCount: o.TotalEventCount,
		
		TotalPageviewCount: o.TotalPageviewCount,
		
		UserAgentString: o.UserAgentString,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		SearchQuery: o.SearchQuery,
		
		Authenticated: o.Authenticated,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		MktCampaign: o.MktCampaign,
		
		Page: o.Page,
		
		Referrer: o.Referrer,
		
		Attributes: o.Attributes,
		
		Traits: o.Traits,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationwebmessage) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationwebmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationwebmessageMap)
	if err != nil {
		return err
	}
	
	if EventName, ok := JourneywebeventsnotificationwebmessageMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if TotalEventCount, ok := JourneywebeventsnotificationwebmessageMap["totalEventCount"].(float64); ok {
		TotalEventCountInt := int(TotalEventCount)
		o.TotalEventCount = &TotalEventCountInt
	}
	
	if TotalPageviewCount, ok := JourneywebeventsnotificationwebmessageMap["totalPageviewCount"].(float64); ok {
		TotalPageviewCountInt := int(TotalPageviewCount)
		o.TotalPageviewCount = &TotalPageviewCountInt
	}
	
	if UserAgentString, ok := JourneywebeventsnotificationwebmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if IpAddress, ok := JourneywebeventsnotificationwebmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneywebeventsnotificationwebmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if SearchQuery, ok := JourneywebeventsnotificationwebmessageMap["searchQuery"].(string); ok {
		o.SearchQuery = &SearchQuery
	}
    
	if Authenticated, ok := JourneywebeventsnotificationwebmessageMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    
	if Browser, ok := JourneywebeventsnotificationwebmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneywebeventsnotificationwebmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneywebeventsnotificationwebmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneywebeventsnotificationwebmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if Page, ok := JourneywebeventsnotificationwebmessageMap["page"].(map[string]interface{}); ok {
		PageString, _ := json.Marshal(Page)
		json.Unmarshal(PageString, &o.Page)
	}
	
	if Referrer, ok := JourneywebeventsnotificationwebmessageMap["referrer"].(map[string]interface{}); ok {
		ReferrerString, _ := json.Marshal(Referrer)
		json.Unmarshal(ReferrerString, &o.Referrer)
	}
	
	if Attributes, ok := JourneywebeventsnotificationwebmessageMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Traits, ok := JourneywebeventsnotificationwebmessageMap["traits"].(map[string]interface{}); ok {
		TraitsString, _ := json.Marshal(Traits)
		json.Unmarshal(TraitsString, &o.Traits)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationwebmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
