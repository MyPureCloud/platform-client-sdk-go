package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentreq
type Knowledgedocumentreq struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Title - Document title.
	Title *string `json:"title,omitempty"`

	// Visible - Indicates if the knowledge document should be included in search results.
	Visible *bool `json:"visible,omitempty"`

	// Alternatives - List of alternate phrases related to the title which improves search results.
	Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`

	// CategoryId - The category associated with the document.
	CategoryId *string `json:"categoryId,omitempty"`

	// LabelIds - The ids of labels associated with the document.
	LabelIds *[]string `json:"labelIds,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentreq) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentreq) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentreq
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Visible *bool `json:"visible,omitempty"`
		
		Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`
		
		CategoryId *string `json:"categoryId,omitempty"`
		
		LabelIds *[]string `json:"labelIds,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Visible: o.Visible,
		
		Alternatives: o.Alternatives,
		
		CategoryId: o.CategoryId,
		
		LabelIds: o.LabelIds,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentreq) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentreqMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentreqMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentreqMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := KnowledgedocumentreqMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Visible, ok := KnowledgedocumentreqMap["visible"].(bool); ok {
		o.Visible = &Visible
	}
    
	if Alternatives, ok := KnowledgedocumentreqMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	
	if CategoryId, ok := KnowledgedocumentreqMap["categoryId"].(string); ok {
		o.CategoryId = &CategoryId
	}
    
	if LabelIds, ok := KnowledgedocumentreqMap["labelIds"].([]interface{}); ok {
		LabelIdsString, _ := json.Marshal(LabelIds)
		json.Unmarshal(LabelIdsString, &o.LabelIds)
	}
	
	if SelfUri, ok := KnowledgedocumentreqMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentreq) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
