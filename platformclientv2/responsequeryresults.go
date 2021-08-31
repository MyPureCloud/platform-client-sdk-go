package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsequeryresults - Used to return response query results
type Responsequeryresults struct { 
	// Results - Contains the query results
	Results *Responseentitylist `json:"results,omitempty"`

}

func (o *Responsequeryresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Responsequeryresults
	
	return json.Marshal(&struct { 
		Results *Responseentitylist `json:"results,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		Alias:    (*Alias)(o),
	})
}

func (o *Responsequeryresults) UnmarshalJSON(b []byte) error {
	var ResponsequeryresultsMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsequeryresultsMap)
	if err != nil {
		return err
	}
	
	if Results, ok := ResponsequeryresultsMap["results"].(map[string]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Responsequeryresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
