package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryservicemapping
type Auditqueryservicemapping struct { 
	// Services - List of Services
	Services *[]Auditqueryservice `json:"services,omitempty"`

}

func (o *Auditqueryservicemapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryservicemapping
	
	return json.Marshal(&struct { 
		Services *[]Auditqueryservice `json:"services,omitempty"`
		*Alias
	}{ 
		Services: o.Services,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditqueryservicemapping) UnmarshalJSON(b []byte) error {
	var AuditqueryservicemappingMap map[string]interface{}
	err := json.Unmarshal(b, &AuditqueryservicemappingMap)
	if err != nil {
		return err
	}
	
	if Services, ok := AuditqueryservicemappingMap["services"].([]interface{}); ok {
		ServicesString, _ := json.Marshal(Services)
		json.Unmarshal(ServicesString, &o.Services)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Auditqueryservicemapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
