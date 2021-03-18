package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Statisticalresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
