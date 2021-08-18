package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerauditrequest
type Dialerauditrequest struct { 
	// QueryPhrase - The word or words to search for.
	QueryPhrase *string `json:"queryPhrase,omitempty"`


	// QueryFields - The fields in which to search for the queryPhrase.
	QueryFields *[]string `json:"queryFields,omitempty"`


	// Facets - The fields to facet on.
	Facets *[]Auditfacet `json:"facets,omitempty"`


	// Filters - The fields to filter on.
	Filters *[]Auditfilter `json:"filters,omitempty"`

}

func (u *Dialerauditrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerauditrequest

	

	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		QueryFields *[]string `json:"queryFields,omitempty"`
		
		Facets *[]Auditfacet `json:"facets,omitempty"`
		
		Filters *[]Auditfilter `json:"filters,omitempty"`
		*Alias
	}{ 
		QueryPhrase: u.QueryPhrase,
		
		QueryFields: u.QueryFields,
		
		Facets: u.Facets,
		
		Filters: u.Filters,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerauditrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
