package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermforecastresult
type Longtermforecastresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PlanningGroups - The forecast data broken up by planning group
	PlanningGroups *[]Longtermforecastplanninggroupdata `json:"planningGroups,omitempty"`

	// ReferenceStartDate - The reference start date relative to the business unit time zone in this forecast. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

	// WeekCount - The number of weeks in this forecast
	WeekCount *int `json:"weekCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Longtermforecastresult) SetField(field string, fieldValue interface{}) {
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

func (o Longtermforecastresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "ReferenceStartDate", }

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
	type Alias Longtermforecastresult
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%d")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		PlanningGroups *[]Longtermforecastplanninggroupdata `json:"planningGroups,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		Alias
	}{ 
		PlanningGroups: o.PlanningGroups,
		
		ReferenceStartDate: ReferenceStartDate,
		
		WeekCount: o.WeekCount,
		Alias:    (Alias)(o),
	})
}

func (o *Longtermforecastresult) UnmarshalJSON(b []byte) error {
	var LongtermforecastresultMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermforecastresultMap)
	if err != nil {
		return err
	}
	
	if PlanningGroups, ok := LongtermforecastresultMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	
	if referenceStartDateString, ok := LongtermforecastresultMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if WeekCount, ok := LongtermforecastresultMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermforecastresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
