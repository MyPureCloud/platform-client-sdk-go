package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdefinitioncreaterequest
type Externalmetricdefinitioncreaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the External Metric Definition
	Name *string `json:"name,omitempty"`

	// Unit - The unit of the External Metric Definition
	Unit *string `json:"unit,omitempty"`

	// UnitDefinition - The unit definition of the External Metric Definition
	UnitDefinition *string `json:"unitDefinition,omitempty"`

	// Precision - The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
	Precision *int `json:"precision,omitempty"`

	// DefaultObjectiveType - The default objective type of the External Metric Definition
	DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`

	// Enabled - True if the External Metric Definition is enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalmetricdefinitioncreaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Externalmetricdefinitioncreaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Externalmetricdefinitioncreaterequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Unit *string `json:"unit,omitempty"`
		
		UnitDefinition *string `json:"unitDefinition,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Unit: o.Unit,
		
		UnitDefinition: o.UnitDefinition,
		
		Precision: o.Precision,
		
		DefaultObjectiveType: o.DefaultObjectiveType,
		
		Enabled: o.Enabled,
		Alias:    (Alias)(o),
	})
}

func (o *Externalmetricdefinitioncreaterequest) UnmarshalJSON(b []byte) error {
	var ExternalmetricdefinitioncreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdefinitioncreaterequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ExternalmetricdefinitioncreaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Unit, ok := ExternalmetricdefinitioncreaterequestMap["unit"].(string); ok {
		o.Unit = &Unit
	}
    
	if UnitDefinition, ok := ExternalmetricdefinitioncreaterequestMap["unitDefinition"].(string); ok {
		o.UnitDefinition = &UnitDefinition
	}
    
	if Precision, ok := ExternalmetricdefinitioncreaterequestMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if DefaultObjectiveType, ok := ExternalmetricdefinitioncreaterequestMap["defaultObjectiveType"].(string); ok {
		o.DefaultObjectiveType = &DefaultObjectiveType
	}
    
	if Enabled, ok := ExternalmetricdefinitioncreaterequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdefinitioncreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
