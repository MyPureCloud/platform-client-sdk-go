package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Selectedcustomcalculationcolumns
type Selectedcustomcalculationcolumns struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CustomCalculation - Custom calculation added as a column
	CustomCalculation *Addressableentityref `json:"customCalculation,omitempty"`

	// Restricted - Indicates if selected custom calculation column is deleted or access revoked for the user
	Restricted *bool `json:"restricted,omitempty"`

	// SoftDeleted - Is selected custom calculation column soft deleted
	SoftDeleted *bool `json:"softDeleted,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Selectedcustomcalculationcolumns) SetField(field string, fieldValue interface{}) {
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

func (o Selectedcustomcalculationcolumns) MarshalJSON() ([]byte, error) {
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
	type Alias Selectedcustomcalculationcolumns
	
	return json.Marshal(&struct { 
		CustomCalculation *Addressableentityref `json:"customCalculation,omitempty"`
		
		Restricted *bool `json:"restricted,omitempty"`
		
		SoftDeleted *bool `json:"softDeleted,omitempty"`
		Alias
	}{ 
		CustomCalculation: o.CustomCalculation,
		
		Restricted: o.Restricted,
		
		SoftDeleted: o.SoftDeleted,
		Alias:    (Alias)(o),
	})
}

func (o *Selectedcustomcalculationcolumns) UnmarshalJSON(b []byte) error {
	var SelectedcustomcalculationcolumnsMap map[string]interface{}
	err := json.Unmarshal(b, &SelectedcustomcalculationcolumnsMap)
	if err != nil {
		return err
	}
	
	if CustomCalculation, ok := SelectedcustomcalculationcolumnsMap["customCalculation"].(map[string]interface{}); ok {
		CustomCalculationString, _ := json.Marshal(CustomCalculation)
		json.Unmarshal(CustomCalculationString, &o.CustomCalculation)
	}
	
	if Restricted, ok := SelectedcustomcalculationcolumnsMap["restricted"].(bool); ok {
		o.Restricted = &Restricted
	}
    
	if SoftDeleted, ok := SelectedcustomcalculationcolumnsMap["softDeleted"].(bool); ok {
		o.SoftDeleted = &SoftDeleted
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Selectedcustomcalculationcolumns) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
