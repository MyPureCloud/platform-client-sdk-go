package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Historicalimportstatus
type Historicalimportstatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Historicalimportstatus) SetField(field string, fieldValue interface{}) {
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

func (o Historicalimportstatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateImportEnded","DateImportStarted","DateCreated","DateModified", }
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
		Alias
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
		Alias:    (Alias)(o),
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
