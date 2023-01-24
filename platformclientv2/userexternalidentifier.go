package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userexternalidentifier - Defines a link between an External Identifier and Authority pair to a Entity Type and Entity Identifier pair. Represents the two way, one to one mapping of an External Authority or System of Record's identifier to a PureCloud entity. e.g. (ExternalId='05001',Authority='XyzCRM') to (entityType=user,entityId='8eb03b33-3acb-4bc1-a244-50b9b9f19495')
type Userexternalidentifier struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AuthorityName - Authority or System of Record which owns the External Identifier
	AuthorityName *string `json:"authorityName,omitempty"`

	// ExternalKey - External Key
	ExternalKey *string `json:"externalKey,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userexternalidentifier) SetField(field string, fieldValue interface{}) {
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

func (o Userexternalidentifier) MarshalJSON() ([]byte, error) {
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
	type Alias Userexternalidentifier
	
	return json.Marshal(&struct { 
		AuthorityName *string `json:"authorityName,omitempty"`
		
		ExternalKey *string `json:"externalKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		AuthorityName: o.AuthorityName,
		
		ExternalKey: o.ExternalKey,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Userexternalidentifier) UnmarshalJSON(b []byte) error {
	var UserexternalidentifierMap map[string]interface{}
	err := json.Unmarshal(b, &UserexternalidentifierMap)
	if err != nil {
		return err
	}
	
	if AuthorityName, ok := UserexternalidentifierMap["authorityName"].(string); ok {
		o.AuthorityName = &AuthorityName
	}
    
	if ExternalKey, ok := UserexternalidentifierMap["externalKey"].(string); ok {
		o.ExternalKey = &ExternalKey
	}
    
	if SelfUri, ok := UserexternalidentifierMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userexternalidentifier) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
