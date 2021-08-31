package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Nluqualityreportsummary
type Nluqualityreportsummary struct { 
	// Metrics - The list of metrics in the summary
	Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`

}

func (o *Nluqualityreportsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nluqualityreportsummary
	
	return json.Marshal(&struct { 
		Metrics *[]Nluqualityreportsummarymetric `json:"metrics,omitempty"`
		*Alias
	}{ 
		Metrics: o.Metrics,
		Alias:    (*Alias)(o),
	})
}

func (o *Nluqualityreportsummary) UnmarshalJSON(b []byte) error {
	var NluqualityreportsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &NluqualityreportsummaryMap)
	if err != nil {
		return err
	}
	
	if Metrics, ok := NluqualityreportsummaryMap["metrics"].([]interface{}); ok {
		MetricsString, _ := json.Marshal(Metrics)
		json.Unmarshal(MetricsString, &o.Metrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nluqualityreportsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
