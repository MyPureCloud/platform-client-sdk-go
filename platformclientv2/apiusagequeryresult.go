package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Apiusagequeryresult
type Apiusagequeryresult struct { 
	// Results - Query results
	Results *[]Apiusagerow `json:"results,omitempty"`


	// QueryStatus - Query status
	QueryStatus *string `json:"queryStatus,omitempty"`

}

func (o *Apiusagequeryresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Apiusagequeryresult
	
	return json.Marshal(&struct { 
		Results *[]Apiusagerow `json:"results,omitempty"`
		
		QueryStatus *string `json:"queryStatus,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		
		QueryStatus: o.QueryStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Apiusagequeryresult) UnmarshalJSON(b []byte) error {
	var ApiusagequeryresultMap map[string]interface{}
	err := json.Unmarshal(b, &ApiusagequeryresultMap)
	if err != nil {
		return err
	}
	
	if Results, ok := ApiusagequeryresultMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if QueryStatus, ok := ApiusagequeryresultMap["queryStatus"].(string); ok {
		o.QueryStatus = &QueryStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Apiusagequeryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
