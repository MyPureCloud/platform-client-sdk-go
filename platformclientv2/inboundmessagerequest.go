package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Inboundmessagerequest
type Inboundmessagerequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueueId - The ID of the queue to use for routing the email conversation. This field is mutually exclusive with flowId
	QueueId *string `json:"queueId,omitempty"`

	// FlowId - The ID of the flow to use for routing email conversation. This field is mutually exclusive with queueId
	FlowId *string `json:"flowId,omitempty"`

	// Provider - The name of the provider that is sourcing the email such as Oracle, Salesforce, etc.
	Provider *string `json:"provider,omitempty"`

	// SkillIds - The list of skill ID's to use for routing.
	SkillIds *[]string `json:"skillIds,omitempty"`

	// LanguageId - The ID of the language to use for routing.
	LanguageId *string `json:"languageId,omitempty"`

	// Priority - The priority to assign to the conversation for routing.
	Priority *int `json:"priority,omitempty"`

	// Attributes - The list of attributes to associate with the customer participant.
	Attributes *map[string]string `json:"attributes,omitempty"`

	// ToAddress - The email address of the recipient of the email.
	ToAddress *string `json:"toAddress,omitempty"`

	// ToName - The name of the recipient of the email.
	ToName *string `json:"toName,omitempty"`

	// FromAddress - The email address of the sender of the email.
	FromAddress *string `json:"fromAddress,omitempty"`

	// FromName - The name of the sender of the email.
	FromName *string `json:"fromName,omitempty"`

	// Subject - The subject of the email
	Subject *string `json:"subject,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Inboundmessagerequest) SetField(field string, fieldValue interface{}) {
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

func (o Inboundmessagerequest) MarshalJSON() ([]byte, error) {
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
	type Alias Inboundmessagerequest
	
	return json.Marshal(&struct { 
		QueueId *string `json:"queueId,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ToAddress *string `json:"toAddress,omitempty"`
		
		ToName *string `json:"toName,omitempty"`
		
		FromAddress *string `json:"fromAddress,omitempty"`
		
		FromName *string `json:"fromName,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		Alias
	}{ 
		QueueId: o.QueueId,
		
		FlowId: o.FlowId,
		
		Provider: o.Provider,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		Priority: o.Priority,
		
		Attributes: o.Attributes,
		
		ToAddress: o.ToAddress,
		
		ToName: o.ToName,
		
		FromAddress: o.FromAddress,
		
		FromName: o.FromName,
		
		Subject: o.Subject,
		Alias:    (Alias)(o),
	})
}

func (o *Inboundmessagerequest) UnmarshalJSON(b []byte) error {
	var InboundmessagerequestMap map[string]interface{}
	err := json.Unmarshal(b, &InboundmessagerequestMap)
	if err != nil {
		return err
	}
	
	if QueueId, ok := InboundmessagerequestMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if FlowId, ok := InboundmessagerequestMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if Provider, ok := InboundmessagerequestMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if SkillIds, ok := InboundmessagerequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := InboundmessagerequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Priority, ok := InboundmessagerequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Attributes, ok := InboundmessagerequestMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ToAddress, ok := InboundmessagerequestMap["toAddress"].(string); ok {
		o.ToAddress = &ToAddress
	}
    
	if ToName, ok := InboundmessagerequestMap["toName"].(string); ok {
		o.ToName = &ToName
	}
    
	if FromAddress, ok := InboundmessagerequestMap["fromAddress"].(string); ok {
		o.FromAddress = &FromAddress
	}
    
	if FromName, ok := InboundmessagerequestMap["fromName"].(string); ok {
		o.FromName = &FromName
	}
    
	if Subject, ok := InboundmessagerequestMap["subject"].(string); ok {
		o.Subject = &Subject
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Inboundmessagerequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
