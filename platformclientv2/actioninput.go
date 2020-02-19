package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Actioninput) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
