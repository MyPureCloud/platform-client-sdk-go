package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbuforecastmodification
type Wfmbushorttermforecastimportcompletetopicbuforecastmodification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType
	VarType *string `json:"type,omitempty"`

	// StartIntervalIndex
	StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`

	// EndIntervalIndex
	EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`

	// Metric
	Metric *string `json:"metric,omitempty"`

	// LegacyMetric
	LegacyMetric *string `json:"legacyMetric,omitempty"`

	// Value
	Value *float32 `json:"value,omitempty"`

	// Values
	Values *[]Wfmbushorttermforecastimportcompletetopicmodificationintervaloffsetvalue `json:"values,omitempty"`

	// Enabled
	Enabled *bool `json:"enabled,omitempty"`

	// Granularity
	Granularity *string `json:"granularity,omitempty"`

	// DisplayGranularity
	DisplayGranularity *string `json:"displayGranularity,omitempty"`

	// PlanningGroupIds
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecastimportcompletetopicbuforecastmodification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecastimportcompletetopicbuforecastmodification) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbushorttermforecastimportcompletetopicbuforecastmodification
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		StartIntervalIndex *int `json:"startIntervalIndex,omitempty"`
		
		EndIntervalIndex *int `json:"endIntervalIndex,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		LegacyMetric *string `json:"legacyMetric,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		Values *[]Wfmbushorttermforecastimportcompletetopicmodificationintervaloffsetvalue `json:"values,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		DisplayGranularity *string `json:"displayGranularity,omitempty"`
		
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
		
		Enabled: o.Enabled,
		
		Granularity: o.Granularity,
		
		DisplayGranularity: o.DisplayGranularity,
		
		PlanningGroupIds: o.PlanningGroupIds,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbushorttermforecastimportcompletetopicbuforecastmodification) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if StartIntervalIndex, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["startIntervalIndex"].(float64); ok {
		StartIntervalIndexInt := int(StartIntervalIndex)
		o.StartIntervalIndex = &StartIntervalIndexInt
	}
	
	if EndIntervalIndex, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["endIntervalIndex"].(float64); ok {
		EndIntervalIndexInt := int(EndIntervalIndex)
		o.EndIntervalIndex = &EndIntervalIndexInt
	}
	
	if Metric, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if LegacyMetric, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["legacyMetric"].(string); ok {
		o.LegacyMetric = &LegacyMetric
	}
    
	if Value, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if Values, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if Enabled, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Granularity, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if DisplayGranularity, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["displayGranularity"].(string); ok {
		o.DisplayGranularity = &DisplayGranularity
	}
    
	if PlanningGroupIds, ok := WfmbushorttermforecastimportcompletetopicbuforecastmodificationMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbuforecastmodification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
