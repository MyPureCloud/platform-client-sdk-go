package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Yearmonth
type Yearmonth struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Year
	Year *int `json:"year,omitempty"`

	// Month
	Month *string `json:"month,omitempty"`

	// MonthValue
	MonthValue *int `json:"monthValue,omitempty"`

	// LeapYear
	LeapYear *bool `json:"leapYear,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Yearmonth) SetField(field string, fieldValue interface{}) {
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

func (o Yearmonth) MarshalJSON() ([]byte, error) {
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
	type Alias Yearmonth
	
	return json.Marshal(&struct { 
		Year *int `json:"year,omitempty"`
		
		Month *string `json:"month,omitempty"`
		
		MonthValue *int `json:"monthValue,omitempty"`
		
		LeapYear *bool `json:"leapYear,omitempty"`
		Alias
	}{ 
		Year: o.Year,
		
		Month: o.Month,
		
		MonthValue: o.MonthValue,
		
		LeapYear: o.LeapYear,
		Alias:    (Alias)(o),
	})
}

func (o *Yearmonth) UnmarshalJSON(b []byte) error {
	var YearmonthMap map[string]interface{}
	err := json.Unmarshal(b, &YearmonthMap)
	if err != nil {
		return err
	}
	
	if Year, ok := YearmonthMap["year"].(float64); ok {
		YearInt := int(Year)
		o.Year = &YearInt
	}
	
	if Month, ok := YearmonthMap["month"].(string); ok {
		o.Month = &Month
	}
    
	if MonthValue, ok := YearmonthMap["monthValue"].(float64); ok {
		MonthValueInt := int(MonthValue)
		o.MonthValue = &MonthValueInt
	}
	
	if LeapYear, ok := YearmonthMap["leapYear"].(bool); ok {
		o.LeapYear = &LeapYear
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Yearmonth) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
