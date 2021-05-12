package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Queryresults) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
