package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate
type Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) SetField(field string, fieldValue interface{}) {
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

func (o Wfmhistoricaldatauploadrequeststatustopichistoricaldatauploadrequestupdate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateImportStarted","DateImportEnded","DateCreated","DateModified", }
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
		Alias
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
		Alias:    (Alias)(o),
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
