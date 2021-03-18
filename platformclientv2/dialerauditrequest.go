package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Dialerauditrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
