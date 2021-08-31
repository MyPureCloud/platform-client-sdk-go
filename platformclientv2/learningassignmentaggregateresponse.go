package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentaggregateresponse
type Learningassignmentaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Learningassignmentaggregatequeryresponsegroupeddata `json:"results,omitempty"`

}

func (o *Learningassignmentaggregateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningassignmentaggregateresponse
	
	return json.Marshal(&struct { 
		Results *[]Learningassignmentaggregatequeryresponsegroupeddata `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Learningassignmentaggregateresponse) UnmarshalJSON(b []byte) error {
	var LearningassignmentaggregateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &LearningassignmentaggregateresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := LearningassignmentaggregateresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
