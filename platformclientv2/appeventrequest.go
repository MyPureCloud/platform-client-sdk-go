package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Appeventrequest
type Appeventrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventName - Represents the action the customer performed. Event types are created for each unique event name and can be faceted on in segment and outcome conditions. A valid event name must only contain alphanumeric characters and underscores. A good event name is typically an object followed by the action performed in past tense, e.g. screen_viewed, search_performed, user_registered.
	EventName *string `json:"eventName,omitempty"`

	// ScreenName - The name of the screen, view, or fragment in the app where the event took place.
	ScreenName *string `json:"screenName,omitempty"`

	// App - Application that the customer is interacting with.
	App *Journeyapp `json:"app,omitempty"`

	// Device - Customer's device.
	Device *Requestdevice `json:"device,omitempty"`

	// SdkLibrary - SDK library used to generate the event.
	SdkLibrary *Sdklibrary `json:"sdkLibrary,omitempty"`

	// NetworkConnectivity - Information relating to the device's network connectivity.
	NetworkConnectivity *Networkconnectivity `json:"networkConnectivity,omitempty"`

	// ReferrerUrl - The referrer URL of the first event in the app session.
	ReferrerUrl *string `json:"referrerUrl,omitempty"`

	// SearchQuery - Represents the keywords in a customer search query.
	SearchQuery *string `json:"searchQuery,omitempty"`

	// Attributes - User-defined attributes associated with a particular event. These attributes provide additional context about the event. For example, items_in_cart or subscription_level.
	Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`

	// Traits - Traits are attributes intrinsic to the customer that may be sent in selected events, (e.g. email, givenName, cellPhone). Traits are used to collect information for identity resolution. For example, the same person might be using an application on different devices which might create two sessions with different customerIds. Additional information can be provided as traits to help link those two sessions and customers to a single external contact through common identifiers that were submitted via a form fill, message, or other input in both sessions.
	Traits *map[string]Customeventattribute `json:"traits,omitempty"`

	// CustomerCookieId - A UUID representing the customer associated with the app event. This is expected to be set per application install or device and can be used to identify a single customer across multiple sessions. This identifier, along with others passed as traits, is used for identity resolution.
	CustomerCookieId *string `json:"customerCookieId,omitempty"`

	// CreatedDate - UTC timestamp indicating when the event actually took place, events older than an hour will be rejected. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Appeventrequest) SetField(field string, fieldValue interface{}) {
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

func (o Appeventrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Appeventrequest
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		EventName *string `json:"eventName,omitempty"`
		
		ScreenName *string `json:"screenName,omitempty"`
		
		App *Journeyapp `json:"app,omitempty"`
		
		Device *Requestdevice `json:"device,omitempty"`
		
		SdkLibrary *Sdklibrary `json:"sdkLibrary,omitempty"`
		
		NetworkConnectivity *Networkconnectivity `json:"networkConnectivity,omitempty"`
		
		ReferrerUrl *string `json:"referrerUrl,omitempty"`
		
		SearchQuery *string `json:"searchQuery,omitempty"`
		
		Attributes *map[string]Customeventattribute `json:"attributes,omitempty"`
		
		Traits *map[string]Customeventattribute `json:"traits,omitempty"`
		
		CustomerCookieId *string `json:"customerCookieId,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		Alias
	}{ 
		EventName: o.EventName,
		
		ScreenName: o.ScreenName,
		
		App: o.App,
		
		Device: o.Device,
		
		SdkLibrary: o.SdkLibrary,
		
		NetworkConnectivity: o.NetworkConnectivity,
		
		ReferrerUrl: o.ReferrerUrl,
		
		SearchQuery: o.SearchQuery,
		
		Attributes: o.Attributes,
		
		Traits: o.Traits,
		
		CustomerCookieId: o.CustomerCookieId,
		
		CreatedDate: CreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Appeventrequest) UnmarshalJSON(b []byte) error {
	var AppeventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AppeventrequestMap)
	if err != nil {
		return err
	}
	
	if EventName, ok := AppeventrequestMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if ScreenName, ok := AppeventrequestMap["screenName"].(string); ok {
		o.ScreenName = &ScreenName
	}
    
	if App, ok := AppeventrequestMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	
	if Device, ok := AppeventrequestMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if SdkLibrary, ok := AppeventrequestMap["sdkLibrary"].(map[string]interface{}); ok {
		SdkLibraryString, _ := json.Marshal(SdkLibrary)
		json.Unmarshal(SdkLibraryString, &o.SdkLibrary)
	}
	
	if NetworkConnectivity, ok := AppeventrequestMap["networkConnectivity"].(map[string]interface{}); ok {
		NetworkConnectivityString, _ := json.Marshal(NetworkConnectivity)
		json.Unmarshal(NetworkConnectivityString, &o.NetworkConnectivity)
	}
	
	if ReferrerUrl, ok := AppeventrequestMap["referrerUrl"].(string); ok {
		o.ReferrerUrl = &ReferrerUrl
	}
    
	if SearchQuery, ok := AppeventrequestMap["searchQuery"].(string); ok {
		o.SearchQuery = &SearchQuery
	}
    
	if Attributes, ok := AppeventrequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Traits, ok := AppeventrequestMap["traits"].(map[string]interface{}); ok {
		TraitsString, _ := json.Marshal(Traits)
		json.Unmarshal(TraitsString, &o.Traits)
	}
	
	if CustomerCookieId, ok := AppeventrequestMap["customerCookieId"].(string); ok {
		o.CustomerCookieId = &CustomerCookieId
	}
    
	if createdDateString, ok := AppeventrequestMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Appeventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
