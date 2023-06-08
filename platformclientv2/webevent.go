package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webevent
type Webevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventName - Represents the action the customer performed. A good event name is typically an object followed by the action performed in past tense (e.g. page_viewed, order_completed, user_registered).
	EventName *string `json:"eventName,omitempty"`

	// TotalEventCount - The total count of events performed by the customer across all sessions.
	TotalEventCount *int `json:"totalEventCount,omitempty"`

	// TotalPageviewCount - The total count of pageviews performed by the customer across all sessions.
	TotalPageviewCount *int `json:"totalPageviewCount,omitempty"`

	// Page - The webpage where the user interaction occurred.
	Page *Journeypage `json:"page,omitempty"`

	// UserAgentString - HTTP User-Agent string (see https://tools.ietf.org/html/rfc1945#section-10.15).
	UserAgentString *string `json:"userAgentString,omitempty"`

	// Browser - Customer's browser.
	Browser *Browser `json:"browser,omitempty"`

	// Device - Customer's device.
	Device *Device `json:"device,omitempty"`

	// Geolocation - Customer's geolocation.
	Geolocation *Journeygeolocation `json:"geolocation,omitempty"`

	// IpAddress - Customer's IP address. May be null if the business configures the tracker to not collect IP addresses.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpOrganization - Customer's IP-based organization or ISP name.
	IpOrganization *string `json:"ipOrganization,omitempty"`

	// MktCampaign - Marketing / traffic source information.
	MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`

	// Referrer - Identifies the page URL that originally generated the request for the current page being viewed.
	Referrer *Referrer `json:"referrer,omitempty"`

	// Attributes - User-defined attributes associated with a particular event.
	Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`

	// Traits - User-defined traits associated with a particular event.
	Traits *map[string]Customeventattribute `json:"traits,omitempty"`

	// SearchQuery - Represents the keywords in a customer search query.
	SearchQuery *string `json:"searchQuery,omitempty"`

	// Authenticated - Indicates whether the event was produced during an authenticated session.
	Authenticated *bool `json:"authenticated,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webevent) SetField(field string, fieldValue interface{}) {
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

func (o Webevent) MarshalJSON() ([]byte, error) {
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
	type Alias Webevent
	
	return json.Marshal(&struct { 
		EventName *string `json:"eventName,omitempty"`
		
		TotalEventCount *int `json:"totalEventCount,omitempty"`
		
		TotalPageviewCount *int `json:"totalPageviewCount,omitempty"`
		
		Page *Journeypage `json:"page,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Browser *Browser `json:"browser,omitempty"`
		
		Device *Device `json:"device,omitempty"`
		
		Geolocation *Journeygeolocation `json:"geolocation,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`
		
		Referrer *Referrer `json:"referrer,omitempty"`
		
		Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`
		
		Traits *map[string]Customeventattribute `json:"traits,omitempty"`
		
		SearchQuery *string `json:"searchQuery,omitempty"`
		
		Authenticated *bool `json:"authenticated,omitempty"`
		Alias
	}{ 
		EventName: o.EventName,
		
		TotalEventCount: o.TotalEventCount,
		
		TotalPageviewCount: o.TotalPageviewCount,
		
		Page: o.Page,
		
		UserAgentString: o.UserAgentString,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		MktCampaign: o.MktCampaign,
		
		Referrer: o.Referrer,
		
		Attributes: o.Attributes,
		
		Traits: o.Traits,
		
		SearchQuery: o.SearchQuery,
		
		Authenticated: o.Authenticated,
		Alias:    (Alias)(o),
	})
}

func (o *Webevent) UnmarshalJSON(b []byte) error {
	var WebeventMap map[string]interface{}
	err := json.Unmarshal(b, &WebeventMap)
	if err != nil {
		return err
	}
	
	if EventName, ok := WebeventMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if TotalEventCount, ok := WebeventMap["totalEventCount"].(float64); ok {
		TotalEventCountInt := int(TotalEventCount)
		o.TotalEventCount = &TotalEventCountInt
	}
	
	if TotalPageviewCount, ok := WebeventMap["totalPageviewCount"].(float64); ok {
		TotalPageviewCountInt := int(TotalPageviewCount)
		o.TotalPageviewCount = &TotalPageviewCountInt
	}
	
	if Page, ok := WebeventMap["page"].(map[string]interface{}); ok {
		PageString, _ := json.Marshal(Page)
		json.Unmarshal(PageString, &o.Page)
	}
	
	if UserAgentString, ok := WebeventMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Browser, ok := WebeventMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := WebeventMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := WebeventMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if IpAddress, ok := WebeventMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := WebeventMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if MktCampaign, ok := WebeventMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if Referrer, ok := WebeventMap["referrer"].(map[string]interface{}); ok {
		ReferrerString, _ := json.Marshal(Referrer)
		json.Unmarshal(ReferrerString, &o.Referrer)
	}
	
	if Attributes, ok := WebeventMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Traits, ok := WebeventMap["traits"].(map[string]interface{}); ok {
		TraitsString, _ := json.Marshal(Traits)
		json.Unmarshal(TraitsString, &o.Traits)
	}
	
	if SearchQuery, ok := WebeventMap["searchQuery"].(string); ok {
		o.SearchQuery = &SearchQuery
	}
    
	if Authenticated, ok := WebeventMap["authenticated"].(bool); ok {
		o.Authenticated = &Authenticated
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
