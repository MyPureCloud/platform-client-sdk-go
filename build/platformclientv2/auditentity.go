package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Auditentity
type Auditentity struct { 
	// VarType - The type of the entity the action of this AuditEntity targeted.
	VarType *string `json:"type,omitempty"`


	// Id - The id of the entity the action of this AuditEntity targeted.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity the action of this AuditEntity targeted.
	Name *string `json:"name,omitempty"`


	// SelfUri - The selfUri for this entity.
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Auditentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
