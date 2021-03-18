package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Auditquerysort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
