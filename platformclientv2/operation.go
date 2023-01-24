package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Operation
type Operation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Complete
	Complete *bool `json:"complete,omitempty"`

	// User
	User *User `json:"user,omitempty"`

	// Client
	Client *Domainentityref `json:"client,omitempty"`

	// ErrorMessage
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// ErrorCode
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorDetails
	ErrorDetails *[]Detail `json:"errorDetails,omitempty"`

	// ErrorMessageParams
	ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`

	// ActionName - Action name
	ActionName *string `json:"actionName,omitempty"`

	// ActionStatus - Action status
	ActionStatus *string `json:"actionStatus,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Operation) SetField(field string, fieldValue interface{}) {
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

func (o Operation) MarshalJSON() ([]byte, error) {
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
	type Alias Operation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Complete *bool `json:"complete,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Client *Domainentityref `json:"client,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorDetails *[]Detail `json:"errorDetails,omitempty"`
		
		ErrorMessageParams *map[string]string `json:"errorMessageParams,omitempty"`
		
		ActionName *string `json:"actionName,omitempty"`
		
		ActionStatus *string `json:"actionStatus,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Complete: o.Complete,
		
		User: o.User,
		
		Client: o.Client,
		
		ErrorMessage: o.ErrorMessage,
		
		ErrorCode: o.ErrorCode,
		
		ErrorDetails: o.ErrorDetails,
		
		ErrorMessageParams: o.ErrorMessageParams,
		
		ActionName: o.ActionName,
		
		ActionStatus: o.ActionStatus,
		Alias:    (Alias)(o),
	})
}

func (o *Operation) UnmarshalJSON(b []byte) error {
	var OperationMap map[string]interface{}
	err := json.Unmarshal(b, &OperationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OperationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Complete, ok := OperationMap["complete"].(bool); ok {
		o.Complete = &Complete
	}
    
	if User, ok := OperationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Client, ok := OperationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if ErrorMessage, ok := OperationMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if ErrorCode, ok := OperationMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorDetails, ok := OperationMap["errorDetails"].([]interface{}); ok {
		ErrorDetailsString, _ := json.Marshal(ErrorDetails)
		json.Unmarshal(ErrorDetailsString, &o.ErrorDetails)
	}
	
	if ErrorMessageParams, ok := OperationMap["errorMessageParams"].(map[string]interface{}); ok {
		ErrorMessageParamsString, _ := json.Marshal(ErrorMessageParams)
		json.Unmarshal(ErrorMessageParamsString, &o.ErrorMessageParams)
	}
	
	if ActionName, ok := OperationMap["actionName"].(string); ok {
		o.ActionName = &ActionName
	}
    
	if ActionStatus, ok := OperationMap["actionStatus"].(string); ok {
		o.ActionStatus = &ActionStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Operation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
