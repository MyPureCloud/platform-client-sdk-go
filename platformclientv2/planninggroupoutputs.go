package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Planninggroupoutputs
type Planninggroupoutputs struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PlanningGroupId - The ID for for the associated planning group result
	PlanningGroupId *string `json:"planningGroupId,omitempty"`

	// ServiceLevelPerInterval - List of Service Level percentage (0.0-100.0) results per interval
	ServiceLevelPerInterval *[]float64 `json:"serviceLevelPerInterval,omitempty"`

	// OccupancyPerInterval - List of Occupancy percentage (0.0-100.0) results per interval
	OccupancyPerInterval *[]float64 `json:"occupancyPerInterval,omitempty"`

	// AverageSpeedOfAnswerSecondsPerInterval - List of Average Speed of Answer (in seconds) results per interval
	AverageSpeedOfAnswerSecondsPerInterval *[]float64 `json:"averageSpeedOfAnswerSecondsPerInterval,omitempty"`

	// AbandonRatePerInterval - List of Abandon rate percentage (0.0-100.0) results per interval
	AbandonRatePerInterval *[]float64 `json:"abandonRatePerInterval,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Planninggroupoutputs) SetField(field string, fieldValue interface{}) {
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

func (o Planninggroupoutputs) MarshalJSON() ([]byte, error) {
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
	type Alias Planninggroupoutputs
	
	return json.Marshal(&struct { 
		PlanningGroupId *string `json:"planningGroupId,omitempty"`
		
		ServiceLevelPerInterval *[]float64 `json:"serviceLevelPerInterval,omitempty"`
		
		OccupancyPerInterval *[]float64 `json:"occupancyPerInterval,omitempty"`
		
		AverageSpeedOfAnswerSecondsPerInterval *[]float64 `json:"averageSpeedOfAnswerSecondsPerInterval,omitempty"`
		
		AbandonRatePerInterval *[]float64 `json:"abandonRatePerInterval,omitempty"`
		Alias
	}{ 
		PlanningGroupId: o.PlanningGroupId,
		
		ServiceLevelPerInterval: o.ServiceLevelPerInterval,
		
		OccupancyPerInterval: o.OccupancyPerInterval,
		
		AverageSpeedOfAnswerSecondsPerInterval: o.AverageSpeedOfAnswerSecondsPerInterval,
		
		AbandonRatePerInterval: o.AbandonRatePerInterval,
		Alias:    (Alias)(o),
	})
}

func (o *Planninggroupoutputs) UnmarshalJSON(b []byte) error {
	var PlanninggroupoutputsMap map[string]interface{}
	err := json.Unmarshal(b, &PlanninggroupoutputsMap)
	if err != nil {
		return err
	}
	
	if PlanningGroupId, ok := PlanninggroupoutputsMap["planningGroupId"].(string); ok {
		o.PlanningGroupId = &PlanningGroupId
	}
    
	if ServiceLevelPerInterval, ok := PlanninggroupoutputsMap["serviceLevelPerInterval"].([]interface{}); ok {
		ServiceLevelPerIntervalString, _ := json.Marshal(ServiceLevelPerInterval)
		json.Unmarshal(ServiceLevelPerIntervalString, &o.ServiceLevelPerInterval)
	}
	
	if OccupancyPerInterval, ok := PlanninggroupoutputsMap["occupancyPerInterval"].([]interface{}); ok {
		OccupancyPerIntervalString, _ := json.Marshal(OccupancyPerInterval)
		json.Unmarshal(OccupancyPerIntervalString, &o.OccupancyPerInterval)
	}
	
	if AverageSpeedOfAnswerSecondsPerInterval, ok := PlanninggroupoutputsMap["averageSpeedOfAnswerSecondsPerInterval"].([]interface{}); ok {
		AverageSpeedOfAnswerSecondsPerIntervalString, _ := json.Marshal(AverageSpeedOfAnswerSecondsPerInterval)
		json.Unmarshal(AverageSpeedOfAnswerSecondsPerIntervalString, &o.AverageSpeedOfAnswerSecondsPerInterval)
	}
	
	if AbandonRatePerInterval, ok := PlanninggroupoutputsMap["abandonRatePerInterval"].([]interface{}); ok {
		AbandonRatePerIntervalString, _ := json.Marshal(AbandonRatePerInterval)
		json.Unmarshal(AbandonRatePerIntervalString, &o.AbandonRatePerInterval)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Planninggroupoutputs) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
