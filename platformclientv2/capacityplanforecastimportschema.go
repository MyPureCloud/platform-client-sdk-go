package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanforecastimportschema
type Capacityplanforecastimportschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Granularity - The granularity at which forecast values are provided for each planning group
	Granularity *string `json:"granularity,omitempty"`

	// BusinessUnitStartDate - The first day of the forecast relative to the business unit time zone in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitStartDate *time.Time `json:"businessUnitStartDate,omitempty"`

	// PlanningGroups - The forecast data for the planning groups included in this forecast
	PlanningGroups *[]Capacityplanforecastplanninggroup `json:"planningGroups,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanforecastimportschema) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanforecastimportschema) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BusinessUnitStartDate", }

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
	type Alias Capacityplanforecastimportschema
	
	BusinessUnitStartDate := new(string)
	if o.BusinessUnitStartDate != nil {
		*BusinessUnitStartDate = timeutil.Strftime(o.BusinessUnitStartDate, "%Y-%m-%d")
	} else {
		BusinessUnitStartDate = nil
	}
	
	return json.Marshal(&struct { 
		Granularity *string `json:"granularity,omitempty"`
		
		BusinessUnitStartDate *string `json:"businessUnitStartDate,omitempty"`
		
		PlanningGroups *[]Capacityplanforecastplanninggroup `json:"planningGroups,omitempty"`
		Alias
	}{ 
		Granularity: o.Granularity,
		
		BusinessUnitStartDate: BusinessUnitStartDate,
		
		PlanningGroups: o.PlanningGroups,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanforecastimportschema) UnmarshalJSON(b []byte) error {
	var CapacityplanforecastimportschemaMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanforecastimportschemaMap)
	if err != nil {
		return err
	}
	
	if Granularity, ok := CapacityplanforecastimportschemaMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if businessUnitStartDateString, ok := CapacityplanforecastimportschemaMap["businessUnitStartDate"].(string); ok {
		BusinessUnitStartDate, _ := time.Parse("2006-01-02", businessUnitStartDateString)
		o.BusinessUnitStartDate = &BusinessUnitStartDate
	}
	
	if PlanningGroups, ok := CapacityplanforecastimportschemaMap["planningGroups"].([]interface{}); ok {
		PlanningGroupsString, _ := json.Marshal(PlanningGroups)
		json.Unmarshal(PlanningGroupsString, &o.PlanningGroups)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanforecastimportschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
