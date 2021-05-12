package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Postoutputcontract - The schemas defining all of the expected responses/outputs.
type Postoutputcontract struct { 
	// SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller.
	SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`

}

// String returns a JSON representation of the model
func (o *Postoutputcontract) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
