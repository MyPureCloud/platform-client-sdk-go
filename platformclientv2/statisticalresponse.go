package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Statisticalresponse
type Statisticalresponse struct { 
	// Interval
	Interval *string `json:"interval,omitempty"`


	// Metrics
	Metrics *[]Aggregatemetricdata `json:"metrics,omitempty"`


	// Views
	Views *[]Aggregateviewdata `json:"views,omitempty"`

}

func (o *Statisticalresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Statisticalresponse
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Aggregatemetricdata `json:"metrics,omitempty"`
		
		Views *[]Aggregateviewdata `json:"views,omitempty"`
		*Alias
	}{ 
		Interval: o.Interval,
		
		Metrics: o.Metrics,
		
		Views: o.Views,
		Alias:    (*Alias)(o),
	})
}

func (o *Statisticalresponse) UnmarshalJSON(b []byte) error {
	var StatisticalresponseMap map[string]interface{}
	err := json.Unmarshal(b, &StatisticalresponseMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := StatisticalresponseMap["interval"].(string); ok {
		o.Interval = &Interval
	}
	
	if Metrics, ok := StatisticalresponseMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	
	if Views, ok := StatisticalresponseMap["views"].([]interface{}); ok {
		ViewsString, _ := json.Marshal(Views)
		json.Unmarshal(ViewsString, &o.Views)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Statisticalresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
