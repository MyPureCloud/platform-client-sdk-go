package platformclientv2
import (
	"encoding/json"
)

// Metric
type Metric struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of this metric
	Name *string `json:"name,omitempty"`


	// MetricDefinitionId - The id of associated metric definition
	MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`


	// Objective - Associated objective for this metric
	Objective *Objective `json:"objective,omitempty"`


	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Metric) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
