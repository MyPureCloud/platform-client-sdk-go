package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TimeZone - The timezone of the window in which any updates to the edges assigned to the site can be applied. The minimum size of the window is 2 hours.
	TimeZone *string `json:"timeZone,omitempty"`

	// Rrule - The recurrence rule for updating the Edges assigned to the site. The only supported frequencies are daily and weekly. Weekly frequencies require a day list with at least oneday specified. All other configurations are not supported.
	Rrule *string `json:"rrule,omitempty"`

	// Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	Start *time.Time `json:"start,omitempty"`

	// End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	End *time.Time `json:"end,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Edgeautoupdateconfig) SetField(field string, fieldValue interface{}) {
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

func (o Edgeautoupdateconfig) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{ "Start","End", }
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
	type Alias Edgeautoupdateconfig
	
	Start := new(string)
	if o.Start != nil {
		*Start = timeutil.Strftime(o.Start, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		Start = nil
	}
	
	End := new(string)
	if o.End != nil {
		*End = timeutil.Strftime(o.End, "%Y-%m-%dT%H:%M:%S.%f")
		
	} else {
		End = nil
	}
	
	return json.Marshal(&struct { 
		TimeZone *string `json:"timeZone,omitempty"`
		
		Rrule *string `json:"rrule,omitempty"`
		
		Start *string `json:"start,omitempty"`
		
		End *string `json:"end,omitempty"`
		Alias
	}{ 
		TimeZone: o.TimeZone,
		
		Rrule: o.Rrule,
		
		Start: Start,
		
		End: End,
		Alias:    (Alias)(o),
	})
}

func (o *Edgeautoupdateconfig) UnmarshalJSON(b []byte) error {
	var EdgeautoupdateconfigMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeautoupdateconfigMap)
	if err != nil {
		return err
	}
	
	if TimeZone, ok := EdgeautoupdateconfigMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if Rrule, ok := EdgeautoupdateconfigMap["rrule"].(string); ok {
		o.Rrule = &Rrule
	}
    
	if startString, ok := EdgeautoupdateconfigMap["start"].(string); ok {
		Start, _ := time.Parse("2006-01-02T15:04:05.999999", startString)
		o.Start = &Start
	}
	
	if endString, ok := EdgeautoupdateconfigMap["end"].(string); ok {
		End, _ := time.Parse("2006-01-02T15:04:05.999999", endString)
		o.End = &End
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeautoupdateconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
