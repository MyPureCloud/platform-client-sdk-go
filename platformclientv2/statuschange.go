package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Statuschange
type Statuschange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Statuschange) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Statuschange) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStatusChanged", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
	}{ 
		DateStatusChanged: DateStatusChanged,
		
		Status: o.Status,
		
		PreviousStatus: o.PreviousStatus,
		
		Namespace: o.Namespace,
		
		Message: o.Message,
		
		RejectReason: o.RejectReason,
		Alias:    (Alias)(o),
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
