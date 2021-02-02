package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecastimportcompletetopicbuforecastmodification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
