package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createmetric
type Createmetric struct { 
	// MetricDefinitionId - The id of associated metric definition
	MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`


	// ExternalMetricDefinitionId - The id of associated external metric definition
	ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`


	// Objective - Associated objective for this metric
	Objective *Createobjective `json:"objective,omitempty"`


	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// Name - The name of this metric
	Name *string `json:"name,omitempty"`

}

func (o *Createmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createmetric
	
	return json.Marshal(&struct { 
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`
		
		Objective *Createobjective `json:"objective,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		Name *string `json:"name,omitempty"`
		*Alias
	}{ 
		MetricDefinitionId: o.MetricDefinitionId,
		
		ExternalMetricDefinitionId: o.ExternalMetricDefinitionId,
		
		Objective: o.Objective,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		Name: o.Name,
		Alias:    (*Alias)(o),
	})
}

func (o *Createmetric) UnmarshalJSON(b []byte) error {
	var CreatemetricMap map[string]interface{}
	err := json.Unmarshal(b, &CreatemetricMap)
	if err != nil {
		return err
	}
	
	if MetricDefinitionId, ok := CreatemetricMap["metricDefinitionId"].(string); ok {
		o.MetricDefinitionId = &MetricDefinitionId
	}
    
	if ExternalMetricDefinitionId, ok := CreatemetricMap["externalMetricDefinitionId"].(string); ok {
		o.ExternalMetricDefinitionId = &ExternalMetricDefinitionId
	}
    
	if Objective, ok := CreatemetricMap["objective"].(map[string]interface{}); ok {
		ObjectiveString, _ := json.Marshal(Objective)
		json.Unmarshal(ObjectiveString, &o.Objective)
	}
	
	if PerformanceProfileId, ok := CreatemetricMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
    
	if Name, ok := CreatemetricMap["name"].(string); ok {
		o.Name = &Name
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
