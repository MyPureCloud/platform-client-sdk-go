package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Smsaddressprovision
type Smsaddressprovision struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - Name associated with this address
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

	// AutoCorrectAddress - This is used when the address is created. If the value is not set or true, then the system will, if necessary, auto-correct the address you provide. Set this value to false if the system should not auto-correct the address.
	AutoCorrectAddress *bool `json:"autoCorrectAddress,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Smsaddressprovision) SetField(field string, fieldValue interface{}) {
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

func (o Smsaddressprovision) MarshalJSON() ([]byte, error) {
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
	type Alias Smsaddressprovision
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Street *string `json:"street,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		AutoCorrectAddress *bool `json:"autoCorrectAddress,omitempty"`
		
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
		
		AutoCorrectAddress: o.AutoCorrectAddress,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Smsaddressprovision) UnmarshalJSON(b []byte) error {
	var SmsaddressprovisionMap map[string]interface{}
	err := json.Unmarshal(b, &SmsaddressprovisionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SmsaddressprovisionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SmsaddressprovisionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Street, ok := SmsaddressprovisionMap["street"].(string); ok {
		o.Street = &Street
	}
    
	if City, ok := SmsaddressprovisionMap["city"].(string); ok {
		o.City = &City
	}
    
	if Region, ok := SmsaddressprovisionMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if PostalCode, ok := SmsaddressprovisionMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if CountryCode, ok := SmsaddressprovisionMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if AutoCorrectAddress, ok := SmsaddressprovisionMap["autoCorrectAddress"].(bool); ok {
		o.AutoCorrectAddress = &AutoCorrectAddress
	}
    
	if SelfUri, ok := SmsaddressprovisionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Smsaddressprovision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
