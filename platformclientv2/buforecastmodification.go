package platformclientv2
import (
	"encoding/json"
)

// Buforecastmodification
type Buforecastmodification struct { 
	// VarType - The type of the modification
	VarType *string `json:"type,omitempty"`


	// StartIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the first interval to which to apply this modification. Must be null if values is populated
	StartIntervalIndex *int32 `json:"startIntervalIndex,omitempty"`


	// EndIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the last interval to which to apply this modification.  Must be null if values is populated
	EndIntervalIndex *int32 `json:"endIntervalIndex,omitempty"`


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

// String returns a JSON representation of the model
func (o *Buforecastmodification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
