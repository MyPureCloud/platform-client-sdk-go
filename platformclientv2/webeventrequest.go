package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webeventrequest
type Webeventrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CustomerCookieId - A UUID representing the customer making the request.
	CustomerCookieId *string `json:"customerCookieId,omitempty"`

	// EventName - Represents the action the customer performed. Event types are created for each unique event name and can be faceted on in segment and outcome conditions. A valid event name must only contain alphanumeric characters and underscores. A good event name is typically an object followed by the action performed in past tense, e.g. page_viewed, order_completed, user_registered.
	EventName *string `json:"eventName,omitempty"`

	// Page - The webpage where the user interaction occurred.
	Page *Requestpage `json:"page,omitempty"`

	// UserAgentString - Override for HTTP User-Agent string from request header (see https://tools.ietf.org/html/rfc1945#section-10.15).
	UserAgentString *string `json:"userAgentString,omitempty"`

	// Browser - Customer's browser.
	Browser *Webeventbrowser `json:"browser,omitempty"`

	// Device - Customer's device.
	Device *Webeventdevice `json:"device,omitempty"`

	// SearchQuery - Represents the keywords in a customer search query.
	SearchQuery *string `json:"searchQuery,omitempty"`

	// IpAddress - Customer's IP address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// ReferrerUrl - Identifies the referrer URL that originally generated the request for the current page being viewed.
	ReferrerUrl *string `json:"referrerUrl,omitempty"`

	// Attributes - User-defined attributes associated with a particular event. These attributes provide additional context about the event. For example, items_in_cart or subscription_level.
	Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`

	// Traits - Traits are attributes intrinsic to the customer that may be sent in selected events, e.g. email, lastName, cellPhone. Traits are used to collect information for identity resolution. For example, the same person might be using an application on different devices which might create two sessions with different customerIds. Additional information can be provided as traits to help link those two sessions and customers to a single external contact through common identifiers that were submitted via a form fill, message, or other input in both sessions.
	Traits *map[string]Customeventattribute `json:"traits,omitempty"`

	// CreatedDate - UTC timestamp indicating when the event actually took place, events older than an hour will be rejected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webeventrequest) SetField(field string, fieldValue interface{}) {
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

func (o Webeventrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Webeventrequest
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		CustomerCookieId *string `json:"customerCookieId,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		Page *Requestpage `json:"page,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Browser *Webeventbrowser `json:"browser,omitempty"`
		
		Device *Webeventdevice `json:"device,omitempty"`
		
		SearchQuery *string `json:"searchQuery,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		ReferrerUrl *string `json:"referrerUrl,omitempty"`
		
		Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`
		
		Traits *map[string]Customeventattribute `json:"traits,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		CustomerCookieId: o.CustomerCookieId,
		
		EventName: o.EventName,
		
		Page: o.Page,
		
		UserAgentString: o.UserAgentString,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		SearchQuery: o.SearchQuery,
		
		IpAddress: o.IpAddress,
		
		ReferrerUrl: o.ReferrerUrl,
		
		Attributes: o.Attributes,
		
		Traits: o.Traits,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Webeventrequest) UnmarshalJSON(b []byte) error {
	var WebeventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WebeventrequestMap)
	if err != nil {
		return err
	}
	
	if CustomerCookieId, ok := WebeventrequestMap["customerCookieId"].(string); ok {
		o.CustomerCookieId = &CustomerCookieId
	}
    
	if EventName, ok := WebeventrequestMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if Page, ok := WebeventrequestMap["page"].(map[string]interface{}); ok {
		PageString, _ := json.Marshal(Page)
		json.Unmarshal(PageString, &o.Page)
	}
	
	if UserAgentString, ok := WebeventrequestMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Browser, ok := WebeventrequestMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := WebeventrequestMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if SearchQuery, ok := WebeventrequestMap["searchQuery"].(string); ok {
		o.SearchQuery = &SearchQuery
	}
    
	if IpAddress, ok := WebeventrequestMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if ReferrerUrl, ok := WebeventrequestMap["referrerUrl"].(string); ok {
		o.ReferrerUrl = &ReferrerUrl
	}
    
	if Attributes, ok := WebeventrequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Traits, ok := WebeventrequestMap["traits"].(map[string]interface{}); ok {
		TraitsString, _ := json.Marshal(Traits)
		json.Unmarshal(TraitsString, &o.Traits)
	}
	
	if createdDateString, ok := WebeventrequestMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
