package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Enrichfieldrules
type Enrichfieldrules struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DefaultAction - Default behavior for combining data from the submitted request with any entity found in the database. The default behavior if unspecified is `PreferProvided`, meaning any non-null fields in the submitted request will override data in the database, but all null fields will remain unchanged. Omitting a field in the request payload means that it will be treated as null.
	DefaultAction *string `json:"defaultAction,omitempty"`

	// Rules - Field-specific behaviors for how to combine data from different sources. For example, you can set a `defaultAction` of `PreferProvided`, but use different behaviors such as `AlwaysUseProvided` or `PreferExisting` for specific fields.
	Rules *[]Enrichfieldrule `json:"rules,omitempty"`

	// DefaultArrayAction - Default behavior for combining items in array field from the submitted request with any array entity found in the database. The default behavior if unspecified is `fill`, meaning the field value will always be the partial concatenation of both the array in the Database and the array in the contact body, up to the size limit of the array
	DefaultArrayAction *string `json:"defaultArrayAction,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Enrichfieldrules) SetField(field string, fieldValue interface{}) {
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

func (o Enrichfieldrules) MarshalJSON() ([]byte, error) {
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
	type Alias Enrichfieldrules
	
	return json.Marshal(&struct { 
		DefaultAction *string `json:"defaultAction,omitempty"`
		
		Rules *[]Enrichfieldrule `json:"rules,omitempty"`
		
		DefaultArrayAction *string `json:"defaultArrayAction,omitempty"`
		Alias
	}{ 
		DefaultAction: o.DefaultAction,
		
		Rules: o.Rules,
		
		DefaultArrayAction: o.DefaultArrayAction,
		Alias:    (Alias)(o),
	})
}

func (o *Enrichfieldrules) UnmarshalJSON(b []byte) error {
	var EnrichfieldrulesMap map[string]interface{}
	err := json.Unmarshal(b, &EnrichfieldrulesMap)
	if err != nil {
		return err
	}
	
	if DefaultAction, ok := EnrichfieldrulesMap["defaultAction"].(string); ok {
		o.DefaultAction = &DefaultAction
	}
    
	if Rules, ok := EnrichfieldrulesMap["rules"].([]interface{}); ok {
		RulesString, _ := json.Marshal(Rules)
		json.Unmarshal(RulesString, &o.Rules)
	}
	
	if DefaultArrayAction, ok := EnrichfieldrulesMap["defaultArrayAction"].(string); ok {
		o.DefaultArrayAction = &DefaultArrayAction
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Enrichfieldrules) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
