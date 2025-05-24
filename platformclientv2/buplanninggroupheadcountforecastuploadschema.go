package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buplanninggroupheadcountforecastuploadschema
type Buplanninggroupheadcountforecastuploadschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RequiredPerInterval - Required headcount per interval, referenced against the reference start date
	RequiredPerInterval *[]float64 `json:"requiredPerInterval,omitempty"`

	// RequiredWithoutShrinkagePerInterval - Required headcount per interval without accounting for shrinkage, referenced against the reference start date
	RequiredWithoutShrinkagePerInterval *[]float64 `json:"requiredWithoutShrinkagePerInterval,omitempty"`

	// PlanningGroupId - The ID of the planning group to which this portion of the headcount forecast applies
	PlanningGroupId *string `json:"planningGroupId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buplanninggroupheadcountforecastuploadschema) SetField(field string, fieldValue interface{}) {
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

func (o Buplanninggroupheadcountforecastuploadschema) MarshalJSON() ([]byte, error) {
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
	type Alias Buplanninggroupheadcountforecastuploadschema
	
	return json.Marshal(&struct { 
		RequiredPerInterval *[]float64 `json:"requiredPerInterval,omitempty"`
		
		RequiredWithoutShrinkagePerInterval *[]float64 `json:"requiredWithoutShrinkagePerInterval,omitempty"`
		
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		Alias
	}{ 
		RequiredPerInterval: o.RequiredPerInterval,
		
		RequiredWithoutShrinkagePerInterval: o.RequiredWithoutShrinkagePerInterval,
		
		PlanningGroupId: o.PlanningGroupId,
		Alias:    (Alias)(o),
	})
}

func (o *Buplanninggroupheadcountforecastuploadschema) UnmarshalJSON(b []byte) error {
	var BuplanninggroupheadcountforecastuploadschemaMap map[string]interface{}
	err := json.Unmarshal(b, &BuplanninggroupheadcountforecastuploadschemaMap)
	if err != nil {
		return err
	}
	
	if RequiredPerInterval, ok := BuplanninggroupheadcountforecastuploadschemaMap["requiredPerInterval"].([]interface{}); ok {
		RequiredPerIntervalString, _ := json.Marshal(RequiredPerInterval)
		json.Unmarshal(RequiredPerIntervalString, &o.RequiredPerInterval)
	}
	
	if RequiredWithoutShrinkagePerInterval, ok := BuplanninggroupheadcountforecastuploadschemaMap["requiredWithoutShrinkagePerInterval"].([]interface{}); ok {
		RequiredWithoutShrinkagePerIntervalString, _ := json.Marshal(RequiredWithoutShrinkagePerInterval)
		json.Unmarshal(RequiredWithoutShrinkagePerIntervalString, &o.RequiredWithoutShrinkagePerInterval)
	}
	
	if PlanningGroupId, ok := BuplanninggroupheadcountforecastuploadschemaMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecastuploadschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
