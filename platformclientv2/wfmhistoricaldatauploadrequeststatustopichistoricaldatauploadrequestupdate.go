package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate
type Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate struct { 
	// RequestId
	RequestId *string `json:"requestId,omitempty"`


	// DateImportStarted
	DateImportStarted *time.Time `json:"dateImportStarted,omitempty"`


	// DateImportEnded
	DateImportEnded *time.Time `json:"dateImportEnded,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// VarError
	VarError *string `json:"error,omitempty"`


	// Active
	Active *bool `json:"active,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate
	
	DateImportStarted := new(string)
	if o.DateImportStarted != nil {
		
		*DateImportStarted = timeutil.Strftime(o.DateImportStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportStarted = nil
	}
	
	DateImportEnded := new(string)
	if o.DateImportEnded != nil {
		
		*DateImportEnded = timeutil.Strftime(o.DateImportEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImportEnded = nil
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
		
		DateImportStarted *string `json:"dateImportStarted,omitempty"`
		
		DateImportEnded *string `json:"dateImportEnded,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		VarError *string `json:"error,omitempty"`
		
		Active *bool `json:"active,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		RequestId: o.RequestId,
		
		DateImportStarted: DateImportStarted,
		
		DateImportEnded: DateImportEnded,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Status: o.Status,
		
		VarError: o.VarError,
		
		Active: o.Active,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) UnmarshalJSON(b []byte) error {
	var WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap map[string]interface{}
	err := json.Unmarshal(b, &WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap)
	if err != nil {
		return err
	}
	
	if RequestId, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["requestId"].(string); ok {
		o.RequestId = &RequestId
	}
    
	if dateImportStartedString, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["dateImportStarted"].(string); ok {
		DateImportStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportStartedString)
		o.DateImportStarted = &DateImportStarted
	}
	
	if dateImportEndedString, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["dateImportEnded"].(string); ok {
		DateImportEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportEndedString)
		o.DateImportEnded = &DateImportEnded
	}
	
	if dateCreatedString, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Status, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarError, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["error"].(string); ok {
		o.VarError = &VarError
	}
    
	if Active, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["active"].(bool); ok {
		o.Active = &Active
	}
    
	if VarType, ok := WfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdateMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
