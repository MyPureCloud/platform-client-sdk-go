package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkresponseresultexternalcontactentity
type Bulkresponseresultexternalcontactentity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Success
	Success *bool `json:"success,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// VarError
	VarError *Bulkerrorentity `json:"error,omitempty"`

}

// String returns a JSON representation of the model
func (o *Bulkresponseresultexternalcontactentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
