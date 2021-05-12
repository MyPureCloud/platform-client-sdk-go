package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Validationservicerequest
type Validationservicerequest struct { 
	// DateImportEnded - The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateImportEnded *time.Time `json:"dateImportEnded,omitempty"`


	// FileUrl - Path to the file in the storage including the file name
	FileUrl *string `json:"fileUrl,omitempty"`

}

// String returns a JSON representation of the model
func (o *Validationservicerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
