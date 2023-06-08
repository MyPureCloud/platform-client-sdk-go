package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Outcomeachievedevent
type Outcomeachievedevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Outcome - The outcome achieved.
	Outcome *Outcomeachievedeventoutcome `json:"outcome,omitempty"`

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

	// VisitCreatedDate - When visit was created (e.g. timestamp of the first event in visit). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	VisitCreatedDate *time.Time `json:"visitCreatedDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Outcomeachievedevent) SetField(field string, fieldValue interface{}) {
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

func (o Outcomeachievedevent) MarshalJSON() ([]byte, error) {
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
	type Alias Outcomeachievedevent
	
	VisitCreatedDate := new(string)
	if o.VisitCreatedDate != nil {
		
		*VisitCreatedDate = timeutil.Strftime(o.VisitCreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VisitCreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		Outcome *Outcomeachievedeventoutcome `json:"outcome,omitempty"`
		
		UserAgentString *string `json:"userAgentString,omitempty"`
		
		Browser *Browser `json:"browser,omitempty"`
		
		Device *Device `json:"device,omitempty"`
		
		Geolocation *Journeygeolocation `json:"geolocation,omitempty"`
		
		IpAddress *string `json:"ipAddress,omitempty"`
		
		IpOrganization *string `json:"ipOrganization,omitempty"`
		
		MktCampaign *Journeycampaign `json:"mktCampaign,omitempty"`
		
		VisitReferrer *Referrer `json:"visitReferrer,omitempty"`
		
		VisitCreatedDate *string `json:"visitCreatedDate,omitempty"`
		Alias
	}{ 
		Outcome: o.Outcome,
		
		UserAgentString: o.UserAgentString,
		
		Browser: o.Browser,
		
		Device: o.Device,
		
		Geolocation: o.Geolocation,
		
		IpAddress: o.IpAddress,
		
		IpOrganization: o.IpOrganization,
		
		MktCampaign: o.MktCampaign,
		
		VisitReferrer: o.VisitReferrer,
		
		VisitCreatedDate: VisitCreatedDate,
		Alias:    (Alias)(o),
	})
}

func (o *Outcomeachievedevent) UnmarshalJSON(b []byte) error {
	var OutcomeachievedeventMap map[string]interface{}
	err := json.Unmarshal(b, &OutcomeachievedeventMap)
	if err != nil {
		return err
	}
	
	if Outcome, ok := OutcomeachievedeventMap["outcome"].(map[string]interface{}); ok {
		OutcomeString, _ := json.Marshal(Outcome)
		json.Unmarshal(OutcomeString, &o.Outcome)
	}
	
	if UserAgentString, ok := OutcomeachievedeventMap["userAgentString"].(string); ok {
		o.UserAgentString = &UserAgentString
	}
    
	if Browser, ok := OutcomeachievedeventMap["browser"].(map[string]interface{}); ok {
		BrowserString, _ := json.Marshal(Browser)
		json.Unmarshal(BrowserString, &o.Browser)
	}
	
	if Device, ok := OutcomeachievedeventMap["device"].(map[string]interface{}); ok {
		DeviceString, _ := json.Marshal(Device)
		json.Unmarshal(DeviceString, &o.Device)
	}
	
	if Geolocation, ok := OutcomeachievedeventMap["geolocation"].(map[string]interface{}); ok {
		GeolocationString, _ := json.Marshal(Geolocation)
		json.Unmarshal(GeolocationString, &o.Geolocation)
	}
	
	if IpAddress, ok := OutcomeachievedeventMap["ipAddress"].(string); ok {
		o.IpAddress = &IpAddress
	}
    
	if IpOrganization, ok := OutcomeachievedeventMap["ipOrganization"].(string); ok {
		o.IpOrganization = &IpOrganization
	}
    
	if MktCampaign, ok := OutcomeachievedeventMap["mktCampaign"].(map[string]interface{}); ok {
		MktCampaignString, _ := json.Marshal(MktCampaign)
		json.Unmarshal(MktCampaignString, &o.MktCampaign)
	}
	
	if VisitReferrer, ok := OutcomeachievedeventMap["visitReferrer"].(map[string]interface{}); ok {
		VisitReferrerString, _ := json.Marshal(VisitReferrer)
		json.Unmarshal(VisitReferrerString, &o.VisitReferrer)
	}
	
	if visitCreatedDateString, ok := OutcomeachievedeventMap["visitCreatedDate"].(string); ok {
		VisitCreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", visitCreatedDateString)
		o.VisitCreatedDate = &VisitCreatedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outcomeachievedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
