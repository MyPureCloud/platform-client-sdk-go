package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatuserref
type Chatuserref struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// Jid
	Jid *string `json:"jid,omitempty"`

	// Inactive
	Inactive *bool `json:"inactive,omitempty"`

	// Integrations
	Integrations *[]Contact `json:"integrations,omitempty"`

	// Presence
	Presence *Chatpresence `json:"presence,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Chatuserref) SetField(field string, fieldValue interface{}) {
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

func (o Chatuserref) MarshalJSON() ([]byte, error) {
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
	type Alias Chatuserref
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Jid *string `json:"jid,omitempty"`
		
		Inactive *bool `json:"inactive,omitempty"`
		
		Integrations *[]Contact `json:"integrations,omitempty"`
		
		Presence *Chatpresence `json:"presence,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		Jid: o.Jid,
		
		Inactive: o.Inactive,
		
		Integrations: o.Integrations,
		
		Presence: o.Presence,
		Alias:    (Alias)(o),
	})
}

func (o *Chatuserref) UnmarshalJSON(b []byte) error {
	var ChatuserrefMap map[string]interface{}
	err := json.Unmarshal(b, &ChatuserrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChatuserrefMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ChatuserrefMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := ChatuserrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Jid, ok := ChatuserrefMap["jid"].(string); ok {
		o.Jid = &Jid
	}
    
	if Inactive, ok := ChatuserrefMap["inactive"].(bool); ok {
		o.Inactive = &Inactive
	}
    
	if Integrations, ok := ChatuserrefMap["integrations"].([]interface{}); ok {
		IntegrationsString, _ := json.Marshal(Integrations)
		json.Unmarshal(IntegrationsString, &o.Integrations)
	}
	
	if Presence, ok := ChatuserrefMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatuserref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
