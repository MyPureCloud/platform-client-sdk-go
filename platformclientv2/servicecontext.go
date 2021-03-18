package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Servicecontext
type Servicecontext struct { 
	// Name - Unused field for the purpose of ensuring a Swagger definition is created for a class with only @JsonIgnore members.
	Name *string `json:"name,omitempty"`

}

// String returns a JSON representation of the model
func (o *Servicecontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
