package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Emailconfig
type Emailconfig struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EmailColumns - The contact list columns specifying the email address(es) of the contact.
	EmailColumns *[]string `json:"emailColumns,omitempty"`

	// ContentTemplate - The content template used to formulate the email to send to the contact.
	ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`

	// FromAddress - The email address that will be used as the sender of the email.
	FromAddress *Fromemailaddress `json:"fromAddress,omitempty"`

	// ReplyToAddress - The email address from which any reply will be sent.
	ReplyToAddress *Replytoemailaddress `json:"replyToAddress,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Emailconfig) SetField(field string, fieldValue interface{}) {
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

func (o Emailconfig) MarshalJSON() ([]byte, error) {
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
	type Alias Emailconfig
	
	return json.Marshal(&struct { 
		EmailColumns *[]string `json:"emailColumns,omitempty"`
		
		ContentTemplate *Domainentityref `json:"contentTemplate,omitempty"`
		
		FromAddress *Fromemailaddress `json:"fromAddress,omitempty"`
		
		ReplyToAddress *Replytoemailaddress `json:"replyToAddress,omitempty"`
		Alias
	}{ 
		EmailColumns: o.EmailColumns,
		
		ContentTemplate: o.ContentTemplate,
		
		FromAddress: o.FromAddress,
		
		ReplyToAddress: o.ReplyToAddress,
		Alias:    (Alias)(o),
	})
}

func (o *Emailconfig) UnmarshalJSON(b []byte) error {
	var EmailconfigMap map[string]interface{}
	err := json.Unmarshal(b, &EmailconfigMap)
	if err != nil {
		return err
	}
	
	if EmailColumns, ok := EmailconfigMap["emailColumns"].([]interface{}); ok {
		EmailColumnsString, _ := json.Marshal(EmailColumns)
		json.Unmarshal(EmailColumnsString, &o.EmailColumns)
	}
	
	if ContentTemplate, ok := EmailconfigMap["contentTemplate"].(map[string]interface{}); ok {
		ContentTemplateString, _ := json.Marshal(ContentTemplate)
		json.Unmarshal(ContentTemplateString, &o.ContentTemplate)
	}
	
	if FromAddress, ok := EmailconfigMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if ReplyToAddress, ok := EmailconfigMap["replyToAddress"].(map[string]interface{}); ok {
		ReplyToAddressString, _ := json.Marshal(ReplyToAddress)
		json.Unmarshal(ReplyToAddressString, &o.ReplyToAddress)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Emailconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
