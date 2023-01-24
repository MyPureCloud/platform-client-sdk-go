package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactablestatus
type Contactablestatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Contactable - Indicates whether or not the entire contact is contactable for the associated media type.
	Contactable *bool `json:"contactable,omitempty"`

	// ColumnStatus - A map of individual contact method columns to whether the individual column is contactable for the associated media type.
	ColumnStatus *map[string]Columnstatus `json:"columnStatus,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactablestatus) SetField(field string, fieldValue interface{}) {
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

func (o Contactablestatus) MarshalJSON() ([]byte, error) {
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
	type Alias Contactablestatus
	
	return json.Marshal(&struct { 
		Contactable *bool `json:"contactable,omitempty"`
		
		ColumnStatus *map[string]Columnstatus `json:"columnStatus,omitempty"`
		Alias
	}{ 
		Contactable: o.Contactable,
		
		ColumnStatus: o.ColumnStatus,
		Alias:    (Alias)(o),
	})
}

func (o *Contactablestatus) UnmarshalJSON(b []byte) error {
	var ContactablestatusMap map[string]interface{}
	err := json.Unmarshal(b, &ContactablestatusMap)
	if err != nil {
		return err
	}
	
	if Contactable, ok := ContactablestatusMap["contactable"].(bool); ok {
		o.Contactable = &Contactable
	}
    
	if ColumnStatus, ok := ContactablestatusMap["columnStatus"].(map[string]interface{}); ok {
		ColumnStatusString, _ := json.Marshal(ColumnStatus)
		json.Unmarshal(ColumnStatusString, &o.ColumnStatus)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactablestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
