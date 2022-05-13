package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationoutcomeachievedmessage
type Journeyoutcomeeventsnotificationoutcomeachievedmessage struct { 
	// Outcome
	Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`


	// Browser
	Browser *Journeyoutcomeeventsnotificationbrowser `json:"browser,omitempty"`


	// VisitCreatedDate
	VisitCreatedDate *time.Time `json:"visitCreatedDate,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`


	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`


	// Device
	Device *Journeyoutcomeeventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeyoutcomeeventsnotificationgeolocation `json:"geolocation,omitempty"`


	// MktCampaign
	MktCampaign *Journeyoutcomeeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// VisitReferrer
	VisitReferrer *Journeyoutcomeeventsnotificationreferrer `json:"visitReferrer,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationoutcomeachievedmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationoutcomeachievedmessage
	
	VisitCreatedDate := new(string)
	if o.VisitCreatedDate != nil {
		
		*VisitCreatedDate = timeutil.Strftime(o.VisitCreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VisitCreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Journeyoutcomeeventsnotificationoutcome `json:"outcome,omitempty"`
		
		Browser *Journeyoutcomeeventsnotificationbrowser `json:"browser,omitempty"`
		
		VisitCreatedDate *string `json:"visitCreatedDate,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Device *Journeyoutcomeeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeyoutcomeeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeyoutcomeeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeyoutcomeeventsnotificationreferrer `json:"visitReferrer,omitempty"`
		*Alias
	}{ 
		Outcome: o.Outcome,
		
		Browser: o.Browser,
		
		VisitCreatedDate: VisitCreatedDate,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		UserAgentString: o.UserAgentString,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		MktCampaign: o.MktCampaign,
		
		VisitReferrer: o.VisitReferrer,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationoutcomeachievedmessage) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationoutcomeachievedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationoutcomeachievedmessageMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if Browser, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if visitCreatedDateString, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["visitCreatedDate"].(string); ok {
		VisitCreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", visitCreatedDateString)
		o.VisitCreatedDate = &VisitCreatedDate
	}
	
	if IpAddress, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if UserAgentString, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Device, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneyoutcomeeventsnotificationoutcomeachievedmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationoutcomeachievedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
