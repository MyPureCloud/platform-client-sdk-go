package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Surveyaggregatequeryresponse
type Surveyaggregatequeryresponse struct { 
	// Results
	Results *[]Surveyaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Surveyaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Surveyaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Surveyaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Surveyaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var SurveyaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SurveyaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := SurveyaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
