package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptaggregatequeryresponse
type Transcriptaggregatequeryresponse struct { 
	// Results
	Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`

}

func (o *Transcriptaggregatequeryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Transcriptaggregatequeryresponse
	
	return json.Marshal(&struct { 
		Results *[]Transcriptaggregatedatacontainer `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Transcriptaggregatequeryresponse) UnmarshalJSON(b []byte) error {
	var TranscriptaggregatequeryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptaggregatequeryresponseMap)
	if err != nil {
		return err
	}
	
	if Results, ok := TranscriptaggregatequeryresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptaggregatequeryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
