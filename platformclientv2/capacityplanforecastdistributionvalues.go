package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanforecastdistributionvalues
type Capacityplanforecastdistributionvalues struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OfferedDistributions - Hourly distribution of forecast offered volume for the forecast week, with one entry per day of the week. Each value is a proportional weight between 0.0 to 1.0. Each day's 24 values must sum to approximately 1.0 (0.999 to 1.001 to allow for rounding)
	OfferedDistributions *[]Capacityplanforecastdistribution `json:"offeredDistributions,omitempty"`

	// AverageHandleTimeDistributions - Hourly distribution of forecast average handle time for the forecast week, with one entry per day of the week. Each value is a weighting factor relative to the daily average handle time. Each day's 24 values must average to approximately 1.0 (0.999 to 1.001 to allow for rounding)
	AverageHandleTimeDistributions *[]Capacityplanforecastdistribution `json:"averageHandleTimeDistributions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanforecastdistributionvalues) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanforecastdistributionvalues) MarshalJSON() ([]byte, error) {
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
	type Alias Capacityplanforecastdistributionvalues
	
	return json.Marshal(&struct { 
		OfferedDistributions *[]Capacityplanforecastdistribution `json:"offeredDistributions,omitempty"`
		
		AverageHandleTimeDistributions *[]Capacityplanforecastdistribution `json:"averageHandleTimeDistributions,omitempty"`
		Alias
	}{ 
		OfferedDistributions: o.OfferedDistributions,
		
		AverageHandleTimeDistributions: o.AverageHandleTimeDistributions,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanforecastdistributionvalues) UnmarshalJSON(b []byte) error {
	var CapacityplanforecastdistributionvaluesMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanforecastdistributionvaluesMap)
	if err != nil {
		return err
	}
	
	if OfferedDistributions, ok := CapacityplanforecastdistributionvaluesMap["offeredDistributions"].([]interface{}); ok {
		OfferedDistributionsString, _ := json.Marshal(OfferedDistributions)
		json.Unmarshal(OfferedDistributionsString, &o.OfferedDistributions)
	}
	
	if AverageHandleTimeDistributions, ok := CapacityplanforecastdistributionvaluesMap["averageHandleTimeDistributions"].([]interface{}); ok {
		AverageHandleTimeDistributionsString, _ := json.Marshal(AverageHandleTimeDistributions)
		json.Unmarshal(AverageHandleTimeDistributionsString, &o.AverageHandleTimeDistributions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanforecastdistributionvalues) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
