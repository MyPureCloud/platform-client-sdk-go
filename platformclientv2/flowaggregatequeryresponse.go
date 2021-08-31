package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowaggregatequeryresponse
type Flowaggregatequeryresponse struct { 
	// Results
	Results *[]Flowaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Flowaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Flowaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var FlowaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FlowaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := FlowaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
