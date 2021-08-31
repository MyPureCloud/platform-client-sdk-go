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

func (o *Queryfacetinfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queryfacetinfo
	
	return json.Marshal(&struct { 
		Attributes *[]Facetkeyattribute `json:"attributes,omitempty"`
		
		Facets *[]Facetentry `json:"facets,omitempty"`
		*Alias
	}{ 
		Attributes: o.Attributes,
		
		Facets: o.Facets,
		Alias:    (*Alias)(o),
	})
}

func (o *Queryfacetinfo) UnmarshalJSON(b []byte) error {
	var QueryfacetinfoMap map[string]interface{}
	err := json.Unmarshal(b, &QueryfacetinfoMap)
	if err != nil {
		return err
	}
	
	if Attributes, ok := QueryfacetinfoMap["attributes"].([]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Facets, ok := QueryfacetinfoMap["facets"].([]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queryfacetinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
