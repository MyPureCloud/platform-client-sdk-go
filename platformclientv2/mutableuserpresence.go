package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mutableuserpresence
type Mutableuserpresence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The user's id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Source - Represents the source where the Presence was set. Some examples are: PURECLOUD, LYNC, OUTLOOK, etc.
	Source *string `json:"source,omitempty"`

	// Primary - A boolean used to tell whether or not to set this presence source as the primary on a PATCH
	Primary *bool `json:"primary,omitempty"`

	// PresenceDefinition
	PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mutableuserpresence) SetField(field string, fieldValue interface{}) {
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

func (o Mutableuserpresence) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifiedDate", }
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
	type Alias Mutableuserpresence
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Source: o.Source,
		
		Primary: o.Primary,
		
		PresenceDefinition: o.PresenceDefinition,
		
		Message: o.Message,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Mutableuserpresence) UnmarshalJSON(b []byte) error {
	var MutableuserpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &MutableuserpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MutableuserpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := MutableuserpresenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Source, ok := MutableuserpresenceMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Primary, ok := MutableuserpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if PresenceDefinition, ok := MutableuserpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if Message, ok := MutableuserpresenceMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if modifiedDateString, ok := MutableuserpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SelfUri, ok := MutableuserpresenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Mutableuserpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
