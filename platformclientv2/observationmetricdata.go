package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Observationmetricdata
type Observationmetricdata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Metric
	Metric *string `json:"metric,omitempty"`

	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`

	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

	// Truncated - Flag for a truncated list of observations. If truncated, the first half of the list of observations will contain the oldest observations and the second half the newest observations.
	Truncated *bool `json:"truncated,omitempty"`

	// Observations - List of observations sorted by timestamp in ascending order. This list may be truncated.
	Observations *[]Observationvalue `json:"observations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Observationmetricdata) SetField(field string, fieldValue interface{}) {
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

func (o Observationmetricdata) MarshalJSON() ([]byte, error) {
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
	type Alias Observationmetricdata
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Qualifier *string `json:"qualifier,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		
		Truncated *bool `json:"truncated,omitempty"`
		
		Observations *[]Observationvalue `json:"observations,omitempty"`
		Alias
	}{ 
		Metric: o.Metric,
		
		Qualifier: o.Qualifier,
		
		Stats: o.Stats,
		
		Truncated: o.Truncated,
		
		Observations: o.Observations,
		Alias:    (Alias)(o),
	})
}

func (o *Observationmetricdata) UnmarshalJSON(b []byte) error {
	var ObservationmetricdataMap map[string]interface{}
	err := json.Unmarshal(b, &ObservationmetricdataMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := ObservationmetricdataMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if Qualifier, ok := ObservationmetricdataMap["qualifier"].(string); ok {
		o.Qualifier = &Qualifier
	}
    
	if Stats, ok := ObservationmetricdataMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	
	if Truncated, ok := ObservationmetricdataMap["truncated"].(bool); ok {
		o.Truncated = &Truncated
	}
    
	if Observations, ok := ObservationmetricdataMap["observations"].([]interface{}); ok {
		ObservationsString, _ := json.Marshal(Observations)
		json.Unmarshal(ObservationsString, &o.Observations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Observationmetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
