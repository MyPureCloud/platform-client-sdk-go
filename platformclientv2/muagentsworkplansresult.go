package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Muagentsworkplansresult
type Muagentsworkplansresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entities
	Entities *[]Agentworkplans `json:"entities,omitempty"`

	// ReferenceStartWeekDate - The reference date in yyyy-MM-dd format rolled back to nearest BU start day of week. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceStartWeekDate *time.Time `json:"referenceStartWeekDate,omitempty"`

	// WorkPlanLookup - Map containing lookup values for agent work plans. The integer keys serves as lookup keys for effective work plan from workPlanLookupKeysPerWeek property
	WorkPlanLookup *map[string]Workplanreference `json:"workPlanLookup,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Muagentsworkplansresult) SetField(field string, fieldValue interface{}) {
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

func (o Muagentsworkplansresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "ReferenceStartWeekDate", }

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
	type Alias Muagentsworkplansresult
	
	ReferenceStartWeekDate := new(string)
	if o.ReferenceStartWeekDate != nil {
		*ReferenceStartWeekDate = timeutil.Strftime(o.ReferenceStartWeekDate, "%Y-%m-%d")
	} else {
		ReferenceStartWeekDate = nil
	}
	
	return json.Marshal(&struct { 
		Entities *[]Agentworkplans `json:"entities,omitempty"`
		
		ReferenceStartWeekDate *string `json:"referenceStartWeekDate,omitempty"`
		
		WorkPlanLookup *map[string]Workplanreference `json:"workPlanLookup,omitempty"`
		Alias
	}{ 
		Entities: o.Entities,
		
		ReferenceStartWeekDate: ReferenceStartWeekDate,
		
		WorkPlanLookup: o.WorkPlanLookup,
		Alias:    (Alias)(o),
	})
}

func (o *Muagentsworkplansresult) UnmarshalJSON(b []byte) error {
	var MuagentsworkplansresultMap map[string]interface{}
	err := json.Unmarshal(b, &MuagentsworkplansresultMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := MuagentsworkplansresultMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if referenceStartWeekDateString, ok := MuagentsworkplansresultMap["referenceStartWeekDate"].(string); ok {
		ReferenceStartWeekDate, _ := time.Parse("2006-01-02", referenceStartWeekDateString)
		o.ReferenceStartWeekDate = &ReferenceStartWeekDate
	}
	
	if WorkPlanLookup, ok := MuagentsworkplansresultMap["workPlanLookup"].(map[string]interface{}); ok {
		WorkPlanLookupString, _ := json.Marshal(WorkPlanLookup)
		json.Unmarshal(WorkPlanLookupString, &o.WorkPlanLookup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Muagentsworkplansresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
