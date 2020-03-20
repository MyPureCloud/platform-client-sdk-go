package platformclientv2
import (
	"encoding/json"
)

// Wfmforecastmodification - A modification to a short term forecast
type Wfmforecastmodification struct { 
	// VarType - The type of the modification
	VarType *string `json:"type,omitempty"`


	// StartIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the first interval to which to apply this modification. Must be null if values is populated
	StartIntervalIndex *int32 `json:"startIntervalIndex,omitempty"`


	// EndIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the last interval to which to apply this modification.  Must be null if values is populated
	EndIntervalIndex *int32 `json:"endIntervalIndex,omitempty"`


	// Metric - The metric to which this modification applies
	Metric *string `json:"metric,omitempty"`


	// Value - The value of the modification.  Must be null if \"values\" is populated
	Value *float64 `json:"value,omitempty"`


	// Values - The list of values to update.  Only applicable for grid-type modifications. Must be null if \"value\" is populated
	Values *[]Wfmforecastmodificationintervaloffsetvalue `json:"values,omitempty"`


	// Enabled - Whether the modification is enabled for the forecast
	Enabled *bool `json:"enabled,omitempty"`


	// Attributes - The attributes defining how this modification applies to the forecast
	Attributes *Wfmforecastmodificationattributes `json:"attributes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmforecastmodification) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
