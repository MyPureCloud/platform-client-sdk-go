package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgebase
type Knowledgebase struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - Knowledge base description
	Description *string `json:"description,omitempty"`

	// CoreLanguage - Core language for knowledge base in which initial content must be created, language codes [en-US, en-UK, en-AU, de-DE] are supported currently. However, the new DX knowledge will support all these language codes, along with 'early preview' language codes [ca-ES, tr-TR, sv-SE, fi-FI, nb-NO, da-DK, ja-JP] which might have a lower accuracy.
	CoreLanguage *string `json:"coreLanguage,omitempty"`

	// DateCreated - Knowledge base creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Knowledge base last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// FaqCount - The count representing the number of documents of type FAQ in the KnowledgeBase
	FaqCount *int `json:"faqCount,omitempty"`

	// DateDocumentLastModified - The date representing when the last document is modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDocumentLastModified *time.Time `json:"dateDocumentLastModified,omitempty"`

	// ArticleCount - The count representing the number of documents of type Article in the KnowledgeBase
	ArticleCount *int `json:"articleCount,omitempty"`

	// Published - Flag that indicates the knowledge base is published
	Published *bool `json:"published,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgebase) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgebase) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateDocumentLastModified", }
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
	type Alias Knowledgebase
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateDocumentLastModified := new(string)
	if o.DateDocumentLastModified != nil {
		
		*DateDocumentLastModified = timeutil.Strftime(o.DateDocumentLastModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDocumentLastModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CoreLanguage *string `json:"coreLanguage,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		FaqCount *int `json:"faqCount,omitempty"`
		
		DateDocumentLastModified *string `json:"dateDocumentLastModified,omitempty"`
		
		ArticleCount *int `json:"articleCount,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CoreLanguage: o.CoreLanguage,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		FaqCount: o.FaqCount,
		
		DateDocumentLastModified: DateDocumentLastModified,
		
		ArticleCount: o.ArticleCount,
		
		Published: o.Published,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgebase) UnmarshalJSON(b []byte) error {
	var KnowledgebaseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgebaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgebaseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KnowledgebaseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := KnowledgebaseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if CoreLanguage, ok := KnowledgebaseMap["coreLanguage"].(string); ok {
		o.CoreLanguage = &CoreLanguage
	}
    
	if dateCreatedString, ok := KnowledgebaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgebaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if FaqCount, ok := KnowledgebaseMap["faqCount"].(float64); ok {
		FaqCountInt := int(FaqCount)
		o.FaqCount = &FaqCountInt
	}
	
	if dateDocumentLastModifiedString, ok := KnowledgebaseMap["dateDocumentLastModified"].(string); ok {
		DateDocumentLastModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDocumentLastModifiedString)
		o.DateDocumentLastModified = &DateDocumentLastModified
	}
	
	if ArticleCount, ok := KnowledgebaseMap["articleCount"].(float64); ok {
		ArticleCountInt := int(ArticleCount)
		o.ArticleCount = &ArticleCountInt
	}
	
	if Published, ok := KnowledgebaseMap["published"].(bool); ok {
		o.Published = &Published
	}
    
	if SelfUri, ok := KnowledgebaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgebase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
