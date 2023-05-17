package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationwebactionmessage
type Journeywebactioneventsnotificationwebactionmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action
	Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`

	// ActionTarget
	ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`

	// ActionMap
	ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`

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
	Browser *Journeywebactioneventsnotificationbrowser `json:"browser,omitempty"`

	// Device
	Device *Journeywebactioneventsnotificationdevice `json:"device,omitempty"`

	// Geolocation
	Geolocation *Journeywebactioneventsnotificationgeolocation `json:"geolocation,omitempty"`

	// MktCampaign
	MktCampaign *Journeywebactioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`

	// VisitReferrer
	VisitReferrer *Journeywebactioneventsnotificationreferrer `json:"visitReferrer,omitempty"`

	// TimeToDisposition
	TimeToDisposition *int `json:"timeToDisposition,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeywebactioneventsnotificationwebactionmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeywebactioneventsnotificationwebactionmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Journeywebactioneventsnotificationwebactionmessage
	
	return json.Marshal(&struct { 
		Action *Journeywebactioneventsnotificationeventaction `json:"action,omitempty"`
		
		ActionTarget *Journeywebactioneventsnotificationactiontarget `json:"actionTarget,omitempty"`
		
		ActionMap *Journeywebactioneventsnotificationactionmap `json:"actionMap,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		Browser *Journeywebactioneventsnotificationbrowser `json:"browser,omitempty"`
		
		Device *Journeywebactioneventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeywebactioneventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeywebactioneventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeywebactioneventsnotificationreferrer `json:"visitReferrer,omitempty"`
		
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

func (o *Journeywebactioneventsnotificationwebactionmessage) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationwebactionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationwebactionmessageMap)
	if err != nil {
		return err
	}
	
	if Action, ok := JourneywebactioneventsnotificationwebactionmessageMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionTarget, ok := JourneywebactioneventsnotificationwebactionmessageMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if ActionMap, ok := JourneywebactioneventsnotificationwebactionmessageMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ErrorCode, ok := JourneywebactioneventsnotificationwebactionmessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := JourneywebactioneventsnotificationwebactionmessageMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if UserAgentString, ok := JourneywebactioneventsnotificationwebactionmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if IpAddress, ok := JourneywebactioneventsnotificationwebactionmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneywebactioneventsnotificationwebactionmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if Browser, ok := JourneywebactioneventsnotificationwebactionmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := JourneywebactioneventsnotificationwebactionmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneywebactioneventsnotificationwebactionmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneywebactioneventsnotificationwebactionmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneywebactioneventsnotificationwebactionmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	
	if TimeToDisposition, ok := JourneywebactioneventsnotificationwebactionmessageMap["timeToDisposition"].(float64); ok {
		TimeToDispositionInt := int(TimeToDisposition)
		o.TimeToDisposition = &TimeToDispositionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationwebactionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
