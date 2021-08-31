package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingappointmentaggregateresponse
type Coachingappointmentaggregateresponse struct { 
	// Results - The results of the query
	Results *[]Queryresponsegroupeddata `json:"results,omitempty"`

}

func (o *Coachingappointmentaggregateresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingappointmentaggregateresponse
	
	return json.Marshal(&struct { 
		Results *[]Queryresponsegroupeddata `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingappointmentaggregateresponse) UnmarshalJSON(b []byte) error {
	var CoachingappointmentaggregateresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingappointmentaggregateresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := CoachingappointmentaggregateresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregateresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
