package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsaddress
type Smsaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The id of this address.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Street - The number and street address where this address is located.
	Street *string `json:"street,omitempty"`

	// City - The city in which this address is in
	City *string `json:"city,omitempty"`

	// Region - The state or region this address is in
	Region *string `json:"region,omitempty"`

	// PostalCode - The postal code this address is in
	PostalCode *string `json:"postalCode,omitempty"`

	// CountryCode - The ISO country code of this address
	CountryCode *string `json:"countryCode,omitempty"`

	// Validated - In some countries, addresses are validated to comply with local regulation. In those countries, if the address you provide does not pass validation, it will not be accepted as an Address. This value will be true if the Address has been validated, or false for countries that don't require validation or if the Address is non-compliant.
	Validated *bool `json:"validated,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Smsaddress) SetField(field string, fieldValue interface{}) {
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

func (o Smsaddress) MarshalJSON() ([]byte, error) {
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
	type Alias Smsaddress
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Street *string `json:"street,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		Validated *bool `json:"validated,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Street: o.Street,
		
		City: o.City,
		
		Region: o.Region,
		
		PostalCode: o.PostalCode,
		
		CountryCode: o.CountryCode,
		
		Validated: o.Validated,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Smsaddress) UnmarshalJSON(b []byte) error {
	var SmsaddressMap map[string]interface{}
	err := json.Unmarshal(b, &SmsaddressMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsaddressMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SmsaddressMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Street, ok := SmsaddressMap["street"].(string); ok {
		o.Street = &Street
	}
    
	if City, ok := SmsaddressMap["city"].(string); ok {
		o.City = &City
	}
    
	if Region, ok := SmsaddressMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if PostalCode, ok := SmsaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if CountryCode, ok := SmsaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if Validated, ok := SmsaddressMap["validated"].(bool); ok {
		o.Validated = &Validated
	}
    
	if SelfUri, ok := SmsaddressMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Smsaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
