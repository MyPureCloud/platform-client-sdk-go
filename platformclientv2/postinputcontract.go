package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Postinputcontract - The schemas defining all of the expected requests/inputs.
type Postinputcontract struct { 
	// InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path.
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`

}

func (u *Postinputcontract) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Postinputcontract

	

	return json.Marshal(&struct { 
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		*Alias
	}{ 
		InputSchema: u.InputSchema,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Postinputcontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
