package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentrequest
type Knowledgedocumentrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - Document type according to assigned template
	VarType *string `json:"type,omitempty"`

	// ExternalUrl - External Url to the document
	ExternalUrl *string `json:"externalUrl,omitempty"`

	// Faq - Faq document details
	Faq *Documentfaq `json:"faq,omitempty"`

	// Categories - Document categories
	Categories *[]Documentcategoryinput `json:"categories,omitempty"`

	// Article - Article details
	Article *Documentarticle `json:"article,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentrequest) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgedocumentrequest
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ExternalUrl *string `json:"externalUrl,omitempty"`
		
		Faq *Documentfaq `json:"faq,omitempty"`
		
		Categories *[]Documentcategoryinput `json:"categories,omitempty"`
		
		Article *Documentarticle `json:"article,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		ExternalUrl: o.ExternalUrl,
		
		Faq: o.Faq,
		
		Categories: o.Categories,
		
		Article: o.Article,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentrequest) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentrequestMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentrequestMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := KnowledgedocumentrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if ExternalUrl, ok := KnowledgedocumentrequestMap["externalUrl"].(string); ok {
		o.ExternalUrl = &ExternalUrl
	}
    
	if Faq, ok := KnowledgedocumentrequestMap["faq"].(map[string]interface{}); ok {
		FaqString, _ := json.Marshal(Faq)
		json.Unmarshal(FaqString, &o.Faq)
	}
	
	if Categories, ok := KnowledgedocumentrequestMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if Article, ok := KnowledgedocumentrequestMap["article"].(map[string]interface{}); ok {
		ArticleString, _ := json.Marshal(Article)
		json.Unmarshal(ArticleString, &o.Article)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
