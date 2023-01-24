package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Policyerrormessage
type Policyerrormessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StatusCode
	StatusCode *int `json:"statusCode,omitempty"`

	// UserMessage
	UserMessage *interface{} `json:"userMessage,omitempty"`

	// UserParamsMessage
	UserParamsMessage *string `json:"userParamsMessage,omitempty"`

	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`

	// UserParams
	UserParams *[]Userparam `json:"userParams,omitempty"`

	// InsertDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	InsertDate *time.Time `json:"insertDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Policyerrormessage) SetField(field string, fieldValue interface{}) {
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

func (o Policyerrormessage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "InsertDate", }
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
	type Alias Policyerrormessage
	
	InsertDate := new(string)
	if o.InsertDate != nil {
		
		*InsertDate = timeutil.Strftime(o.InsertDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		InsertDate = nil
	}
	
	return json.Marshal(&struct { 
		StatusCode *int `json:"statusCode,omitempty"`
		
		UserMessage *interface{} `json:"userMessage,omitempty"`
		
		UserParamsMessage *string `json:"userParamsMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		UserParams *[]Userparam `json:"userParams,omitempty"`
		
		InsertDate *string `json:"insertDate,omitempty"`
		Alias
	}{ 
		StatusCode: o.StatusCode,
		
		UserMessage: o.UserMessage,
		
		UserParamsMessage: o.UserParamsMessage,
		
		ErrorCode: o.ErrorCode,
		
		CorrelationId: o.CorrelationId,
		
		UserParams: o.UserParams,
		
		InsertDate: InsertDate,
		Alias:    (Alias)(o),
	})
}

func (o *Policyerrormessage) UnmarshalJSON(b []byte) error {
	var PolicyerrormessageMap map[string]interface{}
	err := json.Unmarshal(b, &PolicyerrormessageMap)
	if err != nil {
		return err
	}
	
	if StatusCode, ok := PolicyerrormessageMap["statusCode"].(float64); ok {
		StatusCodeInt := int(StatusCode)
		o.StatusCode = &StatusCodeInt
	}
	
	if UserMessage, ok := PolicyerrormessageMap["userMessage"].(map[string]interface{}); ok {
		UserMessageString, _ := json.Marshal(UserMessage)
		json.Unmarshal(UserMessageString, &o.UserMessage)
	}
	
	if UserParamsMessage, ok := PolicyerrormessageMap["userParamsMessage"].(string); ok {
		o.UserParamsMessage = &UserParamsMessage
	}
    
	if ErrorCode, ok := PolicyerrormessageMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if CorrelationId, ok := PolicyerrormessageMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if UserParams, ok := PolicyerrormessageMap["userParams"].([]interface{}); ok {
		UserParamsString, _ := json.Marshal(UserParams)
		json.Unmarshal(UserParamsString, &o.UserParams)
	}
	
	if insertDateString, ok := PolicyerrormessageMap["insertDate"].(string); ok {
		InsertDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", insertDateString)
		o.InsertDate = &InsertDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Policyerrormessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
