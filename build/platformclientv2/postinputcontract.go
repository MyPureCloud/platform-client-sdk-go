package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Postinputcontract - The schemas defining all of the expected requests/inputs.
type Postinputcontract struct { 
	// InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path.
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`

}

// String returns a JSON representation of the model
func (o *Postinputcontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
