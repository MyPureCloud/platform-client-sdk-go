package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Organizationcallmetrics
type Organizationcallmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Usage - The current usage percentage of the organization's call capacity.
	Usage *float64 `json:"usage,omitempty"`

	// AutoScalingTriggerPercentage - The autoscaling trigger percentage of the organization's call capacity.
	AutoScalingTriggerPercentage *float64 `json:"autoScalingTriggerPercentage,omitempty"`

	// CpuIntensity - The current compute intensity of the organization's call capacity.
	CpuIntensity *string `json:"cpuIntensity,omitempty"`

	// MemoryIntensity - The current memory intensity of the organization's call capacity.
	MemoryIntensity *string `json:"memoryIntensity,omitempty"`

	// ConcurrentCallCount - The current number of concurrent calls in the organization.
	ConcurrentCallCount *int `json:"concurrentCallCount,omitempty"`

	// ConcurrentCallSessionCount - The current number of concurrent call sessions in the organization.
	ConcurrentCallSessionCount *int `json:"concurrentCallSessionCount,omitempty"`

	// MaxCallCapacity - The maximum number of concurrent calls allowed in the organization.
	MaxCallCapacity *int `json:"maxCallCapacity,omitempty"`

	// MaxCallSessionCapacity - The maximum number of concurrent call sessions allowed in the organization.
	MaxCallSessionCapacity *int `json:"maxCallSessionCapacity,omitempty"`

	// AutoScaleInProgress - The autoscaling status of the organization's call capacity.
	AutoScaleInProgress *string `json:"autoScaleInProgress,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Organizationcallmetrics) SetField(field string, fieldValue interface{}) {
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

func (o Organizationcallmetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Organizationcallmetrics
	
	return json.Marshal(&struct { 
		Usage *float64 `json:"usage,omitempty"`
		
		AutoScalingTriggerPercentage *float64 `json:"autoScalingTriggerPercentage,omitempty"`
		
		CpuIntensity *string `json:"cpuIntensity,omitempty"`
		
		MemoryIntensity *string `json:"memoryIntensity,omitempty"`
		
		ConcurrentCallCount *int `json:"concurrentCallCount,omitempty"`
		
		ConcurrentCallSessionCount *int `json:"concurrentCallSessionCount,omitempty"`
		
		MaxCallCapacity *int `json:"maxCallCapacity,omitempty"`
		
		MaxCallSessionCapacity *int `json:"maxCallSessionCapacity,omitempty"`
		
		AutoScaleInProgress *string `json:"autoScaleInProgress,omitempty"`
		Alias
	}{ 
		Usage: o.Usage,
		
		AutoScalingTriggerPercentage: o.AutoScalingTriggerPercentage,
		
		CpuIntensity: o.CpuIntensity,
		
		MemoryIntensity: o.MemoryIntensity,
		
		ConcurrentCallCount: o.ConcurrentCallCount,
		
		ConcurrentCallSessionCount: o.ConcurrentCallSessionCount,
		
		MaxCallCapacity: o.MaxCallCapacity,
		
		MaxCallSessionCapacity: o.MaxCallSessionCapacity,
		
		AutoScaleInProgress: o.AutoScaleInProgress,
		Alias:    (Alias)(o),
	})
}

func (o *Organizationcallmetrics) UnmarshalJSON(b []byte) error {
	var OrganizationcallmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &OrganizationcallmetricsMap)
	if err != nil {
		return err
	}
	
	if Usage, ok := OrganizationcallmetricsMap["usage"].(float64); ok {
		o.Usage = &Usage
	}
    
	if AutoScalingTriggerPercentage, ok := OrganizationcallmetricsMap["autoScalingTriggerPercentage"].(float64); ok {
		o.AutoScalingTriggerPercentage = &AutoScalingTriggerPercentage
	}
    
	if CpuIntensity, ok := OrganizationcallmetricsMap["cpuIntensity"].(string); ok {
		o.CpuIntensity = &CpuIntensity
	}
    
	if MemoryIntensity, ok := OrganizationcallmetricsMap["memoryIntensity"].(string); ok {
		o.MemoryIntensity = &MemoryIntensity
	}
    
	if ConcurrentCallCount, ok := OrganizationcallmetricsMap["concurrentCallCount"].(float64); ok {
		ConcurrentCallCountInt := int(ConcurrentCallCount)
		o.ConcurrentCallCount = &ConcurrentCallCountInt
	}
	
	if ConcurrentCallSessionCount, ok := OrganizationcallmetricsMap["concurrentCallSessionCount"].(float64); ok {
		ConcurrentCallSessionCountInt := int(ConcurrentCallSessionCount)
		o.ConcurrentCallSessionCount = &ConcurrentCallSessionCountInt
	}
	
	if MaxCallCapacity, ok := OrganizationcallmetricsMap["maxCallCapacity"].(float64); ok {
		MaxCallCapacityInt := int(MaxCallCapacity)
		o.MaxCallCapacity = &MaxCallCapacityInt
	}
	
	if MaxCallSessionCapacity, ok := OrganizationcallmetricsMap["maxCallSessionCapacity"].(float64); ok {
		MaxCallSessionCapacityInt := int(MaxCallSessionCapacity)
		o.MaxCallSessionCapacity = &MaxCallSessionCapacityInt
	}
	
	if AutoScaleInProgress, ok := OrganizationcallmetricsMap["autoScaleInProgress"].(string); ok {
		o.AutoScaleInProgress = &AutoScaleInProgress
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Organizationcallmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
