package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedmappingpreview
type Agentownedmappingpreview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AgentOwnedColumn - The raw value of the agent-owned column
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`

	// Email - The email address of the user, if it exists
	Email *string `json:"email,omitempty"`

	// UserId - The id of the user, if it exists
	UserId *string `json:"userId,omitempty"`

	// Exists - Whether the user exists
	Exists *bool `json:"exists,omitempty"`

	// IsQueueMember - Whether the user is a member of the campaign's queue
	IsQueueMember *bool `json:"isQueueMember,omitempty"`

	// RecordCount - The number of contact records whose agent-owned column matches the raw value
	RecordCount *int `json:"recordCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agentownedmappingpreview) SetField(field string, fieldValue interface{}) {
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

func (o Agentownedmappingpreview) MarshalJSON() ([]byte, error) {
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
	type Alias Agentownedmappingpreview
	
	return json.Marshal(&struct { 
		AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Exists *bool `json:"exists,omitempty"`
		
		IsQueueMember *bool `json:"isQueueMember,omitempty"`
		
		RecordCount *int `json:"recordCount,omitempty"`
		Alias
	}{ 
		AgentOwnedColumn: o.AgentOwnedColumn,
		
		Email: o.Email,
		
		UserId: o.UserId,
		
		Exists: o.Exists,
		
		IsQueueMember: o.IsQueueMember,
		
		RecordCount: o.RecordCount,
		Alias:    (Alias)(o),
	})
}

func (o *Agentownedmappingpreview) UnmarshalJSON(b []byte) error {
	var AgentownedmappingpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &AgentownedmappingpreviewMap)
	if err != nil {
		return err
	}
	
	if AgentOwnedColumn, ok := AgentownedmappingpreviewMap["agentOwnedColumn"].(string); ok {
		o.AgentOwnedColumn = &AgentOwnedColumn
	}
    
	if Email, ok := AgentownedmappingpreviewMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if UserId, ok := AgentownedmappingpreviewMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Exists, ok := AgentownedmappingpreviewMap["exists"].(bool); ok {
		o.Exists = &Exists
	}
    
	if IsQueueMember, ok := AgentownedmappingpreviewMap["isQueueMember"].(bool); ok {
		o.IsQueueMember = &IsQueueMember
	}
    
	if RecordCount, ok := AgentownedmappingpreviewMap["recordCount"].(float64); ok {
		RecordCountInt := int(RecordCount)
		o.RecordCount = &RecordCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
