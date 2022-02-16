package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Stateventcampaigntopicintervalmetrics
type Stateventcampaigntopicintervalmetrics struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Stateventcampaigntopicmetricstats `json:"metrics,omitempty"`

}

func (o *Stateventcampaigntopicintervalmetrics) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Stateventcampaigntopicintervalmetrics
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Stateventcampaigntopicmetricstats `json:"metrics,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Stateventcampaigntopicintervalmetrics) UnmarshalJSON(b []byte) error {
	var StateventcampaigntopicintervalmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &StateventcampaigntopicintervalmetricsMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StateventcampaigntopicintervalmetricsMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StateventcampaigntopicintervalmetricsMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stateventcampaigntopicintervalmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
