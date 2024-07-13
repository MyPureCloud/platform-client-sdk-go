package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2wemengagementcelebrationupdatestopicengagementcelebration
type V2wemengagementcelebrationupdatestopicengagementcelebration struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Recipient
	Recipient *V2wemengagementcelebrationupdatestopicuserid `json:"recipient,omitempty"`

	// CreatedBy
	CreatedBy *V2wemengagementcelebrationupdatestopicuserid `json:"createdBy,omitempty"`

	// DateCreated
	DateCreated *string `json:"dateCreated,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Note
	Note *string `json:"note,omitempty"`

	// SourceEntity
	SourceEntity *V2wemengagementcelebrationupdatestopicsourceentity `json:"sourceEntity,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2wemengagementcelebrationupdatestopicengagementcelebration) SetField(field string, fieldValue interface{}) {
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

func (o V2wemengagementcelebrationupdatestopicengagementcelebration) MarshalJSON() ([]byte, error) {
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
	type Alias V2wemengagementcelebrationupdatestopicengagementcelebration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Recipient *V2wemengagementcelebrationupdatestopicuserid `json:"recipient,omitempty"`
		
		CreatedBy *V2wemengagementcelebrationupdatestopicuserid `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		SourceEntity *V2wemengagementcelebrationupdatestopicsourceentity `json:"sourceEntity,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Recipient: o.Recipient,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: o.DateCreated,
		
		VarType: o.VarType,
		
		Title: o.Title,
		
		Note: o.Note,
		
		SourceEntity: o.SourceEntity,
		Alias:    (Alias)(o),
	})
}

func (o *V2wemengagementcelebrationupdatestopicengagementcelebration) UnmarshalJSON(b []byte) error {
	var V2wemengagementcelebrationupdatestopicengagementcelebrationMap map[string]interface{}
	err := json.Unmarshal(b, &V2wemengagementcelebrationupdatestopicengagementcelebrationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Recipient, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if CreatedBy, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if DateCreated, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["dateCreated"].(string); ok {
		o.DateCreated = &DateCreated
	}
    
	if VarType, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Title, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Note, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["note"].(string); ok {
		o.Note = &Note
	}
    
	if SourceEntity, ok := V2wemengagementcelebrationupdatestopicengagementcelebrationMap["sourceEntity"].(map[string]interface{}); ok {
		SourceEntityString, _ := json.Marshal(SourceEntity)
		json.Unmarshal(SourceEntityString, &o.SourceEntity)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2wemengagementcelebrationupdatestopicengagementcelebration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
