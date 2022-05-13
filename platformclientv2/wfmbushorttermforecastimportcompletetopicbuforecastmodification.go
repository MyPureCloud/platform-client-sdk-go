package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecastimportcompletetopicbuforecastmodification
type Wfmbushorttermforecastimportcompletetopicbuforecastmodification struct { 
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

func (o *Wfmbushorttermforecastimportcompletetopicbuforecastmodification) MarshalJSON() ([]byte, error) {
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
		*Alias
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
		Alias:    (*Alias)(o),
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
