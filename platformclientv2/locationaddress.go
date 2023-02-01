package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Locationaddress
type Locationaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// City
	City *string `json:"city,omitempty"`

	// Country
	Country *string `json:"country,omitempty"`

	// CountryName
	CountryName *string `json:"countryName,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// Street1
	Street1 *string `json:"street1,omitempty"`

	// Street2
	Street2 *string `json:"street2,omitempty"`

	// Zipcode
	Zipcode *string `json:"zipcode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Locationaddress) SetField(field string, fieldValue interface{}) {
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

func (o Locationaddress) MarshalJSON() ([]byte, error) {
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
	type Alias Locationaddress
	
	return json.Marshal(&struct { 
		City *string `json:"city,omitempty"`
		
		Country *string `json:"country,omitempty"`
		
		CountryName *string `json:"countryName,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Street1 *string `json:"street1,omitempty"`
		
		Street2 *string `json:"street2,omitempty"`
		
		Zipcode *string `json:"zipcode,omitempty"`
		Alias
	}{ 
		City: o.City,
		
		Country: o.Country,
		
		CountryName: o.CountryName,
		
		State: o.State,
		
		Street1: o.Street1,
		
		Street2: o.Street2,
		
		Zipcode: o.Zipcode,
		Alias:    (Alias)(o),
	})
}

func (o *Locationaddress) UnmarshalJSON(b []byte) error {
	var LocationaddressMap map[string]interface{}
	err := json.Unmarshal(b, &LocationaddressMap)
	if err != nil {
		return err
	}
	
	if City, ok := LocationaddressMap["city"].(string); ok {
		o.City = &City
	}
    
	if Country, ok := LocationaddressMap["country"].(string); ok {
		o.Country = &Country
	}
    
	if CountryName, ok := LocationaddressMap["countryName"].(string); ok {
		o.CountryName = &CountryName
	}
    
	if State, ok := LocationaddressMap["state"].(string); ok {
		o.State = &State
	}
    
	if Street1, ok := LocationaddressMap["street1"].(string); ok {
		o.Street1 = &Street1
	}
    
	if Street2, ok := LocationaddressMap["street2"].(string); ok {
		o.Street2 = &Street2
	}
    
	if Zipcode, ok := LocationaddressMap["zipcode"].(string); ok {
		o.Zipcode = &Zipcode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Locationaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
