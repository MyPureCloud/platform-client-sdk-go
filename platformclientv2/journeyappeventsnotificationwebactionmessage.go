package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyappeventsnotificationwebactionmessage
type Journeyappeventsnotificationwebactionmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action
	Action *Journeyappeventsnotificationeventaction `json:"action,omitempty"`

	// ActionTarget
	ActionTarget *Journeyappeventsnotificationactiontarget `json:"actionTarget,omitempty"`

	// ActionMap
	ActionMap *Journeyappeventsnotificationactionmap `json:"actionMap,omitempty"`

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
	Browser *Journeyappeventsnotificationbrowser `json:"browser,omitempty"`

	// Device
	Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`

	// Geolocation
	Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`

	// MktCampaign
	MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`

	// VisitReferrer
	VisitReferrer *Journeyappeventsnotificationreferrer `json:"visitReferrer,omitempty"`

	// TimeToDisposition
	TimeToDisposition *int `json:"timeToDisposition,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyappeventsnotificationwebactionmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeyappeventsnotificationwebactionmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeyappeventsnotificationwebactionmessage
	
	return json.Marshal(&struct { 
		Action *Journeyappeventsnotificationeventaction `json:"action,omitempty"`
		
		ActionTarget *Journeyappeventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		ActionMap *Journeyappeventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		Browser *Journeyappeventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeyappeventsnotificationreferrer `json:"visitReferrer,omitempty"`
		
		TimeToDisposition *int `json:"timeToDisposition,omitempty"`
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
		
		TimeToDisposition: o.TimeToDisposition,
		Alias:    (Alias)(o),
	})
}

func (o *Journeyappeventsnotificationwebactionmessage) UnmarshalJSON(b []byte) error {
	var JourneyappeventsnotificationwebactionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappeventsnotificationwebactionmessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneyappeventsnotificationwebactionmessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionTarget, ok := JourneyappeventsnotificationwebactionmessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if ActionMap, ok := JourneyappeventsnotificationwebactionmessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ErrorCode, ok := JourneyappeventsnotificationwebactionmessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := JourneyappeventsnotificationwebactionmessageMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if UserAgentString, ok := JourneyappeventsnotificationwebactionmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if IpAddress, ok := JourneyappeventsnotificationwebactionmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneyappeventsnotificationwebactionmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if Browser, ok := JourneyappeventsnotificationwebactionmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneyappeventsnotificationwebactionmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneyappeventsnotificationwebactionmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneyappeventsnotificationwebactionmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneyappeventsnotificationwebactionmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	
	if TimeToDisposition, ok := JourneyappeventsnotificationwebactionmessageMap["timeToDisposition"].(float64); ok {
		TimeToDispositionInt := int(TimeToDisposition)
		o.TimeToDisposition = &TimeToDispositionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyappeventsnotificationwebactionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
