package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facet
type Facet struct { 
	// Name - The name of the field on which to facet.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the facet, DATE or STRING.
	VarType *string `json:"type,omitempty"`

}

func (o *Facet) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facet
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Facet) UnmarshalJSON(b []byte) error {
	var FacetMap map[string]interface{}
	err := json.Unmarshal(b, &FacetMap)
	if err != nil {
		return err
	}
	
	if Name, ok := FacetMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := FacetMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facet) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
