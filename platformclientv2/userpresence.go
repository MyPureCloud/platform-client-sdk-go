package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userpresence
type Userpresence struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Source - Deprecated - The sourceID field should be used as a replacement.
	Source *string `json:"source,omitempty"`

	// SourceId - Represents the ID of a registered source
	SourceId *string `json:"sourceId,omitempty"`

	// Primary - A boolean used to tell whether or not to set this presence source as the primary on a PATCH
	Primary *bool `json:"primary,omitempty"`

	// PresenceDefinition
	PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`

	// FuturePresenceDefinition - The definition of the user's desired future presence
	FuturePresenceDefinition *Presencedefinition `json:"futurePresenceDefinition,omitempty"`

	// Message
	Message *string `json:"message,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userpresence) SetField(field string, fieldValue interface{}) {
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

func (o Userpresence) MarshalJSON() ([]byte, error) {
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
	type Alias Userpresence
	
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
		
		SourceId *string `json:"sourceId,omitempty"`
		
		Primary *bool `json:"primary,omitempty"`
		
		PresenceDefinition *Presencedefinition `json:"presenceDefinition,omitempty"`
		
		FuturePresenceDefinition *Presencedefinition `json:"futurePresenceDefinition,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Source: o.Source,
		
		SourceId: o.SourceId,
		
		Primary: o.Primary,
		
		PresenceDefinition: o.PresenceDefinition,
		
		FuturePresenceDefinition: o.FuturePresenceDefinition,
		
		Message: o.Message,
		
		ModifiedDate: ModifiedDate,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Userpresence) UnmarshalJSON(b []byte) error {
	var UserpresenceMap map[string]interface{}
	err := json.Unmarshal(b, &UserpresenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserpresenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserpresenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Source, ok := UserpresenceMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if SourceId, ok := UserpresenceMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    
	if Primary, ok := UserpresenceMap["primary"].(bool); ok {
		o.Primary = &Primary
	}
    
	if PresenceDefinition, ok := UserpresenceMap["presenceDefinition"].(map[string]interface{}); ok {
		PresenceDefinitionString, _ := json.Marshal(PresenceDefinition)
		json.Unmarshal(PresenceDefinitionString, &o.PresenceDefinition)
	}
	
	if FuturePresenceDefinition, ok := UserpresenceMap["futurePresenceDefinition"].(map[string]interface{}); ok {
		FuturePresenceDefinitionString, _ := json.Marshal(FuturePresenceDefinition)
		json.Unmarshal(FuturePresenceDefinitionString, &o.FuturePresenceDefinition)
	}
	
	if Message, ok := UserpresenceMap["message"].(string); ok {
		o.Message = &Message
	}
    
	if modifiedDateString, ok := UserpresenceMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if SelfUri, ok := UserpresenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userpresence) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
