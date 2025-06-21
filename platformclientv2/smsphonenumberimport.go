package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsphonenumberimport
type Smsphonenumberimport struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// PhoneNumber - A phone number to be used for SMS communications. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// PhoneNumberType - Type of the phone number provisioned.
	PhoneNumberType *string `json:"phoneNumberType,omitempty"`

	// CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode *string `json:"countryCode,omitempty"`

	// IntegrationId - The id of the Genesys Cloud integration this phone number belongs to.
	IntegrationId *string `json:"integrationId,omitempty"`

	// Compliance - Compliance configuration for short codes, including help, stop and opt in.
	Compliance *Compliance `json:"compliance,omitempty"`

	// SupportedContent - Defines the media SupportedContent profile configured for an MMS capable phone number.
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Smsphonenumberimport) SetField(field string, fieldValue interface{}) {
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

func (o Smsphonenumberimport) MarshalJSON() ([]byte, error) {
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
	type Alias Smsphonenumberimport
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		PhoneNumberType *string `json:"phoneNumberType,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		Compliance *Compliance `json:"compliance,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		PhoneNumber: o.PhoneNumber,
		
		PhoneNumberType: o.PhoneNumberType,
		
		CountryCode: o.CountryCode,
		
		IntegrationId: o.IntegrationId,
		
		Compliance: o.Compliance,
		
		SupportedContent: o.SupportedContent,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Smsphonenumberimport) UnmarshalJSON(b []byte) error {
	var SmsphonenumberimportMap map[string]interface{}
	err := json.Unmarshal(b, &SmsphonenumberimportMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsphonenumberimportMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if PhoneNumber, ok := SmsphonenumberimportMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if PhoneNumberType, ok := SmsphonenumberimportMap["phoneNumberType"].(string); ok {
		o.PhoneNumberType = &PhoneNumberType
	}
    
	if CountryCode, ok := SmsphonenumberimportMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if IntegrationId, ok := SmsphonenumberimportMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if Compliance, ok := SmsphonenumberimportMap["compliance"].(map[string]interface{}); ok {
		ComplianceString, _ := json.Marshal(Compliance)
		json.Unmarshal(ComplianceString, &o.Compliance)
	}
	
	if SupportedContent, ok := SmsphonenumberimportMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if SelfUri, ok := SmsphonenumberimportMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Smsphonenumberimport) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
