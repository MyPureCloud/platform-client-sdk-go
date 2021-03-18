package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Parameter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
