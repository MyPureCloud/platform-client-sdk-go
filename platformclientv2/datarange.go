package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Datarange
type Datarange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateMin - The first event timestamp found. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateMin *time.Time `json:"dateMin,omitempty"`

	// DateMax - The last event timestamp found. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateMax *time.Time `json:"dateMax,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Datarange) SetField(field string, fieldValue interface{}) {
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

func (o Datarange) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateMin","DateMax", }
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
	type Alias Datarange
	
	DateMin := new(string)
	if o.DateMin != nil {
		
		*DateMin = timeutil.Strftime(o.DateMin, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMin = nil
	}
	
	DateMax := new(string)
	if o.DateMax != nil {
		
		*DateMax = timeutil.Strftime(o.DateMax, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateMax = nil
	}
	
	return json.Marshal(&struct { 
		DateMin *string `json:"dateMin,omitempty"`
		
		DateMax *string `json:"dateMax,omitempty"`
		Alias
	}{ 
		DateMin: DateMin,
		
		DateMax: DateMax,
		Alias:    (Alias)(o),
	})
}

func (o *Datarange) UnmarshalJSON(b []byte) error {
	var DatarangeMap map[string]interface{}
	err := json.Unmarshal(b, &DatarangeMap)
	if err != nil {
		return err
	}
	
	if dateMinString, ok := DatarangeMap["dateMin"].(string); ok {
		DateMin, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMinString)
		o.DateMin = &DateMin
	}
	
	if dateMaxString, ok := DatarangeMap["dateMax"].(string); ok {
		DateMax, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateMaxString)
		o.DateMax = &DateMax
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Datarange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
