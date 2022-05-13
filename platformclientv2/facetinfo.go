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

func (o *Facetinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetinfo
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Entries *[]Entry `json:"entries,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Entries: o.Entries,
		Alias:    (*Alias)(o),
	})
}

func (o *Facetinfo) UnmarshalJSON(b []byte) error {
	var FacetinfoMap map[string]interface{}
	err := json.Unmarshal(b, &FacetinfoMap)
	if err != nil {
		return err
	}
	
	if Name, ok := FacetinfoMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Entries, ok := FacetinfoMap["entries"].([]interface{}); ok {
		EntriesString, _ := json.Marshal(Entries)
		json.Unmarshal(EntriesString, &o.Entries)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
