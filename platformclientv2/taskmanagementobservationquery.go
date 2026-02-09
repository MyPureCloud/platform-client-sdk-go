package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Taskmanagementobservationquery
type Taskmanagementobservationquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// GroupBy - Dimension(s) to group by. Determines how the results will be grouped in the response.
	GroupBy *[]string `json:"groupBy,omitempty"`

	// Metrics - List of metrics to be retrieved. Specifies which observational metrics should be included in the response.
	Metrics *[]Taskmanagementquerymetric `json:"metrics,omitempty"`

	// Filter - Filter to return a subset of observations.
	Filter *Taskmanagementobservationqueryfilter `json:"filter,omitempty"`

	// Expands - List of properties to expand. Additional details about the objects returned in the results will be included in the response if supplied.
	Expands *[]string `json:"expands,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Taskmanagementobservationquery) SetField(field string, fieldValue interface{}) {
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

func (o Taskmanagementobservationquery) MarshalJSON() ([]byte, error) {
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
	type Alias Taskmanagementobservationquery
	
	return json.Marshal(&struct { 
		GroupBy *[]string `json:"groupBy,omitempty"`
		
		Metrics *[]Taskmanagementquerymetric `json:"metrics,omitempty"`
		
		Filter *Taskmanagementobservationqueryfilter `json:"filter,omitempty"`
		
		Expands *[]string `json:"expands,omitempty"`
		Alias
	}{ 
		GroupBy: o.GroupBy,
		
		Metrics: o.Metrics,
		
		Filter: o.Filter,
		
		Expands: o.Expands,
		Alias:    (Alias)(o),
	})
}

func (o *Taskmanagementobservationquery) UnmarshalJSON(b []byte) error {
	var TaskmanagementobservationqueryMap map[string]interface{}
	err := json.Unmarshal(b, &TaskmanagementobservationqueryMap)
	if err != nil {
		return err
	}
	
	if GroupBy, ok := TaskmanagementobservationqueryMap["groupBy"].([]interface{}); ok {
		GroupByString, _ := json.Marshal(GroupBy)
		json.Unmarshal(GroupByString, &o.GroupBy)
	}
	
	if Metrics, ok := TaskmanagementobservationqueryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Filter, ok := TaskmanagementobservationqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if Expands, ok := TaskmanagementobservationqueryMap["expands"].([]interface{}); ok {
		ExpandsString, _ := json.Marshal(Expands)
		json.Unmarshal(ExpandsString, &o.Expands)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
