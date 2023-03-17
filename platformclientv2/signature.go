package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Signature
type Signature struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Enabled - A toggle to enable the signature on email send.
	Enabled *bool `json:"enabled,omitempty"`

	// CannedResponseId - The identifier referring to an email signature canned response.
	CannedResponseId *string `json:"cannedResponseId,omitempty"`

	// AlwaysIncluded - A toggle that defines if a signature is always included or only set on the first email in an email chain.
	AlwaysIncluded *bool `json:"alwaysIncluded,omitempty"`

	// InclusionType - The configuration to indicate when the signature of a conversation has to be included
	InclusionType *string `json:"inclusionType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Signature) SetField(field string, fieldValue interface{}) {
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

func (o Signature) MarshalJSON() ([]byte, error) {
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
	type Alias Signature
	
	return json.Marshal(&struct { 
		Enabled *bool `json:"enabled,omitempty"`
		
		CannedResponseId *string `json:"cannedResponseId,omitempty"`
		
		AlwaysIncluded *bool `json:"alwaysIncluded,omitempty"`
		
		InclusionType *string `json:"inclusionType,omitempty"`
		Alias
	}{ 
		Enabled: o.Enabled,
		
		CannedResponseId: o.CannedResponseId,
		
		AlwaysIncluded: o.AlwaysIncluded,
		
		InclusionType: o.InclusionType,
		Alias:    (Alias)(o),
	})
}

func (o *Signature) UnmarshalJSON(b []byte) error {
	var SignatureMap map[string]interface{}
	err := json.Unmarshal(b, &SignatureMap)
	if err != nil {
		return err
	}
	
	if Enabled, ok := SignatureMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if CannedResponseId, ok := SignatureMap["cannedResponseId"].(string); ok {
		o.CannedResponseId = &CannedResponseId
	}
    
	if AlwaysIncluded, ok := SignatureMap["alwaysIncluded"].(bool); ok {
		o.AlwaysIncluded = &AlwaysIncluded
	}
    
	if InclusionType, ok := SignatureMap["inclusionType"].(string); ok {
		o.InclusionType = &InclusionType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Signature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
