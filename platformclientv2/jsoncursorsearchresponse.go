package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Jsoncursorsearchresponse
type Jsoncursorsearchresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Types - Resource types the search was performed against
	Types *[]string `json:"types,omitempty"`

	// Results - Search results
	Results *interface{} `json:"results,omitempty"`

	// Aggregations
	Aggregations *interface{} `json:"aggregations,omitempty"`

	// Cursor - The page cursor
	Cursor *string `json:"cursor,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Jsoncursorsearchresponse) SetField(field string, fieldValue interface{}) {
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

func (o Jsoncursorsearchresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Jsoncursorsearchresponse
	
	return json.Marshal(&struct { 
		Types *[]string `json:"types,omitempty"`
		
		Results *interface{} `json:"results,omitempty"`
		
		Aggregations *interface{} `json:"aggregations,omitempty"`
		
		Cursor *string `json:"cursor,omitempty"`
		Alias
	}{ 
		Types: o.Types,
		
		Results: o.Results,
		
		Aggregations: o.Aggregations,
		
		Cursor: o.Cursor,
		Alias:    (Alias)(o),
	})
}

func (o *Jsoncursorsearchresponse) UnmarshalJSON(b []byte) error {
	var JsoncursorsearchresponseMap map[string]interface{}
	err := json.Unmarshal(b, &JsoncursorsearchresponseMap)
	if err != nil {
		return err
	}
	
	if Types, ok := JsoncursorsearchresponseMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if Results, ok := JsoncursorsearchresponseMap["results"].(map[string]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if Aggregations, ok := JsoncursorsearchresponseMap["aggregations"].(map[string]interface{}); ok {
		AggregationsString, _ := json.Marshal(Aggregations)
		json.Unmarshal(AggregationsString, &o.Aggregations)
	}
	
	if Cursor, ok := JsoncursorsearchresponseMap["cursor"].(string); ok {
		o.Cursor = &Cursor
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Jsoncursorsearchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
