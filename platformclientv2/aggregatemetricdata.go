package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregatemetricdata
type Aggregatemetricdata struct { 
	// Metric
	Metric *string `json:"metric,omitempty"`


	// Qualifier
	Qualifier *string `json:"qualifier,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

func (o *Aggregatemetricdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregatemetricdata
	
	return json.Marshal(&struct { 
		Metric *string `json:"metric,omitempty"`
		
		Qualifier *string `json:"qualifier,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		*Alias
	}{ 
		Metric: o.Metric,
		
		Qualifier: o.Qualifier,
		
		Stats: o.Stats,
		Alias:    (*Alias)(o),
	})
}

func (o *Aggregatemetricdata) UnmarshalJSON(b []byte) error {
	var AggregatemetricdataMap map[string]interface{}
	err := json.Unmarshal(b, &AggregatemetricdataMap)
	if err != nil {
		return err
	}
	
	if Metric, ok := AggregatemetricdataMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if Qualifier, ok := AggregatemetricdataMap["qualifier"].(string); ok {
		o.Qualifier = &Qualifier
	}
	
	if Stats, ok := AggregatemetricdataMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregatemetricdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
