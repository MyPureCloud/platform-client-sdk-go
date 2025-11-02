package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Selfagentgreeting
type Selfagentgreeting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// InboundPrompt - The agent greeting prompt to use when inbound calls are connected
	InboundPrompt *Prompt `json:"inboundPrompt,omitempty"`

	// OutboundPrompt - The agent greeting prompt to use when outbound calls are connected
	OutboundPrompt *Prompt `json:"outboundPrompt,omitempty"`

	// InboundPromptDefaultLanguage - The default language to use for the agent greeting inbound prompt
	InboundPromptDefaultLanguage *string `json:"inboundPromptDefaultLanguage,omitempty"`

	// OutboundPromptDefaultLanguage - The default language to use for the agent greeting outbound prompt
	OutboundPromptDefaultLanguage *string `json:"outboundPromptDefaultLanguage,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Selfagentgreeting) SetField(field string, fieldValue interface{}) {
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

func (o Selfagentgreeting) MarshalJSON() ([]byte, error) {
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
	type Alias Selfagentgreeting
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		InboundPrompt *Prompt `json:"inboundPrompt,omitempty"`
		
		OutboundPrompt *Prompt `json:"outboundPrompt,omitempty"`
		
		InboundPromptDefaultLanguage *string `json:"inboundPromptDefaultLanguage,omitempty"`
		
		OutboundPromptDefaultLanguage *string `json:"outboundPromptDefaultLanguage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		InboundPrompt: o.InboundPrompt,
		
		OutboundPrompt: o.OutboundPrompt,
		
		InboundPromptDefaultLanguage: o.InboundPromptDefaultLanguage,
		
		OutboundPromptDefaultLanguage: o.OutboundPromptDefaultLanguage,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Selfagentgreeting) UnmarshalJSON(b []byte) error {
	var SelfagentgreetingMap map[string]interface{}
	err := json.Unmarshal(b, &SelfagentgreetingMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SelfagentgreetingMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SelfagentgreetingMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InboundPrompt, ok := SelfagentgreetingMap["inboundPrompt"].(map[string]interface{}); ok {
		InboundPromptString, _ := json.Marshal(InboundPrompt)
		json.Unmarshal(InboundPromptString, &o.InboundPrompt)
	}
	
	if OutboundPrompt, ok := SelfagentgreetingMap["outboundPrompt"].(map[string]interface{}); ok {
		OutboundPromptString, _ := json.Marshal(OutboundPrompt)
		json.Unmarshal(OutboundPromptString, &o.OutboundPrompt)
	}
	
	if InboundPromptDefaultLanguage, ok := SelfagentgreetingMap["inboundPromptDefaultLanguage"].(string); ok {
		o.InboundPromptDefaultLanguage = &InboundPromptDefaultLanguage
	}
    
	if OutboundPromptDefaultLanguage, ok := SelfagentgreetingMap["outboundPromptDefaultLanguage"].(string); ok {
		o.OutboundPromptDefaultLanguage = &OutboundPromptDefaultLanguage
	}
    
	if SelfUri, ok := SelfagentgreetingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Selfagentgreeting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
