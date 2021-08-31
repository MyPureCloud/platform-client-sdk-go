package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventcampaigntopicmetric
type Stateventcampaigntopicmetric struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *map[string]float32 `json:"stats,omitempty"`

}

func (o *Stateventcampaigntopicmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventcampaigntopicmetric
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Qualifier *string `json:"qualifier,omitempty"`
		
		Stats *map[string]float32 `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Qualifier: o.Qualifier,
		
		Stats: o.Stats,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventcampaigntopicmetric) UnmarshalJSON(b []byte) error {
	var StateventcampaigntopicmetricMap map[string]interface{}
	err := json.Unmarshal(b, &StateventcampaigntopicmetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := StateventcampaigntopicmetricMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Qualifier, ok := StateventcampaigntopicmetricMap["qualifier"].(string); ok {
		o.Qualifier = &Qualifier
	}
	
	if Stats, ok := StateventcampaigntopicmetricMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
