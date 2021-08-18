package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queryfacetinfo
type Queryfacetinfo struct { 
	// Attributes
	Attributes *[]Facetkeyattribute `json:"attributes,omitempty"`


	// Facets
	Facets *[]Facetentry `json:"facets,omitempty"`

}

func (u *Queryfacetinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryfacetinfo

	

	return json.Marshal(&struct { 
		Attributes *[]Facetkeyattribute `json:"attributes,omitempty"`
		
		Facets *[]Facetentry `json:"facets,omitempty"`
		*Alias
	}{ 
		Attributes: u.Attributes,
		
		Facets: u.Facets,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queryfacetinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
