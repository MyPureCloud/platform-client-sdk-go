package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Intradayperformancepredictiondata
type Intradayperformancepredictiondata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ServiceLevelPercent - Percentage of interactions that meets service level target as defined in the matching service goal templates
	ServiceLevelPercent *float64 `json:"serviceLevelPercent,omitempty"`

	// AverageSpeedOfAnswerSeconds - Predicted average time in seconds it takes to answer an interaction once the interaction becomes available to be routed
	AverageSpeedOfAnswerSeconds *float64 `json:"averageSpeedOfAnswerSeconds,omitempty"`

	// OccupancyPercent - Percentage of on-queue time for all agents in this group that are occupied handling interactions
	OccupancyPercent *float64 `json:"occupancyPercent,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Intradayperformancepredictiondata) SetField(field string, fieldValue interface{}) {
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

func (o Intradayperformancepredictiondata) MarshalJSON() ([]byte, error) {
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
	type Alias Intradayperformancepredictiondata
	
	return json.Marshal(&struct { 
		ServiceLevelPercent *float64 `json:"serviceLevelPercent,omitempty"`
		
		AverageSpeedOfAnswerSeconds *float64 `json:"averageSpeedOfAnswerSeconds,omitempty"`
		
		OccupancyPercent *float64 `json:"occupancyPercent,omitempty"`
		Alias
	}{ 
		ServiceLevelPercent: o.ServiceLevelPercent,
		
		AverageSpeedOfAnswerSeconds: o.AverageSpeedOfAnswerSeconds,
		
		OccupancyPercent: o.OccupancyPercent,
		Alias:    (Alias)(o),
	})
}

func (o *Intradayperformancepredictiondata) UnmarshalJSON(b []byte) error {
	var IntradayperformancepredictiondataMap map[string]interface{}
	err := json.Unmarshal(b, &IntradayperformancepredictiondataMap)
	if err != nil {
		return err
	}
	
	if ServiceLevelPercent, ok := IntradayperformancepredictiondataMap["serviceLevelPercent"].(float64); ok {
		o.ServiceLevelPercent = &ServiceLevelPercent
	}
    
	if AverageSpeedOfAnswerSeconds, ok := IntradayperformancepredictiondataMap["averageSpeedOfAnswerSeconds"].(float64); ok {
		o.AverageSpeedOfAnswerSeconds = &AverageSpeedOfAnswerSeconds
	}
    
	if OccupancyPercent, ok := IntradayperformancepredictiondataMap["occupancyPercent"].(float64); ok {
		o.OccupancyPercent = &OccupancyPercent
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Intradayperformancepredictiondata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
