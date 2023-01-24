package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestquerybody
type Timeoffrequestquerybody struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Ids - The set of ids to filter time off requests
	Ids *[]string `json:"ids,omitempty"`

	// UserIds - The set of user ids to filter time off requests
	UserIds *[]string `json:"userIds,omitempty"`

	// Statuses - The set of statuses to filter time off requests
	Statuses *[]string `json:"statuses,omitempty"`

	// Substatuses - The set of substatuses to filter time off requests
	Substatuses *[]string `json:"substatuses,omitempty"`

	// DateRange - The inclusive range of dates to filter time off requests
	DateRange *Daterange `json:"dateRange,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Timeoffrequestquerybody) SetField(field string, fieldValue interface{}) {
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

func (o Timeoffrequestquerybody) MarshalJSON() ([]byte, error) {
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
	type Alias Timeoffrequestquerybody
	
	return json.Marshal(&struct { 
		Ids *[]string `json:"ids,omitempty"`
		
		UserIds *[]string `json:"userIds,omitempty"`
		
		Statuses *[]string `json:"statuses,omitempty"`
		
		Substatuses *[]string `json:"substatuses,omitempty"`
		
		DateRange *Daterange `json:"dateRange,omitempty"`
		Alias
	}{ 
		Ids: o.Ids,
		
		UserIds: o.UserIds,
		
		Statuses: o.Statuses,
		
		Substatuses: o.Substatuses,
		
		DateRange: o.DateRange,
		Alias:    (Alias)(o),
	})
}

func (o *Timeoffrequestquerybody) UnmarshalJSON(b []byte) error {
	var TimeoffrequestquerybodyMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestquerybodyMap)
	if err != nil {
		return err
	}
	
	if Ids, ok := TimeoffrequestquerybodyMap["ids"].([]interface{}); ok {
		IdsString, _ := json.Marshal(Ids)
		json.Unmarshal(IdsString, &o.Ids)
	}
	
	if UserIds, ok := TimeoffrequestquerybodyMap["userIds"].([]interface{}); ok {
		UserIdsString, _ := json.Marshal(UserIds)
		json.Unmarshal(UserIdsString, &o.UserIds)
	}
	
	if Statuses, ok := TimeoffrequestquerybodyMap["statuses"].([]interface{}); ok {
		StatusesString, _ := json.Marshal(Statuses)
		json.Unmarshal(StatusesString, &o.Statuses)
	}
	
	if Substatuses, ok := TimeoffrequestquerybodyMap["substatuses"].([]interface{}); ok {
		SubstatusesString, _ := json.Marshal(Substatuses)
		json.Unmarshal(SubstatusesString, &o.Substatuses)
	}
	
	if DateRange, ok := TimeoffrequestquerybodyMap["dateRange"].(map[string]interface{}); ok {
		DateRangeString, _ := json.Marshal(DateRange)
		json.Unmarshal(DateRangeString, &o.DateRange)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestquerybody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
