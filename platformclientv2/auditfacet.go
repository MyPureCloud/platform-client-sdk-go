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

func (u *Auditfacet) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditfacet

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditfacet) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
