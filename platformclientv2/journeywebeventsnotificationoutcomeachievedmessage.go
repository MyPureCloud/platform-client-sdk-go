package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationoutcomeachievedmessage
type Journeywebeventsnotificationoutcomeachievedmessage struct { 
	// Outcome
	Outcome *Journeywebeventsnotificationoutcome `json:"outcome,omitempty"`


	// Browser
	Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`


	// VisitCreatedDate
	VisitCreatedDate *time.Time `json:"visitCreatedDate,omitempty"`


	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`


	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`


	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`


	// Device
	Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`


	// Geolocation
	Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`


	// MktCampaign
	MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`


	// VisitReferrer
	VisitReferrer *Journeywebeventsnotificationreferrer `json:"visitReferrer,omitempty"`


	// AssociatedValue
	AssociatedValue *Journeywebeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`

}

func (o *Journeywebeventsnotificationoutcomeachievedmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationoutcomeachievedmessage
	
	VisitCreatedDate := new(string)
	if o.VisitCreatedDate != nil {
		
		*VisitCreatedDate = timeutil.Strftime(o.VisitCreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VisitCreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Journeywebeventsnotificationoutcome `json:"outcome,omitempty"`
		
		Browser *Journeywebeventsnotificationbrowser `json:"browser,omitempty"`
		
		VisitCreatedDate *string `json:"visitCreatedDate,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Device *Journeywebeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeywebeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeywebeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeywebeventsnotificationreferrer `json:"visitReferrer,omitempty"`
		
		AssociatedValue *Journeywebeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`
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
		
		AssociatedValue: o.AssociatedValue,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationoutcomeachievedmessage) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationoutcomeachievedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationoutcomeachievedmessageMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if Browser, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if visitCreatedDateString, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["visitCreatedDate"].(string); ok {
		VisitCreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", visitCreatedDateString)
		o.VisitCreatedDate = &VisitCreatedDate
	}
	
	if IpAddress, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if UserAgentString, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Device, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	
	if AssociatedValue, ok := JourneywebeventsnotificationoutcomeachievedmessageMap["associatedValue"].(map[string]interface{}); ok {
		AssociatedValueString, _ := json.Marshal(AssociatedValue)
		json.Unmarshal(AssociatedValueString, &o.AssociatedValue)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationoutcomeachievedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
