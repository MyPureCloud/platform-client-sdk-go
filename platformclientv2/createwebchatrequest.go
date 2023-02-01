package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatrequest
type Createwebchatrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - The ID of the queue to use for routing the chat conversation.
	QueueId *string `json:"queueId,omitempty"`

	// Provider - The name of the provider that is sourcing the web chat.
	Provider *string `json:"provider,omitempty"`

	// SkillIds - The list of skill ID's to use for routing.
	SkillIds *[]string `json:"skillIds,omitempty"`

	// LanguageId - The ID of the langauge to use for routing.
	LanguageId *string `json:"languageId,omitempty"`

	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`

	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`

	// CustomerName - The name of the customer participating in the web chat.
	CustomerName *string `json:"customerName,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createwebchatrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createwebchatrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createwebchatrequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		CustomerName *string `json:"customerName,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		Provider: o.Provider,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		Priority: o.Priority,
		
		Attributes: o.Attributes,
		
		CustomerName: o.CustomerName,
		Alias:    (Alias)(o),
	})
}

func (o *Createwebchatrequest) UnmarshalJSON(b []byte) error {
	var CreatewebchatrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatrequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := CreatewebchatrequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Provider, ok := CreatewebchatrequestMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if SkillIds, ok := CreatewebchatrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := CreatewebchatrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Priority, ok := CreatewebchatrequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Attributes, ok := CreatewebchatrequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if CustomerName, ok := CreatewebchatrequestMap["customerName"].(string); ok {
		o.CustomerName = &CustomerName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
