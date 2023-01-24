package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Columndatatypespecification
type Columndatatypespecification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ColumnName - The column name of a column selected for dynamic queueing
	ColumnName *string `json:"columnName,omitempty"`

	// ColumnDataType - The data type of the column selected for dynamic queueing (TEXT, NUMERIC or TIMESTAMP)
	ColumnDataType *string `json:"columnDataType,omitempty"`

	// Min - The minimum length of the numeric column selected for dynamic queueing
	Min *int `json:"min,omitempty"`

	// Max - The maximum length of the numeric column selected for dynamic queueing
	Max *int `json:"max,omitempty"`

	// MaxLength - The maximum length of the text column selected for dynamic queueing
	MaxLength *int `json:"maxLength,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Columndatatypespecification) SetField(field string, fieldValue interface{}) {
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

func (o Columndatatypespecification) MarshalJSON() ([]byte, error) {
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
	type Alias Columndatatypespecification
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		ColumnDataType *string `json:"columnDataType,omitempty"`
		
		Min *int `json:"min,omitempty"`
		
		Max *int `json:"max,omitempty"`
		
		MaxLength *int `json:"maxLength,omitempty"`
		Alias
	}{ 
		ColumnName: o.ColumnName,
		
		ColumnDataType: o.ColumnDataType,
		
		Min: o.Min,
		
		Max: o.Max,
		
		MaxLength: o.MaxLength,
		Alias:    (Alias)(o),
	})
}

func (o *Columndatatypespecification) UnmarshalJSON(b []byte) error {
	var ColumndatatypespecificationMap map[string]interface{}
	err := json.Unmarshal(b, &ColumndatatypespecificationMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := ColumndatatypespecificationMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if ColumnDataType, ok := ColumndatatypespecificationMap["columnDataType"].(string); ok {
		o.ColumnDataType = &ColumnDataType
	}
    
	if Min, ok := ColumndatatypespecificationMap["min"].(float64); ok {
		MinInt := int(Min)
		o.Min = &MinInt
	}
	
	if Max, ok := ColumndatatypespecificationMap["max"].(float64); ok {
		MaxInt := int(Max)
		o.Max = &MaxInt
	}
	
	if MaxLength, ok := ColumndatatypespecificationMap["maxLength"].(float64); ok {
		MaxLengthInt := int(MaxLength)
		o.MaxLength = &MaxLengthInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Columndatatypespecification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
