package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsphonenumberprovision
type Smsphonenumberprovision struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// PhoneNumber - A phone number to be used for SMS communications. E.g. +13175555555 or +34234234234
	PhoneNumber *string `json:"phoneNumber,omitempty"`

	// PhoneNumberType - Type of the phone number provisioned.
	PhoneNumberType *string `json:"phoneNumberType,omitempty"`

	// CountryCode - The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode *string `json:"countryCode,omitempty"`

	// AddressId - The id of an address added on your account. Due to regulatory requirements in some countries, an address may be required when provisioning a sms number. In those cases you should provide the provisioned sms address id here
	AddressId *string `json:"addressId,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Smsphonenumberprovision) SetField(field string, fieldValue interface{}) {
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

func (o Smsphonenumberprovision) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Smsphonenumberprovision
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		PhoneNumberType *string `json:"phoneNumberType,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		AddressId *string `json:"addressId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		PhoneNumber: o.PhoneNumber,
		
		PhoneNumberType: o.PhoneNumberType,
		
		CountryCode: o.CountryCode,
		
		AddressId: o.AddressId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Smsphonenumberprovision) UnmarshalJSON(b []byte) error {
	var SmsphonenumberprovisionMap map[string]interface{}
	err := json.Unmarshal(b, &SmsphonenumberprovisionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsphonenumberprovisionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SmsphonenumberprovisionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if PhoneNumber, ok := SmsphonenumberprovisionMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if PhoneNumberType, ok := SmsphonenumberprovisionMap["phoneNumberType"].(string); ok {
		o.PhoneNumberType = &PhoneNumberType
	}
    
	if CountryCode, ok := SmsphonenumberprovisionMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if AddressId, ok := SmsphonenumberprovisionMap["addressId"].(string); ok {
		o.AddressId = &AddressId
	}
    
	if SelfUri, ok := SmsphonenumberprovisionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Smsphonenumberprovision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
