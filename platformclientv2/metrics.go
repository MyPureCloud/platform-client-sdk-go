package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Metrics
type Metrics struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Order - The order of metric within a performance profile
	Order *int `json:"order,omitempty"`


	// MetricDefinitionName - The name of associated metric definition
	MetricDefinitionName *string `json:"metricDefinitionName,omitempty"`


	// MetricDefinitionId - The id of associated metric definition
	MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`


	// ExternalMetricDefinitionId - The id of associated external metric definition
	ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`


	// UnitType - Corresponding unit type for this metric
	UnitType *string `json:"unitType,omitempty"`


	// Enabled - A flag for whether this metric is enabled for a performance profile
	Enabled *bool `json:"enabled,omitempty"`


	// TemplateName - The name of associated objective template
	TemplateName *string `json:"templateName,omitempty"`


	// MaxPoints - Achievable maximum points for this metric
	MaxPoints *int `json:"maxPoints,omitempty"`


	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// UnitDefinition - Unit definition of linked external metric
	UnitDefinition *string `json:"unitDefinition,omitempty"`


	// Precision - Precision of linked external metric
	Precision *int `json:"precision,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Metrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metrics
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		MetricDefinitionName *string `json:"metricDefinitionName,omitempty"`
		
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		ExternalMetricDefinitionId *string `json:"externalMetricDefinitionId,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		TemplateName *string `json:"templateName,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		UnitDefinition *string `json:"unitDefinition,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Order: o.Order,
		
		MetricDefinitionName: o.MetricDefinitionName,
		
		MetricDefinitionId: o.MetricDefinitionId,
		
		ExternalMetricDefinitionId: o.ExternalMetricDefinitionId,
		
		UnitType: o.UnitType,
		
		Enabled: o.Enabled,
		
		TemplateName: o.TemplateName,
		
		MaxPoints: o.MaxPoints,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		UnitDefinition: o.UnitDefinition,
		
		Precision: o.Precision,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Metrics) UnmarshalJSON(b []byte) error {
	var MetricsMap map[string]interface{}
	err := json.Unmarshal(b, &MetricsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MetricsMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MetricsMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Order, ok := MetricsMap["order"].(float64); ok {
		OrderInt := int(Order)
		o.Order = &OrderInt
	}
	
	if MetricDefinitionName, ok := MetricsMap["metricDefinitionName"].(string); ok {
		o.MetricDefinitionName = &MetricDefinitionName
	}
	
	if MetricDefinitionId, ok := MetricsMap["metricDefinitionId"].(string); ok {
		o.MetricDefinitionId = &MetricDefinitionId
	}
	
	if ExternalMetricDefinitionId, ok := MetricsMap["externalMetricDefinitionId"].(string); ok {
		o.ExternalMetricDefinitionId = &ExternalMetricDefinitionId
	}
	
	if UnitType, ok := MetricsMap["unitType"].(string); ok {
		o.UnitType = &UnitType
	}
	
	if Enabled, ok := MetricsMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if TemplateName, ok := MetricsMap["templateName"].(string); ok {
		o.TemplateName = &TemplateName
	}
	
	if MaxPoints, ok := MetricsMap["maxPoints"].(float64); ok {
		MaxPointsInt := int(MaxPoints)
		o.MaxPoints = &MaxPointsInt
	}
	
	if PerformanceProfileId, ok := MetricsMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
	
	if UnitDefinition, ok := MetricsMap["unitDefinition"].(string); ok {
		o.UnitDefinition = &UnitDefinition
	}
	
	if Precision, ok := MetricsMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if SelfUri, ok := MetricsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Metrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
