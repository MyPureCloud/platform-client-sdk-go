package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationwebactionmessage
type Journeywebeventsnotificationwebactionmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebeventsnotificationwebactionmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebeventsnotificationwebactionmessage) MarshalJSON() ([]byte, error) {
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
		Alias
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
		Alias:    (Alias)(o),
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
