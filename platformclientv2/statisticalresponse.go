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

func (u *Statisticalresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Statisticalresponse

	

	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		Metrics *[]Aggregatemetricdata `json:"metrics,omitempty"`
		
		Views *[]Aggregateviewdata `json:"views,omitempty"`
		*Alias
	}{ 
		Interval: u.Interval,
		
		Metrics: u.Metrics,
		
		Views: u.Views,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Statisticalresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
