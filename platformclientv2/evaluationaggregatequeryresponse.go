package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evaluationaggregatequeryresponse
type Evaluationaggregatequeryresponse struct { 
	// Results
	Results *[]Evaluationaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Evaluationaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evaluationaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Evaluationaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Evaluationaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var EvaluationaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &EvaluationaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := EvaluationaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Evaluationaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
