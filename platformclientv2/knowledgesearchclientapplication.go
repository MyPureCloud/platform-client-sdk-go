package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesearchclientapplication
type Knowledgesearchclientapplication struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Application type.
	VarType *string `json:"type,omitempty"`

	// Deployment - Application details when type is MessengerKnowledgeApp or SupportCenter.
	Deployment *Addressableentityref `json:"deployment,omitempty"`

	// BotFlow - Application details when type is BotFlow.
	BotFlow *Addressableentityref `json:"botFlow,omitempty"`

	// Assistant - Application details when type is Assistant.
	Assistant *Addressableentityref `json:"assistant,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesearchclientapplication) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesearchclientapplication) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesearchclientapplication
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Deployment *Addressableentityref `json:"deployment,omitempty"`
		
		BotFlow *Addressableentityref `json:"botFlow,omitempty"`
		
		Assistant *Addressableentityref `json:"assistant,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Deployment: o.Deployment,
		
		BotFlow: o.BotFlow,
		
		Assistant: o.Assistant,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesearchclientapplication) UnmarshalJSON(b []byte) error {
	var KnowledgesearchclientapplicationMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesearchclientapplicationMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := KnowledgesearchclientapplicationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Deployment, ok := KnowledgesearchclientapplicationMap["deployment"].(map[string]interface{}); ok {
		DeploymentString, _ := json.Marshal(Deployment)
		json.Unmarshal(DeploymentString, &o.Deployment)
	}
	
	if BotFlow, ok := KnowledgesearchclientapplicationMap["botFlow"].(map[string]interface{}); ok {
		BotFlowString, _ := json.Marshal(BotFlow)
		json.Unmarshal(BotFlowString, &o.BotFlow)
	}
	
	if Assistant, ok := KnowledgesearchclientapplicationMap["assistant"].(map[string]interface{}); ok {
		AssistantString, _ := json.Marshal(Assistant)
		json.Unmarshal(AssistantString, &o.Assistant)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesearchclientapplication) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
