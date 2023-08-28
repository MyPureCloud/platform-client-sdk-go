package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyappeventsnotificationappmessage
type Journeyappeventsnotificationappmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventName
	EventName *string `json:"eventName,omitempty"`

	// ScreenName
	ScreenName *string `json:"screenName,omitempty"`

	// App
	App *Journeyappeventsnotificationapp `json:"app,omitempty"`

	// Device
	Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`

	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`

	// Geolocation
	Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`

	// SdkLibrary
	SdkLibrary *Journeyappeventsnotificationsdklibrary `json:"sdkLibrary,omitempty"`

	// NetworkConnectivity
	NetworkConnectivity *Journeyappeventsnotificationnetworkconnectivity `json:"networkConnectivity,omitempty"`

	// MktCampaign
	MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`

	// SearchQuery
	SearchQuery *string `json:"searchQuery,omitempty"`

	// Attributes
	Attributes *map[string]Journeyappeventsnotificationcustomeventattribute `json:"attributes,omitempty"`

	// Traits
	Traits *map[string]Journeyappeventsnotificationcustomeventattribute `json:"traits,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyappeventsnotificationappmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeyappeventsnotificationappmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyappeventsnotificationappmessage
	
	return json.Marshal(&struct { 
		EventName *string `json:"eventName,omitempty"`
		
		ScreenName *string `json:"screenName,omitempty"`
		
		App *Journeyappeventsnotificationapp `json:"app,omitempty"`
		
		Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		SdkLibrary *Journeyappeventsnotificationsdklibrary `json:"sdkLibrary,omitempty"`
		
		NetworkConnectivity *Journeyappeventsnotificationnetworkconnectivity `json:"networkConnectivity,omitempty"`
		
		MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		SearchQuery *string `json:"searchQuery,omitempty"`
		
		Attributes *map[string]Journeyappeventsnotificationcustomeventattribute `json:"attributes,omitempty"`
		
		Traits *map[string]Journeyappeventsnotificationcustomeventattribute `json:"traits,omitempty"`
		Alias
	}{ 
		EventName: o.EventName,
		
		ScreenName: o.ScreenName,
		
		App: o.App,
		
		Device: o.Device,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		Geolocation: o.Geolocation,
		
		SdkLibrary: o.SdkLibrary,
		
		NetworkConnectivity: o.NetworkConnectivity,
		
		MktCampaign: o.MktCampaign,
		
		SearchQuery: o.SearchQuery,
		
		Attributes: o.Attributes,
		
		Traits: o.Traits,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyappeventsnotificationappmessage) UnmarshalJSON(b []byte) error {
	var JourneyappeventsnotificationappmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappeventsnotificationappmessageMap)
	if err != nil {
		return err
	}
	
	if EventName, ok := JourneyappeventsnotificationappmessageMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if ScreenName, ok := JourneyappeventsnotificationappmessageMap["screenName"].(string); ok {
		o.ScreenName = &ScreenName
	}
    
	if App, ok := JourneyappeventsnotificationappmessageMap["app"].(map[string]interface{}); ok {
		AppString, _ := json.Marshal(App)
		json.Unmarshal(AppString, &o.App)
	}
	
	if Device, ok := JourneyappeventsnotificationappmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if IpAddress, ok := JourneyappeventsnotificationappmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneyappeventsnotificationappmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if Geolocation, ok := JourneyappeventsnotificationappmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if SdkLibrary, ok := JourneyappeventsnotificationappmessageMap["sdkLibrary"].(map[string]interface{}); ok {
		SdkLibraryString, _ := json.Marshal(SdkLibrary)
		json.Unmarshal(SdkLibraryString, &o.SdkLibrary)
	}
	
	if NetworkConnectivity, ok := JourneyappeventsnotificationappmessageMap["networkConnectivity"].(map[string]interface{}); ok {
		NetworkConnectivityString, _ := json.Marshal(NetworkConnectivity)
		json.Unmarshal(NetworkConnectivityString, &o.NetworkConnectivity)
	}
	
	if MktCampaign, ok := JourneyappeventsnotificationappmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if SearchQuery, ok := JourneyappeventsnotificationappmessageMap["searchQuery"].(string); ok {
		o.SearchQuery = &SearchQuery
	}
    
	if Attributes, ok := JourneyappeventsnotificationappmessageMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Traits, ok := JourneyappeventsnotificationappmessageMap["traits"].(map[string]interface{}); ok {
		TraitsString, _ := json.Marshal(Traits)
		json.Unmarshal(TraitsString, &o.Traits)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyappeventsnotificationappmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
