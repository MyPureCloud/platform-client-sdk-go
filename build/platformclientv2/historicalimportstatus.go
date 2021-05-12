package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalimportstatus
type Historicalimportstatus struct { 
	// RequestId - Request id of the historical import in the organization
	RequestId *string `json:"requestId,omitempty"`


	// DateImportEnded - The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateImportEnded *time.Time `json:"dateImportEnded,omitempty"`


	// DateImportStarted - The first day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateImportStarted *time.Time `json:"dateImportStarted,omitempty"`


	// Status - Status of the historical import in the organization.
	Status *string `json:"status,omitempty"`


	// VarError - Error occured if the status of the import is failed
	VarError *string `json:"error,omitempty"`


	// DateCreated - Date in which the historical import is initiated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date in which the historical import is modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Active - Whether this historical import is active or not
	Active *bool `json:"active,omitempty"`


	// VarType - Whether this historical import is of type csv or json
	VarType *string `json:"type,omitempty"`

}

// String returns a JSON representation of the model
func (o *Historicalimportstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
