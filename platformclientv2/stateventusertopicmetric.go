package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventusertopicmetric
type Stateventusertopicmetric struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *map[string]float32 `json:"stats,omitempty"`

}

func (o *Stateventusertopicmetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventusertopicmetric
	
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

func (o *Stateventusertopicmetric) UnmarshalJSON(b []byte) error {
	var StateventusertopicmetricMap map[string]interface{}
	err := json.Unmarshal(b, &StateventusertopicmetricMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := StateventusertopicmetricMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Qualifier, ok := StateventusertopicmetricMap["qualifier"].(string); ok {
		o.Qualifier = &Qualifier
	}
	
	if Stats, ok := StateventusertopicmetricMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventusertopicmetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
