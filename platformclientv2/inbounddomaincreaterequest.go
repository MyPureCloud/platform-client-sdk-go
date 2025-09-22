package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Inbounddomaincreaterequest
type Inbounddomaincreaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Unique Id of the domain such as: example.com
	Id *string `json:"id,omitempty"`

	// SubDomain - Indicates if this a PureCloud sub-domain. If true, then the appropriate DNS records are created for sending/receiving email.
	SubDomain *bool `json:"subDomain,omitempty"`

	// MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
	MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`

	// CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
	CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`

	// EmailSetting - The email settings to associate with this domain.
	EmailSetting *Emailsettingreference `json:"emailSetting,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Inbounddomaincreaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Inbounddomaincreaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Inbounddomaincreaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SubDomain *bool `json:"subDomain,omitempty"`
		
		MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`
		
		CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`
		
		EmailSetting *Emailsettingreference `json:"emailSetting,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SubDomain: o.SubDomain,
		
		MailFromSettings: o.MailFromSettings,
		
		CustomSMTPServer: o.CustomSMTPServer,
		
		EmailSetting: o.EmailSetting,
		Alias:    (Alias)(o),
	})
}

func (o *Inbounddomaincreaterequest) UnmarshalJSON(b []byte) error {
	var InbounddomaincreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &InbounddomaincreaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InbounddomaincreaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SubDomain, ok := InbounddomaincreaterequestMap["subDomain"].(bool); ok {
		o.SubDomain = &SubDomain
	}
    
	if MailFromSettings, ok := InbounddomaincreaterequestMap["mailFromSettings"].(map[string]interface{}); ok {
		MailFromSettingsString, _ := json.Marshal(MailFromSettings)
		json.Unmarshal(MailFromSettingsString, &o.MailFromSettings)
	}
	
	if CustomSMTPServer, ok := InbounddomaincreaterequestMap["customSMTPServer"].(map[string]interface{}); ok {
		CustomSMTPServerString, _ := json.Marshal(CustomSMTPServer)
		json.Unmarshal(CustomSMTPServerString, &o.CustomSMTPServer)
	}
	
	if EmailSetting, ok := InbounddomaincreaterequestMap["emailSetting"].(map[string]interface{}); ok {
		EmailSettingString, _ := json.Marshal(EmailSetting)
		json.Unmarshal(EmailSettingString, &o.EmailSetting)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Inbounddomaincreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
