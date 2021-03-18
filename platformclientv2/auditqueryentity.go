package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryentity
type Auditqueryentity struct { 
	// Name - Name of the Entity
	Name *string `json:"name,omitempty"`


	// Actions - List of Actions
	Actions *[]string `json:"actions,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditqueryentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
