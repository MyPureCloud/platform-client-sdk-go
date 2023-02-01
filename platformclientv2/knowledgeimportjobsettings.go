package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeimportjobsettings
type Knowledgeimportjobsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ImportAsNew - If enabled import creates a new document even if update is available.
	ImportAsNew *bool `json:"importAsNew,omitempty"`

	// Visible - If specified, import will override the visibility of the imported documents.
	Visible *bool `json:"visible,omitempty"`

	// CategoryId - If specified, import will override the category of the imported documents.
	CategoryId *string `json:"categoryId,omitempty"`

	// LabelIds - If specified, import will add this labels to the imported documents.
	LabelIds *[]string `json:"labelIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgeimportjobsettings) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgeimportjobsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgeimportjobsettings
	
	return json.Marshal(&struct { 
		ImportAsNew *bool `json:"importAsNew,omitempty"`
		
		Visible *bool `json:"visible,omitempty"`
		
		CategoryId *string `json:"categoryId,omitempty"`
		
		LabelIds *[]string `json:"labelIds,omitempty"`
		Alias
	}{ 
		ImportAsNew: o.ImportAsNew,
		
		Visible: o.Visible,
		
		CategoryId: o.CategoryId,
		
		LabelIds: o.LabelIds,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgeimportjobsettings) UnmarshalJSON(b []byte) error {
	var KnowledgeimportjobsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeimportjobsettingsMap)
	if err != nil {
		return err
	}
	
	if ImportAsNew, ok := KnowledgeimportjobsettingsMap["importAsNew"].(bool); ok {
		o.ImportAsNew = &ImportAsNew
	}
    
	if Visible, ok := KnowledgeimportjobsettingsMap["visible"].(bool); ok {
		o.Visible = &Visible
	}
    
	if CategoryId, ok := KnowledgeimportjobsettingsMap["categoryId"].(string); ok {
		o.CategoryId = &CategoryId
	}
    
	if LabelIds, ok := KnowledgeimportjobsettingsMap["labelIds"].([]interface{}); ok {
		LabelIdsString, _ := json.Marshal(LabelIds)
		json.Unmarshal(LabelIdsString, &o.LabelIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
