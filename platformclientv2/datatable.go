package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Datatable - Contains a metadata representation for a JSON schema stored in DataTables along with an optional field for the schema itself
type Datatable struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Writabledivision `json:"division,omitempty"`


	// Description - The description from the JSON schema (equates to the Description field on the JSON schema.)
	Description *string `json:"description,omitempty"`


	// Schema - the schema as stored in the system.
	Schema *Jsonschemadocument `json:"schema,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Datatable) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
