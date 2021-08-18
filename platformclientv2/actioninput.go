package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actioninput - Input requirements of Action.
type Actioninput struct { 
	// InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. If the 'flatten' query parameter is omitted or false, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`


	// InputSchemaFlattened - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
	InputSchemaFlattened *Jsonschemadocument `json:"inputSchemaFlattened,omitempty"`


	// InputSchemaUri - The URI of the input schema
	InputSchemaUri *string `json:"inputSchemaUri,omitempty"`

}

func (u *Actioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actioninput

	

	return json.Marshal(&struct { 
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		
		InputSchemaFlattened *Jsonschemadocument `json:"inputSchemaFlattened,omitempty"`
		
		InputSchemaUri *string `json:"inputSchemaUri,omitempty"`
		*Alias
	}{ 
		InputSchema: u.InputSchema,
		
		InputSchemaFlattened: u.InputSchemaFlattened,
		
		InputSchemaUri: u.InputSchemaUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
