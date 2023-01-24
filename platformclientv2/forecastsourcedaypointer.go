package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Forecastsourcedaypointer
type Forecastsourcedaypointer struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DayOfWeek - The forecast day of week for this source data
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// Weight - The relative weight to apply to this source data item for weighted averages
	Weight *int `json:"weight,omitempty"`

	// Date - The date this source data represents, in yyyy-MM-dd format
	Date *string `json:"date,omitempty"`

	// FileName - The name of the source file this data came from if it originated from a data import
	FileName *string `json:"fileName,omitempty"`

	// DataKey - The key to look up the forecast source data for this source day
	DataKey *string `json:"dataKey,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Forecastsourcedaypointer) SetField(field string, fieldValue interface{}) {
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

func (o Forecastsourcedaypointer) MarshalJSON() ([]byte, error) {
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
	type Alias Forecastsourcedaypointer
	
	return json.Marshal(&struct { 
		DayOfWeek *string `json:"dayOfWeek,omitempty"`
		
		Weight *int `json:"weight,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		FileName *string `json:"fileName,omitempty"`
		
		DataKey *string `json:"dataKey,omitempty"`
		Alias
	}{ 
		DayOfWeek: o.DayOfWeek,
		
		Weight: o.Weight,
		
		Date: o.Date,
		
		FileName: o.FileName,
		
		DataKey: o.DataKey,
		Alias:    (Alias)(o),
	})
}

func (o *Forecastsourcedaypointer) UnmarshalJSON(b []byte) error {
	var ForecastsourcedaypointerMap map[string]interface{}
	err := json.Unmarshal(b, &ForecastsourcedaypointerMap)
	if err != nil {
		return err
	}
	
	if DayOfWeek, ok := ForecastsourcedaypointerMap["dayOfWeek"].(string); ok {
		o.DayOfWeek = &DayOfWeek
	}
    
	if Weight, ok := ForecastsourcedaypointerMap["weight"].(float64); ok {
		WeightInt := int(Weight)
		o.Weight = &WeightInt
	}
	
	if Date, ok := ForecastsourcedaypointerMap["date"].(string); ok {
		o.Date = &Date
	}
    
	if FileName, ok := ForecastsourcedaypointerMap["fileName"].(string); ok {
		o.FileName = &FileName
	}
    
	if DataKey, ok := ForecastsourcedaypointerMap["dataKey"].(string); ok {
		o.DataKey = &DataKey
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Forecastsourcedaypointer) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
