package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticssessionmetric
type Analyticssessionmetric struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EmitDate - Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EmitDate *time.Time `json:"emitDate,omitempty"`

	// Name - Unique name of this metric
	Name *string `json:"name,omitempty"`

	// Value - The metric value
	Value *int `json:"value,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticssessionmetric) SetField(field string, fieldValue interface{}) {
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

func (o Analyticssessionmetric) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EmitDate", }
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
	type Alias Analyticssessionmetric
	
	EmitDate := new(string)
	if o.EmitDate != nil {
		
		*EmitDate = timeutil.Strftime(o.EmitDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EmitDate = nil
	}
	
	return json.Marshal(&struct { 
		EmitDate *string `json:"emitDate,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Value *int `json:"value,omitempty"`
		Alias
	}{ 
		EmitDate: EmitDate,
		
		Name: o.Name,
		
		Value: o.Value,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticssessionmetric) UnmarshalJSON(b []byte) error {
	var AnalyticssessionmetricMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticssessionmetricMap)
	if err != nil {
		return err
	}
	
	if emitDateString, ok := AnalyticssessionmetricMap["emitDate"].(string); ok {
		EmitDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", emitDateString)
		o.EmitDate = &EmitDate
	}
	
	if Name, ok := AnalyticssessionmetricMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Value, ok := AnalyticssessionmetricMap["value"].(float64); ok {
		ValueInt := int(Value)
		o.Value = &ValueInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticssessionmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
