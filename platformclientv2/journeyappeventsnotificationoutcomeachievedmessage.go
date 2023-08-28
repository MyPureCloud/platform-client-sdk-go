package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyappeventsnotificationoutcomeachievedmessage
type Journeyappeventsnotificationoutcomeachievedmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Outcome
	Outcome *Journeyappeventsnotificationoutcome `json:"outcome,omitempty"`

	// Browser
	Browser *Journeyappeventsnotificationbrowser `json:"browser,omitempty"`

	// VisitCreatedDate
	VisitCreatedDate *time.Time `json:"visitCreatedDate,omitempty"`

	// IpAddress
	IpAddress *string `json:"ipAddress,omitempty"`

	// IpOrganization
	IpOrganization *string `json:"ipOrganization,omitempty"`

	// UserAgentString
	UserAgentString *string `json:"userAgentString,omitempty"`

	// Device
	Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`

	// Geolocation
	Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`

	// MktCampaign
	MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`

	// VisitReferrer
	VisitReferrer *Journeyappeventsnotificationreferrer `json:"visitReferrer,omitempty"`

	// AssociatedValue
	AssociatedValue *Journeyappeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Journeyappeventsnotificationoutcomeachievedmessage) SetField(field string, fieldValue interface{}) {
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

func (o Journeyappeventsnotificationoutcomeachievedmessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "VisitCreatedDate", }
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
	type Alias Journeyappeventsnotificationoutcomeachievedmessage
	
	VisitCreatedDate := new(string)
	if o.VisitCreatedDate != nil {
		
		*VisitCreatedDate = timeutil.Strftime(o.VisitCreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VisitCreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Journeyappeventsnotificationoutcome `json:"outcome,omitempty"`
		
		Browser *Journeyappeventsnotificationbrowser `json:"browser,omitempty"`
		
		VisitCreatedDate *string `json:"visitCreatedDate,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Device *Journeyappeventsnotificationdevice `json:"device,omitempty"`
		
		Geolocation *Journeyappeventsnotificationgeolocation `json:"geolocation,omitempty"`
		
		MktCampaign *Journeyappeventsnotificationmktcampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Journeyappeventsnotificationreferrer `json:"visitReferrer,omitempty"`
		
		AssociatedValue *Journeyappeventsnotificationassociatedvalue `json:"associatedValue,omitempty"`
		Alias
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
		Alias:    (Alias)(o),
	})
}

func (o *Journeyappeventsnotificationoutcomeachievedmessage) UnmarshalJSON(b []byte) error {
	var JourneyappeventsnotificationoutcomeachievedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyappeventsnotificationoutcomeachievedmessageMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if Browser, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if visitCreatedDateString, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["visitCreatedDate"].(string); ok {
		VisitCreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", visitCreatedDateString)
		o.VisitCreatedDate = &VisitCreatedDate
	}
	
	if IpAddress, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if UserAgentString, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Device, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if MktCampaign, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	
	if AssociatedValue, ok := JourneyappeventsnotificationoutcomeachievedmessageMap["associatedValue"].(map[string]interface{}); ok {
		AssociatedValueString, _ := json.Marshal(AssociatedValue)
		json.Unmarshal(AssociatedValueString, &o.AssociatedValue)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyappeventsnotificationoutcomeachievedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
