package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Exporturi
type Exporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ExportTimestamp *time.Time `json:"exportTimestamp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Exporturi) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
