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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Metrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metrics

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Order *int `json:"order,omitempty"`
		
		MetricDefinitionName *string `json:"metricDefinitionName,omitempty"`
		
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		UnitType *string `json:"unitType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		TemplateName *string `json:"templateName,omitempty"`
		
		MaxPoints *int `json:"maxPoints,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Order: u.Order,
		
		MetricDefinitionName: u.MetricDefinitionName,
		
		MetricDefinitionId: u.MetricDefinitionId,
		
		UnitType: u.UnitType,
		
		Enabled: u.Enabled,
		
		TemplateName: u.TemplateName,
		
		MaxPoints: u.MaxPoints,
		
		PerformanceProfileId: u.PerformanceProfileId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Metrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
