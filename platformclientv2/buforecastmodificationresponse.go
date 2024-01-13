package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastmodificationresponse
type Buforecastmodificationresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the modification
	VarType *string `json:"type,omitempty"`

	// StartIntervalIndex - The number of intervals past referenceStartDate representing the first interval to which this modification applies
	StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`

	// EndIntervalIndex - The number of intervals past referenceStartDate representing the last interval to which this modification applies
	EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`

	// Metric - The metric to which this modification applies
	Metric *string `json:"metric,omitempty"`

	// LegacyMetric - The legacy metric to which this modification applies if applicable
	LegacyMetric *string `json:"legacyMetric,omitempty"`

	// Value - The value of the modification
	Value *float64 `json:"value,omitempty"`

	// Values - The list of modification values. Only applicable for grid-type modifications
	Values *[]Wfmforecastmodificationintervaloffsetvalue `json:"values,omitempty"`

	// SecondaryValues - The list of modification secondary values. Only applicable for multi granularity modifications
	SecondaryValues *[]Wfmforecastmodificationintervaloffsetvalue `json:"secondaryValues,omitempty"`

	// DisplayGranularity - The client side display granularity of the modification, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	DisplayGranularity *string `json:"displayGranularity,omitempty"`

	// Granularity - The actual granularity of the modification as stored behind the scenes, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	Granularity *string `json:"granularity,omitempty"`

	// SecondaryGranularity - The granularity of the 'secondaryValues' modification as stored behind the scenes, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	SecondaryGranularity *string `json:"secondaryGranularity,omitempty"`

	// Enabled - Whether the modification is enabled for the forecast
	Enabled *bool `json:"enabled,omitempty"`

	// PlanningGroupIds - The IDs of the planning groups to which this forecast modification applies
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buforecastmodificationresponse) SetField(field string, fieldValue interface{}) {
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

func (o Buforecastmodificationresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Buforecastmodificationresponse
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`
		
		EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		LegacyMetric *string `json:"legacyMetric,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Values *[]Wfmforecastmodificationintervaloffsetvalue `json:"values,omitempty"`
		
		SecondaryValues *[]Wfmforecastmodificationintervaloffsetvalue `json:"secondaryValues,omitempty"`
		
		DisplayGranularity *string `json:"displayGranularity,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		SecondaryGranularity *string `json:"secondaryGranularity,omitempty"`
		
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
		
		SecondaryValues: o.SecondaryValues,
		
		DisplayGranularity: o.DisplayGranularity,
		
		Granularity: o.Granularity,
		
		SecondaryGranularity: o.SecondaryGranularity,
		
		Enabled: o.Enabled,
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
	})
}

func (o *Buforecastmodificationresponse) UnmarshalJSON(b []byte) error {
	var BuforecastmodificationresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecastmodificationresponseMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := BuforecastmodificationresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if StartIntervalIndex, ok := BuforecastmodificationresponseMap["startIntervalIndex"].(float64); ok {
		StartIntervalIndexInt := int(StartIntervalIndex)
		o.StartIntervalIndex = &StartIntervalIndexInt
	}
	
	if EndIntervalIndex, ok := BuforecastmodificationresponseMap["endIntervalIndex"].(float64); ok {
		EndIntervalIndexInt := int(EndIntervalIndex)
		o.EndIntervalIndex = &EndIntervalIndexInt
	}
	
	if Metric, ok := BuforecastmodificationresponseMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if LegacyMetric, ok := BuforecastmodificationresponseMap["legacyMetric"].(string); ok {
		o.LegacyMetric = &LegacyMetric
	}
    
	if Value, ok := BuforecastmodificationresponseMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if Values, ok := BuforecastmodificationresponseMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if SecondaryValues, ok := BuforecastmodificationresponseMap["secondaryValues"].([]interface{}); ok {
		SecondaryValuesString, _ := json.Marshal(SecondaryValues)
		json.Unmarshal(SecondaryValuesString, &o.SecondaryValues)
	}
	
	if DisplayGranularity, ok := BuforecastmodificationresponseMap["displayGranularity"].(string); ok {
		o.DisplayGranularity = &DisplayGranularity
	}
    
	if Granularity, ok := BuforecastmodificationresponseMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if SecondaryGranularity, ok := BuforecastmodificationresponseMap["secondaryGranularity"].(string); ok {
		o.SecondaryGranularity = &SecondaryGranularity
	}
    
	if Enabled, ok := BuforecastmodificationresponseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if PlanningGroupIds, ok := BuforecastmodificationresponseMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecastmodificationresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
