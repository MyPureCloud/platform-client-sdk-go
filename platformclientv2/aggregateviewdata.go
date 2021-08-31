package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Aggregateviewdata
type Aggregateviewdata struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// Stats
	Stats *Statisticalsummary `json:"stats,omitempty"`

}

func (o *Aggregateviewdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Aggregateviewdata
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Stats *Statisticalsummary `json:"stats,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Stats: o.Stats,
		Alias:    (*Alias)(o),
	})
}

func (o *Aggregateviewdata) UnmarshalJSON(b []byte) error {
	var AggregateviewdataMap map[string]interface{}
	err := json.Unmarshal(b, &AggregateviewdataMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AggregateviewdataMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Stats, ok := AggregateviewdataMap["stats"].(map[string]interface{}); ok {
		StatsString, _ := json.Marshal(Stats)
		json.Unmarshal(StatsString, &o.Stats)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Aggregateviewdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
