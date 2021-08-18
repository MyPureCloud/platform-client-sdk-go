package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditfilter
type Auditfilter struct { 
	// Name - The name of the field by which to filter.
	Name *string `json:"name,omitempty"`


	// VarType - The type of the filter, DATE or STRING.
	VarType *string `json:"type,omitempty"`


	// Operator - The operation that the filter performs.
	Operator *string `json:"operator,omitempty"`


	// Values - The values to make the filter comparison against.
	Values *[]string `json:"values,omitempty"`

}

func (u *Auditfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditfilter

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		VarType: u.VarType,
		
		Operator: u.Operator,
		
		Values: u.Values,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
