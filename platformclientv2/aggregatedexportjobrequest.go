package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregatedexportjobrequest
type Aggregatedexportjobrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Granularity - Granularity of the exported data, defaults to day
	Granularity *string `json:"granularity,omitempty"`

	// TimeZone - The requested time zone of the exported data, in Olson format. Defaults to business unit time zone
	TimeZone *string `json:"timeZone,omitempty"`

	// Delimiter - The delimiter to use between fields in the export, defaults to comma
	Delimiter *string `json:"delimiter,omitempty"`

	// PlanningGroupIds - The IDs of the planning groups to include in the export, defaults to all planning groups in the business unit
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`

	// DateStart - Start date-time of the export range in ISO-8601 format
	DateStart *time.Time `json:"dateStart,omitempty"`

	// DateEnd - End date-time of the export range in ISO-8601 format
	DateEnd *time.Time `json:"dateEnd,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Aggregatedexportjobrequest) SetField(field string, fieldValue interface{}) {
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

func (o Aggregatedexportjobrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStart","DateEnd", }
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
	type Alias Aggregatedexportjobrequest
	
	DateStart := new(string)
	if o.DateStart != nil {
		
		*DateStart = timeutil.Strftime(o.DateStart, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStart = nil
	}
	
	DateEnd := new(string)
	if o.DateEnd != nil {
		
		*DateEnd = timeutil.Strftime(o.DateEnd, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnd = nil
	}
	
	return json.Marshal(&struct { 
		Granularity *string `json:"granularity,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Delimiter *string `json:"delimiter,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		
		DateStart *string `json:"dateStart,omitempty"`
		
		DateEnd *string `json:"dateEnd,omitempty"`
		Alias
	}{ 
		Granularity: o.Granularity,
		
		TimeZone: o.TimeZone,
		
		Delimiter: o.Delimiter,
		
		PlanningGroupIds: o.PlanningGroupIds,
		
		DateStart: DateStart,
		
		DateEnd: DateEnd,
		Alias:    (Alias)(o),
	})
}

func (o *Aggregatedexportjobrequest) UnmarshalJSON(b []byte) error {
	var AggregatedexportjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AggregatedexportjobrequestMap)
	if err != nil {
		return err
	}
	
	if Granularity, ok := AggregatedexportjobrequestMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if TimeZone, ok := AggregatedexportjobrequestMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if Delimiter, ok := AggregatedexportjobrequestMap["delimiter"].(string); ok {
		o.Delimiter = &Delimiter
	}
    
	if PlanningGroupIds, ok := AggregatedexportjobrequestMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	
	if dateStartString, ok := AggregatedexportjobrequestMap["dateStart"].(string); ok {
		DateStart, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartString)
		o.DateStart = &DateStart
	}
	
	if dateEndString, ok := AggregatedexportjobrequestMap["dateEnd"].(string); ok {
		DateEnd, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndString)
		o.DateEnd = &DateEnd
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregatedexportjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
