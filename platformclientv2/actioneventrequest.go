package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actioneventrequest
type Actioneventrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// SessionId - UUID of the customer session for this action.
	SessionId *string `json:"sessionId,omitempty"`

	// ActionId - UUID for the action, as returned by the Ping endpoint when the action was qualified.
	ActionId *string `json:"actionId,omitempty"`

	// ActionState - State the action is transitioning to.
	ActionState *string `json:"actionState,omitempty"`

	// ErrorCode - Client defined error code (when state transitions to errored)
	ErrorCode *string `json:"errorCode,omitempty"`

	// ErrorMessage - Message of the error returned when the action fails (when state transitions to errored)
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actioneventrequest) SetField(field string, fieldValue interface{}) {
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

func (o Actioneventrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Actioneventrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		ActionId *string `json:"actionId,omitempty"`
		
		ActionState *string `json:"actionState,omitempty"`
		
		ErrorCode *string `json:"errorCode,omitempty"`
		
		ErrorMessage *string `json:"errorMessage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		SessionId: o.SessionId,
		
		ActionId: o.ActionId,
		
		ActionState: o.ActionState,
		
		ErrorCode: o.ErrorCode,
		
		ErrorMessage: o.ErrorMessage,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Actioneventrequest) UnmarshalJSON(b []byte) error {
	var ActioneventrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ActioneventrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActioneventrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SessionId, ok := ActioneventrequestMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if ActionId, ok := ActioneventrequestMap["actionId"].(string); ok {
		o.ActionId = &ActionId
	}
    
	if ActionState, ok := ActioneventrequestMap["actionState"].(string); ok {
		o.ActionState = &ActionState
	}
    
	if ErrorCode, ok := ActioneventrequestMap["errorCode"].(string); ok {
		o.ErrorCode = &ErrorCode
	}
    
	if ErrorMessage, ok := ActioneventrequestMap["errorMessage"].(string); ok {
		o.ErrorMessage = &ErrorMessage
	}
    
	if SelfUri, ok := ActioneventrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Actioneventrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
