package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Parameter
type Parameter struct { 
	// Name
	Name *string `json:"name,omitempty"`


	// ParameterType
	ParameterType *string `json:"parameterType,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// Required
	Required *bool `json:"required,omitempty"`

}

func (u *Parameter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Parameter

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ParameterType *string `json:"parameterType,omitempty"`
		
		Domain *string `json:"domain,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		ParameterType: u.ParameterType,
		
		Domain: u.Domain,
		
		Required: u.Required,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Parameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
