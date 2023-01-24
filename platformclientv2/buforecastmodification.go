package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastmodification
type Buforecastmodification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the modification
	VarType *string `json:"type,omitempty"`

	// StartIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the first interval to which to apply this modification. Must be null if values is populated
	StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`

	// EndIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the last interval to which to apply this modification.  Must be null if values is populated
	EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`

	// Metric - The metric to which this modification applies
	Metric *string `json:"metric,omitempty"`

	// LegacyMetric - The legacy metric to which this modification applies if applicable
	LegacyMetric *string `json:"legacyMetric,omitempty"`

	// Value - The value of the modification.  Must be null if \"values\" is populated
	Value *float64 `json:"value,omitempty"`

	// Values - The list of values to update.  Only applicable for grid-type modifications. Must be null if \"value\" is populated
	Values *[]Wfmforecastmodificationintervaloffsetvalue `json:"values,omitempty"`

	// DisplayGranularity - The client side display granularity of the modification, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	DisplayGranularity *string `json:"displayGranularity,omitempty"`

	// Granularity - The actual granularity of the modification as stored behind the scenes, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Granularity *string `json:"granularity,omitempty"`

	// Enabled - Whether the modification is enabled for the forecast
	Enabled *bool `json:"enabled,omitempty"`

	// PlanningGroupIds - The IDs of the planning groups to which this forecast modification applies.  Leave empty to apply to all
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buforecastmodification) SetField(field string, fieldValue interface{}) {
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

func (o Buforecastmodification) MarshalJSON() ([]byte, error) {
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
	type Alias Buforecastmodification
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`
		
		EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		LegacyMetric *string `json:"legacyMetric,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Values *[]Wfmforecastmodificationintervaloffsetvalue `json:"values,omitempty"`
		
		DisplayGranularity *string `json:"displayGranularity,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		StartIntervalIndex: o.StartIntervalIndex,
		
		EndIntervalIndex: o.EndIntervalIndex,
		
		Metric: o.Metric,
		
		LegacyMetric: o.LegacyMetric,
		
		Value: o.Value,
		
		Values: o.Values,
		
		DisplayGranularity: o.DisplayGranularity,
		
		Granularity: o.Granularity,
		
		Enabled: o.Enabled,
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
	})
}

func (o *Buforecastmodification) UnmarshalJSON(b []byte) error {
	var BuforecastmodificationMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecastmodificationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := BuforecastmodificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if StartIntervalIndex, ok := BuforecastmodificationMap["startIntervalIndex"].(float64); ok {
		StartIntervalIndexInt := int(StartIntervalIndex)
		o.StartIntervalIndex = &StartIntervalIndexInt
	}
	
	if EndIntervalIndex, ok := BuforecastmodificationMap["endIntervalIndex"].(float64); ok {
		EndIntervalIndexInt := int(EndIntervalIndex)
		o.EndIntervalIndex = &EndIntervalIndexInt
	}
	
	if Metric, ok := BuforecastmodificationMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if LegacyMetric, ok := BuforecastmodificationMap["legacyMetric"].(string); ok {
		o.LegacyMetric = &LegacyMetric
	}
    
	if Value, ok := BuforecastmodificationMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Values, ok := BuforecastmodificationMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if DisplayGranularity, ok := BuforecastmodificationMap["displayGranularity"].(string); ok {
		o.DisplayGranularity = &DisplayGranularity
	}
    
	if Granularity, ok := BuforecastmodificationMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if Enabled, ok := BuforecastmodificationMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if PlanningGroupIds, ok := BuforecastmodificationMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecastmodification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
