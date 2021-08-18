package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Historicalimportstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalimportstatus

	
	DateImportEnded := new(string)
	if u.DateImportEnded != nil {
		
		*DateImportEnded = timeutil.Strftime(u.DateImportEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportEnded = nil
	}
	
	DateImportStarted := new(string)
	if u.DateImportStarted != nil {
		
		*DateImportStarted = timeutil.Strftime(u.DateImportStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportStarted = nil
	}
	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		RequestId *string `json:"requestId,omitempty"`
		
		DateImportEnded *string `json:"dateImportEnded,omitempty"`
		
		DateImportStarted *string `json:"dateImportStarted,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *string `json:"error,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		RequestId: u.RequestId,
		
		DateImportEnded: DateImportEnded,
		
		DateImportStarted: DateImportStarted,
		
		Status: u.Status,
		
		VarError: u.VarError,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Active: u.Active,
		
		VarType: u.VarType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Historicalimportstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
