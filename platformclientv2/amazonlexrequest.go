package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Amazonlexrequest
type Amazonlexrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RequestAttributes - AttributeName/AttributeValue pairs of User Defined Request Attributes to be sent to the amazon bot See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs
	RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`

	// SessionAttributes - AttributeName/AttributeValue pairs of Session Attributes to be sent to the amazon bot. See - https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs
	SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Amazonlexrequest) SetField(field string, fieldValue interface{}) {
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

func (o Amazonlexrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Amazonlexrequest
	
	return json.Marshal(&struct { 
		RequestAttributes *map[string]string `json:"requestAttributes,omitempty"`
		
		SessionAttributes *map[string]string `json:"sessionAttributes,omitempty"`
		Alias
	}{ 
		RequestAttributes: o.RequestAttributes,
		
		SessionAttributes: o.SessionAttributes,
		Alias:    (Alias)(o),
	})
}

func (o *Amazonlexrequest) UnmarshalJSON(b []byte) error {
	var AmazonlexrequestMap map[string]interface{}
	err := json.Unmarshal(b, &AmazonlexrequestMap)
	if err != nil {
		return err
	}
	
	if RequestAttributes, ok := AmazonlexrequestMap["requestAttributes"].(map[string]interface{}); ok {
		RequestAttributesString, _ := json.Marshal(RequestAttributes)
		json.Unmarshal(RequestAttributesString, &o.RequestAttributes)
	}
	
	if SessionAttributes, ok := AmazonlexrequestMap["sessionAttributes"].(map[string]interface{}); ok {
		SessionAttributesString, _ := json.Marshal(SessionAttributes)
		json.Unmarshal(SessionAttributesString, &o.SessionAttributes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Amazonlexrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
