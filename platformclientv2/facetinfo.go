package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetinfo
type Facetinfo struct { 
	// Name - The name of the field that was faceted on.
	Name *string `json:"name,omitempty"`


	// Entries - The entries resulting from this facet.
	Entries *[]Entry `json:"entries,omitempty"`

}

func (u *Facetinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetinfo

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Entries *[]Entry `json:"entries,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Entries: u.Entries,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Facetinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
