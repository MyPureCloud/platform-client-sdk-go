package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopiccontactaddress
type Externalcontactsunresolvedcontactchangedtopiccontactaddress struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Address1
	Address1 *string `json:"address1,omitempty"`

	// Address2
	Address2 *string `json:"address2,omitempty"`

	// City
	City *string `json:"city,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// PostalCode
	PostalCode *string `json:"postalCode,omitempty"`

	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactsunresolvedcontactchangedtopiccontactaddress) MarshalJSON() ([]byte, error) {
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
	type Alias Externalcontactsunresolvedcontactchangedtopiccontactaddress
	
	return json.Marshal(&struct { 
		Address1 *string `json:"address1,omitempty"`
		
		Address2 *string `json:"address2,omitempty"`
		
		City *string `json:"city,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		PostalCode *string `json:"postalCode,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		Alias
	}{ 
		Address1: o.Address1,
		
		Address2: o.Address2,
		
		City: o.City,
		
		State: o.State,
		
		PostalCode: o.PostalCode,
		
		CountryCode: o.CountryCode,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap)
	if err != nil {
		return err
	}
	
	if Address1, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["address1"].(string); ok {
		o.Address1 = &Address1
	}
    
	if Address2, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["address2"].(string); ok {
		o.Address2 = &Address2
	}
    
	if City, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["city"].(string); ok {
		o.City = &City
	}
    
	if State, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["state"].(string); ok {
		o.State = &State
	}
    
	if PostalCode, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["postalCode"].(string); ok {
		o.PostalCode = &PostalCode
	}
    
	if CountryCode, ok := ExternalcontactsunresolvedcontactchangedtopiccontactaddressMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopiccontactaddress) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
