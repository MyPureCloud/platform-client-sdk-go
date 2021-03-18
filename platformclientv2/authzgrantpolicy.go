package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Authzgrantpolicy
type Authzgrantpolicy struct { 
	// Actions
	Actions *[]string `json:"actions,omitempty"`


	// Condition
	Condition *string `json:"condition,omitempty"`


	// Domain
	Domain *string `json:"domain,omitempty"`


	// EntityName
	EntityName *string `json:"entityName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Authzgrantpolicy) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
