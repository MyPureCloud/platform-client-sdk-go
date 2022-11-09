package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionaggregatequeryresponse
type Actionaggregatequeryresponse struct { 
	// Results
	Results *[]Actionaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Actionaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Actionaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Actionaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var ActionaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ActionaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := ActionaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
