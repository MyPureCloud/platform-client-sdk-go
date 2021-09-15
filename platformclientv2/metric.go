package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Metric
type Metric struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of this metric
	Name *string `json:"name,omitempty"`


	// MetricDefinitionId - The id of associated metric definition
	MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`


	// ExternalMetricDefinitionId - The id of associated external metric definition
	ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`


	// Objective - Associated objective for this metric
	Objective *Objective `json:"objective,omitempty"`


	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// LinkedMetric - The linked metric entity reference
	LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`


	// DateCreated - The created date of this metric
	DateCreated *int `json:"dateCreated,omitempty"`


	// DateUnlinked - The unlinked workday for this metric if this metric was ever unlinked. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateUnlinked *time.Time `json:"dateUnlinked,omitempty"`


	// SourcePerformanceProfile - The source performance profile when this metric is linked
	SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Metric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metric
	
	DateUnlinked := new(string)
	if o.DateUnlinked != nil {
		*DateUnlinked = timeutil.Strftime(o.DateUnlinked, "%Y-%m-%d")
	} else {
		DateUnlinked = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`
		
		Objective *Objective `json:"objective,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`
		
		DateCreated *int `json:"dateCreated,omitempty"`
		
		DateUnlinked *string `json:"dateUnlinked,omitempty"`
		
		SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MetricDefinitionId: o.MetricDefinitionId,
		
		ExternalMetricDefinitionId: o.ExternalMetricDefinitionId,
		
		Objective: o.Objective,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		LinkedMetric: o.LinkedMetric,
		
		DateCreated: o.DateCreated,
		
		DateUnlinked: DateUnlinked,
		
		SourcePerformanceProfile: o.SourcePerformanceProfile,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Metric) UnmarshalJSON(b []byte) error {
	var MetricMap map[string]interface{}
	err := json.Unmarshal(b, &MetricMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MetricMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MetricMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if MetricDefinitionId, ok := MetricMap["metricDefinitionId"].(string); ok {
		o.MetricDefinitionId = &MetricDefinitionId
	}
	
	if ExternalMetricDefinitionId, ok := MetricMap["externalMetricDefinitionId"].(string); ok {
		o.ExternalMetricDefinitionId = &ExternalMetricDefinitionId
	}
	
	if Objective, ok := MetricMap["objective"].(map[string]interface{}); ok {
		ObjectiveString, _ := json.Marshal(Objective)
		json.Unmarshal(ObjectiveString, &o.Objective)
	}
	
	if PerformanceProfileId, ok := MetricMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
	
	if LinkedMetric, ok := MetricMap["linkedMetric"].(map[string]interface{}); ok {
		LinkedMetricString, _ := json.Marshal(LinkedMetric)
		json.Unmarshal(LinkedMetricString, &o.LinkedMetric)
	}
	
	if DateCreated, ok := MetricMap["dateCreated"].(float64); ok {
		DateCreatedInt := int(DateCreated)
		o.DateCreated = &DateCreatedInt
	}
	
	if dateUnlinkedString, ok := MetricMap["dateUnlinked"].(string); ok {
		DateUnlinked, _ := time.Parse("2006-01-02", dateUnlinkedString)
		o.DateUnlinked = &DateUnlinked
	}
	
	if SourcePerformanceProfile, ok := MetricMap["sourcePerformanceProfile"].(map[string]interface{}); ok {
		SourcePerformanceProfileString, _ := json.Marshal(SourcePerformanceProfile)
		json.Unmarshal(SourcePerformanceProfileString, &o.SourcePerformanceProfile)
	}
	
	if SelfUri, ok := MetricMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Metric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
