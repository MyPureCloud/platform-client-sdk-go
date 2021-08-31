package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowobservationqueryresponse
type Flowobservationqueryresponse struct { 
	// Results
	Results *[]Flowobservationdatacontainer `json:"results,omitempty"`

}

func (o *Flowobservationqueryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Flowobservationqueryresponse
	
	return json.Marshal(&struct { 
		Results *[]Flowobservationdatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Flowobservationqueryresponse) UnmarshalJSON(b []byte) error {
	var FlowobservationqueryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &FlowobservationqueryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := FlowobservationqueryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowobservationqueryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
