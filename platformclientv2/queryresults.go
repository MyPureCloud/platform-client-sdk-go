package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryresults
type Queryresults struct { 
	// Results
	Results *Domainentitylistingqueryresult `json:"results,omitempty"`


	// FacetInfo
	FacetInfo *Queryfacetinfo `json:"facetInfo,omitempty"`

}

func (o *Queryresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryresults
	
	return json.Marshal(&struct { 
		Results *Domainentitylistingqueryresult `json:"results,omitempty"`
		
		FacetInfo *Queryfacetinfo `json:"facetInfo,omitempty"`
		*Alias
	}{ 
		Results: o.Results,
		
		FacetInfo: o.FacetInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryresults) UnmarshalJSON(b []byte) error {
	var QueryresultsMap map[string]interface{}
	err := json.Unmarshal(b, &QueryresultsMap)
	if err != nil {
		return err
	}
	
	if Results, ok := QueryresultsMap["results"].(map[string]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if FacetInfo, ok := QueryresultsMap["facetInfo"].(map[string]interface{}); ok {
		FacetInfoString, _ := json.Marshal(FacetInfo)
		json.Unmarshal(FacetInfoString, &o.FacetInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queryresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
