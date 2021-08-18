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

func (u *Queryresults) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryresults

	

	return json.Marshal(&struct { 
		Results *Domainentitylistingqueryresult `json:"results,omitempty"`
		
		FacetInfo *Queryfacetinfo `json:"facetInfo,omitempty"`
		*Alias
	}{ 
		Results: u.Results,
		
		FacetInfo: u.FacetInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queryresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
