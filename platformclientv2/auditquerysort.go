package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditquerysort
type Auditquerysort struct { 
	// Name - Name of the property to sort.
	Name *string `json:"name,omitempty"`


	// SortOrder - Sort Order
	SortOrder *string `json:"sortOrder,omitempty"`

}

func (o *Auditquerysort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditquerysort
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		SortOrder: o.SortOrder,
		Alias:    (*Alias)(o),
	})
}

func (o *Auditquerysort) UnmarshalJSON(b []byte) error {
	var AuditquerysortMap map[string]interface{}
	err := json.Unmarshal(b, &AuditquerysortMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AuditquerysortMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SortOrder, ok := AuditquerysortMap["sortOrder"].(string); ok {
		o.SortOrder = &SortOrder
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Auditquerysort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
