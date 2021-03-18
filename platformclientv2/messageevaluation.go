package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Messageevaluation
type Messageevaluation struct { 
	// ContactColumn
	ContactColumn *string `json:"contactColumn,omitempty"`


	// ContactAddress
	ContactAddress *string `json:"contactAddress,omitempty"`


	// WrapupCodeId
	WrapupCodeId *string `json:"wrapupCodeId,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Messageevaluation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
