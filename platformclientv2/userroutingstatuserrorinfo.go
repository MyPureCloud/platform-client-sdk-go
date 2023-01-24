package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatuserrorinfo - Error information that the Public API will receive in a response body. This allows backend services to pass an error message to consumers of the Public API.
type Userroutingstatuserrorinfo struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ErrorCode - A code unique to this error. Typically prefixed with the service that originated the error. For example CONFIG_USER_NOT_FOUND
	ErrorCode *string `json:"errorCode,omitempty"`

	// Status - The HTTP status code for this message. If left blank the status code from the HTTP response is used.
	Status *int `json:"status,omitempty"`

	// CorrelationId - The correlation Id or context Id for this message. If left blank the Public API will look at the HTTP response header 'ININ-Correlation-Id' instead.
	CorrelationId *string `json:"correlationId,omitempty"`

	// UserMessage - A customer friendly message. This should be a complete sentence, use proper grammar and only include information useful to a customer. This is not a dev message and should not include things like Org Id
	UserMessage *string `json:"userMessage,omitempty"`

	// UserParamsMessage - This is the same as userMessage except it uses template fields for variable replacement. For instance: 'User {username} was not found'
	UserParamsMessage *string `json:"userParamsMessage,omitempty"`

	// UserParams - Used in conjunction with userParamsMessage. These are the template parameters. For instance: UserParam.key = 'username', UserParam.value = 'chuck.pulfer'
	UserParams *[]Userroutingstatususerparam `json:"userParams,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userroutingstatuserrorinfo) SetField(field string, fieldValue interface{}) {
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

func (o Userroutingstatuserrorinfo) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Userroutingstatuserrorinfo
	
	return json.Marshal(&struct { 
		ErrorCode *string `json:"errorCode,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		UserMessage *string `json:"userMessage,omitempty"`
		
		UserParamsMessage *string `json:"userParamsMessage,omitempty"`
		
		UserParams *[]Userroutingstatususerparam `json:"userParams,omitempty"`
		Alias
	}{ 
		ErrorCode: o.ErrorCode,
		
		Status: o.Status,
		
		CorrelationId: o.CorrelationId,
		
		UserMessage: o.UserMessage,
		
		UserParamsMessage: o.UserParamsMessage,
		
		UserParams: o.UserParams,
		Alias:    (Alias)(o),
	})
}

func (o *Userroutingstatuserrorinfo) UnmarshalJSON(b []byte) error {
	var UserroutingstatuserrorinfoMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatuserrorinfoMap)
	if err != nil {
		return err
	}
	
	if ErrorCode, ok := UserroutingstatuserrorinfoMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if Status, ok := UserroutingstatuserrorinfoMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if CorrelationId, ok := UserroutingstatuserrorinfoMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if UserMessage, ok := UserroutingstatuserrorinfoMap["userMessage"].(string); ok {
		o.UserMessage = &UserMessage
	}
    
	if UserParamsMessage, ok := UserroutingstatuserrorinfoMap["userParamsMessage"].(string); ok {
		o.UserParamsMessage = &UserParamsMessage
	}
    
	if UserParams, ok := UserroutingstatuserrorinfoMap["userParams"].([]interface{}); ok {
		UserParamsString, _ := json.Marshal(UserParams)
		json.Unmarshal(UserParamsString, &o.UserParams)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatuserrorinfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
