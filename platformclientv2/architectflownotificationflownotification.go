package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationflownotification
type Architectflownotificationflownotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The flow ID
	Id *string `json:"id,omitempty"`

	// Name - The flow name
	Name *string `json:"name,omitempty"`

	// Description - The flow description
	Description *string `json:"description,omitempty"`

	// Deleted - The flow deleted state
	Deleted *bool `json:"deleted,omitempty"`

	// CheckedInVersion
	CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`

	// SavedVersion - A bare-bones flow version object
	SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`

	// PublishedVersion - A bare-bones flow version object
	PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`

	// CurrentOperation
	CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Architectflownotificationflownotification) SetField(field string, fieldValue interface{}) {
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

func (o Architectflownotificationflownotification) MarshalJSON() ([]byte, error) {
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
	type Alias Architectflownotificationflownotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`
		
		SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`
		
		PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`
		
		CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Deleted: o.Deleted,
		
		CheckedInVersion: o.CheckedInVersion,
		
		SavedVersion: o.SavedVersion,
		
		PublishedVersion: o.PublishedVersion,
		
		CurrentOperation: o.CurrentOperation,
		Alias:    (Alias)(o),
	})
}

func (o *Architectflownotificationflownotification) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationflownotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationflownotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationflownotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectflownotificationflownotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ArchitectflownotificationflownotificationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Deleted, ok := ArchitectflownotificationflownotificationMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if CheckedInVersion, ok := ArchitectflownotificationflownotificationMap["checkedInVersion"].(map[string]interface{}); ok {
		CheckedInVersionString, _ := json.Marshal(CheckedInVersion)
		json.Unmarshal(CheckedInVersionString, &o.CheckedInVersion)
	}
	
	if SavedVersion, ok := ArchitectflownotificationflownotificationMap["savedVersion"].(map[string]interface{}); ok {
		SavedVersionString, _ := json.Marshal(SavedVersion)
		json.Unmarshal(SavedVersionString, &o.SavedVersion)
	}
	
	if PublishedVersion, ok := ArchitectflownotificationflownotificationMap["publishedVersion"].(map[string]interface{}); ok {
		PublishedVersionString, _ := json.Marshal(PublishedVersion)
		json.Unmarshal(PublishedVersionString, &o.PublishedVersion)
	}
	
	if CurrentOperation, ok := ArchitectflownotificationflownotificationMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationflownotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
