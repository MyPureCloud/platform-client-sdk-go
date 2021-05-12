package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Facetinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
