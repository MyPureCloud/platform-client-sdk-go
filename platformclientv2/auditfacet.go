package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditfacet
type Auditfacet struct { 
	// Name - The name of the field on which to facet.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the facet, DATE or STRING.
	VarType *string `json:"type,omitempty"`

}

func (o *Auditfacet) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditfacet
	
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

func (o *Auditfacet) UnmarshalJSON(b []byte) error {
	var AuditfacetMap map[string]interface{}
	err := json.Unmarshal(b, &AuditfacetMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AuditfacetMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if VarType, ok := AuditfacetMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditfacet) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
