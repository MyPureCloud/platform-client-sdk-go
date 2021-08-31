package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecastmodification
type Buforecastmodification struct { 
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

func (o *Buforecastmodification) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
