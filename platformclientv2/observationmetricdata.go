package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Observationmetricdata
type Observationmetricdata struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`


	// Truncated - Flag for a truncated list of observations. If truncated, the first half of the list of observations will contain the oldest observations and the second half the newest observations.
	Truncated *bool `json:"truncated,omitempty"`


	// Observations - List of observations sorted by timestamp in ascending order. This list may be truncated.
	Observations *[]Observationvalue `json:"observations,omitempty"`

}

func (o *Observationmetricdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Observationmetricdata
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Qualifier *string `json:"qualifier,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		
		Truncated *bool `json:"truncated,omitempty"`
		
		Observations *[]Observationvalue `json:"observations,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Qualifier: o.Qualifier,
		
		Stats: o.Stats,
		
		Truncated: o.Truncated,
		
		Observations: o.Observations,
		Alias:    (*Alias)(o),
	})
}

func (o *Observationmetricdata) UnmarshalJSON(b []byte) error {
	var ObservationmetricdataMap map[string]interface{}
	err := json.Unmarshal(b, &ObservationmetricdataMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := ObservationmetricdataMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Qualifier, ok := ObservationmetricdataMap["qualifier"].(string); ok {
		o.Qualifier = &Qualifier
	}
	
	if Stats, ok := ObservationmetricdataMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	
	if Truncated, ok := ObservationmetricdataMap["truncated"].(bool); ok {
		o.Truncated = &Truncated
	}
	
	if Observations, ok := ObservationmetricdataMap["observations"].([]interface{}); ok {
		ObservationsString, _ := json.Marshal(Observations)
		json.Unmarshal(ObservationsString, &o.Observations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Observationmetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
