package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebmessage
type Journeywebeventsnotificationwebmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebeventsnotificationwebmessage) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Journeywebeventsnotificationwebmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
