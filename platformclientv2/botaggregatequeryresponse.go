package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Botaggregatequeryresponse
type Botaggregatequeryresponse struct { 
	// Results
	Results *[]Botaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Botaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Botaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Botaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Botaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var BotaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BotaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := BotaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Botaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
