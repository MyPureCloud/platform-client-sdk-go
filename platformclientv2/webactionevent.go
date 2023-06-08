package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webactionevent
type Webactionevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Action - The action that triggered the event.
	Action *Eventaction `json:"action,omitempty"`

	// ActionMap - The action map that triggered the action.
	ActionMap *Actioneventactionmap `json:"actionMap,omitempty"`

	// ActionTarget - The target for engagement actions.
	ActionTarget *Addressableentityref `json:"actionTarget,omitempty"`

	// TimeToDisposition - Milliseconds elapsed until the action is disposed.
	TimeToDisposition *int `json:"timeToDisposition,omitempty"`

	// ErrorCode - Code of the error returned when the action fails.
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessage - Message of the error returned when the action fails.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// UserAgentString - HTTP User-Agent string (see https://tools.ietf.org/html/rfc1945#section-10.15).
	UserAgentString *string `json:"userAgentString,omitempty"`

	// Browser - Customer's browser.
	Browser *Browser `json:"browser,omitempty"`

	// Device - Customer's device.
	Device *Device `json:"device,omitempty"`

	// Geolocation - Customer's geolocation.
	Geolocation *Journeygeolocation `json:"geolocation,omitempty"`

	// IpAddress - Visitor's IP address.
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpOrganization - Visitor's IP-based organization or ISP name.
	IpOrganization *string `json:"ipOrganization,omitempty"`

	// MktCampaign - Marketing / traffic source information.
	MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`

	// VisitReferrer - Visit's referrer.
	VisitReferrer *Referrer `json:"visitReferrer,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webactionevent) SetField(field string, fieldValue interface{}) {
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

func (o Webactionevent) MarshalJSON() ([]byte, error) {
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
	type Alias Webactionevent
	
	return json.Marshal(&struct { 
		Action *Eventaction `json:"action,omitempty"`
		
		ActionMap *Actioneventactionmap `json:"actionMap,omitempty"`
		
		ActionTarget *Addressableentityref `json:"actionTarget,omitempty"`
		
		TimeToDisposition *int `json:"timeToDisposition,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Browser *Browser `json:"browser,omitempty"`
		
		Device *Device `json:"device,omitempty"`
		
		Geolocation *Journeygeolocation `json:"geolocation,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Referrer `json:"visitReferrer,omitempty"`
		Alias
	}{ 
		Action: o.Action,
		
		ActionMap: o.ActionMap,
		
		ActionTarget: o.ActionTarget,
		
		TimeToDisposition: o.TimeToDisposition,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		
		UserAgentString: o.UserAgentString,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		MktCampaign: o.MktCampaign,
		
		VisitReferrer: o.VisitReferrer,
		Alias:    (Alias)(o),
	})
}

func (o *Webactionevent) UnmarshalJSON(b []byte) error {
	var WebactioneventMap map[string]interface{}
	err := json.Unmarshal(b, &WebactioneventMap)
	if err != nil {
		return err
	}
	
	if Action, ok := WebactioneventMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionMap, ok := WebactioneventMap["actionMap"].(map[string]interface{}); ok {
		ActionMapString, _ := json.Marshal(ActionMap)
		json.Unmarshal(ActionMapString, &o.ActionMap)
	}
	
	if ActionTarget, ok := WebactioneventMap["actionTarget"].(map[string]interface{}); ok {
		ActionTargetString, _ := json.Marshal(ActionTarget)
		json.Unmarshal(ActionTargetString, &o.ActionTarget)
	}
	
	if TimeToDisposition, ok := WebactioneventMap["timeToDisposition"].(float64); ok {
		TimeToDispositionInt := int(TimeToDisposition)
		o.TimeToDisposition = &TimeToDispositionInt
	}
	
	if ErrorCode, ok := WebactioneventMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := WebactioneventMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if UserAgentString, ok := WebactioneventMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Browser, ok := WebactioneventMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := WebactioneventMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := WebactioneventMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if IpAddress, ok := WebactioneventMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := WebactioneventMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if MktCampaign, ok := WebactioneventMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := WebactioneventMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webactionevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
