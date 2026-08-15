package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectjobmessagedetail
type Architectjobmessagedetail struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The kind of information carried by this entry, which determines which of the other properties are set.
	VarType *string `json:"type,omitempty"`

	// Url - The URL of the request.
	Url *string `json:"url,omitempty"`

	// Method - The HTTP method of the request.
	Method *string `json:"method,omitempty"`

	// RequestBody - The body of the request, reported as sent and without redaction. Omitted when the request had no body, so it is absent for ordinary GET lookups and present for calls such as POST searches. Truncated to 4096 characters with a `...<truncated N chars>` suffix when longer.
	RequestBody *string `json:"requestBody,omitempty"`

	// StatusCode - The HTTP status code of the response. Set only when a response was received, and never alongside errorCode.
	StatusCode *int `json:"statusCode,omitempty"`

	// StatusMessage - The HTTP status message of the response. Set only when a response was received, and never alongside errorMessage.
	StatusMessage *string `json:"statusMessage,omitempty"`

	// CorrelationId - The Genesys Cloud correlation id of the response, to quote when escalating to Genesys Cloud support. Set only when a response was received.
	CorrelationId *string `json:"correlationId,omitempty"`

	// ResponseBody - The body of the response, reported as received and without redaction. Because entries are captured for requests that succeeded as well, this can carry data returned by a lookup that was unrelated to the failure. Omitted when the response had no body. Truncated to 4096 characters with a `...<truncated N chars>` suffix when longer.
	ResponseBody *string `json:"responseBody,omitempty"`

	// ErrorCode - The transport error code, such as ECONNRESET. Set only when the request failed before any HTTP response was received, and never alongside statusCode.
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessage - The transport error message. Set only when the request failed before any HTTP response was received, and never alongside statusMessage.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Architectjobmessagedetail) SetField(field string, fieldValue interface{}) {
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

func (o Architectjobmessagedetail) MarshalJSON() ([]byte, error) {
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
	type Alias Architectjobmessagedetail
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Method *string `json:"method,omitempty"`
		
		RequestBody *string `json:"requestBody,omitempty"`
		
		StatusCode *int `json:"statusCode,omitempty"`
		
		StatusMessage *string `json:"statusMessage,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		ResponseBody *string `json:"responseBody,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Url: o.Url,
		
		Method: o.Method,
		
		RequestBody: o.RequestBody,
		
		StatusCode: o.StatusCode,
		
		StatusMessage: o.StatusMessage,
		
		CorrelationId: o.CorrelationId,
		
		ResponseBody: o.ResponseBody,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		Alias:    (Alias)(o),
	})
}

func (o *Architectjobmessagedetail) UnmarshalJSON(b []byte) error {
	var ArchitectjobmessagedetailMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectjobmessagedetailMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ArchitectjobmessagedetailMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Url, ok := ArchitectjobmessagedetailMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Method, ok := ArchitectjobmessagedetailMap["method"].(string); ok {
		o.Method = &Method
	}
    
	if RequestBody, ok := ArchitectjobmessagedetailMap["requestBody"].(string); ok {
		o.RequestBody = &RequestBody
	}
    
	if StatusCode, ok := ArchitectjobmessagedetailMap["statusCode"].(float64); ok {
		StatusCodeInt := int(StatusCode)
		o.StatusCode = &StatusCodeInt
	}
	
	if StatusMessage, ok := ArchitectjobmessagedetailMap["statusMessage"].(string); ok {
		o.StatusMessage = &StatusMessage
	}
    
	if CorrelationId, ok := ArchitectjobmessagedetailMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if ResponseBody, ok := ArchitectjobmessagedetailMap["responseBody"].(string); ok {
		o.ResponseBody = &ResponseBody
	}
    
	if ErrorCode, ok := ArchitectjobmessagedetailMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := ArchitectjobmessagedetailMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Architectjobmessagedetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
