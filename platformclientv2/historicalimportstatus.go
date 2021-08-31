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

func (o *Historicalimportstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Historicalimportstatus
	
	DateImportEnded := new(string)
	if o.DateImportEnded != nil {
		
		*DateImportEnded = timeutil.Strftime(o.DateImportEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportEnded = nil
	}
	
	DateImportStarted := new(string)
	if o.DateImportStarted != nil {
		
		*DateImportStarted = timeutil.Strftime(o.DateImportStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportStarted = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		RequestId: o.RequestId,
		
		DateImportEnded: DateImportEnded,
		
		DateImportStarted: DateImportStarted,
		
		Status: o.Status,
		
		VarError: o.VarError,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Active: o.Active,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Historicalimportstatus) UnmarshalJSON(b []byte) error {
	var HistoricalimportstatusMap map[string]interface{}
	err := json.Unmarshal(b, &HistoricalimportstatusMap)
	if err != nil {
		return err
	}
	
	if RequestId, ok := HistoricalimportstatusMap["requestId"].(string); ok {
		o.RequestId = &RequestId
	}
	
	if dateImportEndedString, ok := HistoricalimportstatusMap["dateImportEnded"].(string); ok {
		DateImportEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportEndedString)
		o.DateImportEnded = &DateImportEnded
	}
	
	if dateImportStartedString, ok := HistoricalimportstatusMap["dateImportStarted"].(string); ok {
		DateImportStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportStartedString)
		o.DateImportStarted = &DateImportStarted
	}
	
	if Status, ok := HistoricalimportstatusMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if VarError, ok := HistoricalimportstatusMap["error"].(string); ok {
		o.VarError = &VarError
	}
	
	if dateCreatedString, ok := HistoricalimportstatusMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := HistoricalimportstatusMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Active, ok := HistoricalimportstatusMap["active"].(bool); ok {
		o.Active = &Active
	}
	
	if VarType, ok := HistoricalimportstatusMap["type"].(string); ok {
		o.VarType = &VarType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Historicalimportstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
