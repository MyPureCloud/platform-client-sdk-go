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


	// Objective - Associated objective for this metric
	Objective *Objective `json:"objective,omitempty"`


	// PerformanceProfileId - Performance profile id of this metric
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`


	// LinkedMetric - The linked metric entity reference
	LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`


	// DateCreated - The created date of this metric. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateUnlinked - The unlinked workday for this metric if this metric was ever unlinked. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateUnlinked *time.Time `json:"dateUnlinked,omitempty"`


	// SourcePerformanceProfile - The source performance profile when this metric is linked
	SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Metric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Metric

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateUnlinked := new(string)
	if u.DateUnlinked != nil {
		*DateUnlinked = timeutil.Strftime(u.DateUnlinked, "%Y-%m-%d")
	} else {
		DateUnlinked = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MetricDefinitionId *string `json:"metricDefinitionId,omitempty"`
		
		Objective *Objective `json:"objective,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		LinkedMetric *Addressableentityref `json:"linkedMetric,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateUnlinked *string `json:"dateUnlinked,omitempty"`
		
		SourcePerformanceProfile *Performanceprofile `json:"sourcePerformanceProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		MetricDefinitionId: u.MetricDefinitionId,
		
		Objective: u.Objective,
		
		PerformanceProfileId: u.PerformanceProfileId,
		
		LinkedMetric: u.LinkedMetric,
		
		DateCreated: DateCreated,
		
		DateUnlinked: DateUnlinked,
		
		SourcePerformanceProfile: u.SourcePerformanceProfile,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Metric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
