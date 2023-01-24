package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistconfigchangecontactphonenumbercolumn
type Dialercontactlistconfigchangecontactphonenumbercolumn struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ColumnName - name of the phone column
	ColumnName *string `json:"columnName,omitempty"`

	// VarType - type of the phone column, for example, 'cell' or 'home'
	VarType *string `json:"type,omitempty"`

	// CallableTimeColumn - name of the column indicating the timezone to be considered for determining callable times
	CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) SetField(field string, fieldValue interface{}) {
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

func (o Dialercontactlistconfigchangecontactphonenumbercolumn) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercontactlistconfigchangecontactphonenumbercolumn
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		CallableTimeColumn *string `json:"callableTimeColumn,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		Alias
	}{ 
		ColumnName: o.ColumnName,
		
		VarType: o.VarType,
		
		CallableTimeColumn: o.CallableTimeColumn,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) UnmarshalJSON(b []byte) error {
	var DialercontactlistconfigchangecontactphonenumbercolumnMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistconfigchangecontactphonenumbercolumnMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if VarType, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if CallableTimeColumn, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["callableTimeColumn"].(string); ok {
		o.CallableTimeColumn = &CallableTimeColumn
	}
    
	if AdditionalProperties, ok := DialercontactlistconfigchangecontactphonenumbercolumnMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistconfigchangecontactphonenumbercolumn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
