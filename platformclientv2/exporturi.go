package platformclientv2
import (
	"time"
	"encoding/json"
)

// Exporturi
type Exporturi struct { 
	// Uri
	Uri *string `json:"uri,omitempty"`


	// ExportTimestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	ExportTimestamp *time.Time `json:"exportTimestamp,omitempty"`

}

// String returns a JSON representation of the model
func (o *Exporturi) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
