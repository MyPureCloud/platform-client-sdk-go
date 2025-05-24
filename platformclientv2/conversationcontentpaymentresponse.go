package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentpaymentresponse - Payment response object representing the status of a payment transaction.
type Conversationcontentpaymentresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OriginatingMessageId - Reference to the ID of the original payment request message this response is for.
	OriginatingMessageId *string `json:"originatingMessageId,omitempty"`

	// PaymentStatus - The status of the payment transaction.
	PaymentStatus *string `json:"paymentStatus,omitempty"`

	// FailureReason - The reason the payment request failed.
	FailureReason *string `json:"failureReason,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationcontentpaymentresponse) SetField(field string, fieldValue interface{}) {
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

func (o Conversationcontentpaymentresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationcontentpaymentresponse
	
	return json.Marshal(&struct { 
		OriginatingMessageId *string `json:"originatingMessageId,omitempty"`
		
		PaymentStatus *string `json:"paymentStatus,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		Alias
	}{ 
		OriginatingMessageId: o.OriginatingMessageId,
		
		PaymentStatus: o.PaymentStatus,
		
		FailureReason: o.FailureReason,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationcontentpaymentresponse) UnmarshalJSON(b []byte) error {
	var ConversationcontentpaymentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentpaymentresponseMap)
	if err != nil {
		return err
	}
	
	if OriginatingMessageId, ok := ConversationcontentpaymentresponseMap["originatingMessageId"].(string); ok {
		o.OriginatingMessageId = &OriginatingMessageId
	}
    
	if PaymentStatus, ok := ConversationcontentpaymentresponseMap["paymentStatus"].(string); ok {
		o.PaymentStatus = &PaymentStatus
	}
    
	if FailureReason, ok := ConversationcontentpaymentresponseMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentpaymentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
