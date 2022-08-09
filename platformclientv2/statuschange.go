package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Statuschange
type Statuschange struct { 
	// DateStatusChanged - The date of this status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStatusChanged *time.Time `json:"dateStatusChanged,omitempty"`


	// Status - The status the change request transitioned to
	Status *string `json:"status,omitempty"`


	// PreviousStatus - The status the change request transitioned from
	PreviousStatus *string `json:"previousStatus,omitempty"`


	// Namespace - The namespace for the status change
	Namespace *string `json:"namespace,omitempty"`


	// Message - A short message describing the status change
	Message *string `json:"message,omitempty"`


	// RejectReason - The reason for rejecting the limit override request
	RejectReason *string `json:"rejectReason,omitempty"`

}

func (o *Statuschange) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Statuschange
	
	DateStatusChanged := new(string)
	if o.DateStatusChanged != nil {
		
		*DateStatusChanged = timeutil.Strftime(o.DateStatusChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStatusChanged = nil
	}
	
	return json.Marshal(&struct { 
		DateStatusChanged *string `json:"dateStatusChanged,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PreviousStatus *string `json:"previousStatus,omitempty"`
		
		Namespace *string `json:"namespace,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		RejectReason *string `json:"rejectReason,omitempty"`
		*Alias
	}{ 
		DateStatusChanged: DateStatusChanged,
		
		Status: o.Status,
		
		PreviousStatus: o.PreviousStatus,
		
		Namespace: o.Namespace,
		
		Message: o.Message,
		
		RejectReason: o.RejectReason,
		Alias:    (*Alias)(o),
	})
}

func (o *Statuschange) UnmarshalJSON(b []byte) error {
	var StatuschangeMap map[string]interface{}
	err := json.Unmarshal(b, &StatuschangeMap)
	if err != nil {
		return err
	}
	
	if dateStatusChangedString, ok := StatuschangeMap["dateStatusChanged"].(string); ok {
		DateStatusChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStatusChangedString)
		o.DateStatusChanged = &DateStatusChanged
	}
	
	if Status, ok := StatuschangeMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if PreviousStatus, ok := StatuschangeMap["previousStatus"].(string); ok {
		o.PreviousStatus = &PreviousStatus
	}
    
	if Namespace, ok := StatuschangeMap["namespace"].(string); ok {
		o.Namespace = &Namespace
	}
    
	if Message, ok := StatuschangeMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if RejectReason, ok := StatuschangeMap["rejectReason"].(string); ok {
		o.RejectReason = &RejectReason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Statuschange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
