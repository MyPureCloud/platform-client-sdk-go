package platformclientv2
import (
	"encoding/json"
)

// Auditqueryresponse
type Auditqueryresponse struct { }

// String returns a JSON representation of the model
func (o *Auditqueryresponse) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
