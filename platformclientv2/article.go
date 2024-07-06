package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Article
type Article struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - The article title.
	Title *string `json:"title,omitempty"`

	// Uri - The URI for the article.
	Uri *string `json:"uri,omitempty"`

	// Snippets - This contains snippets of text from the article matching the query.
	Snippets *[]string `json:"snippets,omitempty"`

	// Confidence - Value between 0 and 1. 1 corresponds to very confident, 0 to not confident at all.
	Confidence *float32 `json:"confidence,omitempty"`

	// Metadata - A map that contains custom metadata about the article answer.
	Metadata *map[string]Metadataattribute `json:"metadata,omitempty"`

	// Version - The version of the Article.
	Version *Addressableentityref `json:"version,omitempty"`

	// Variations - Variations of the Article.
	Variations *[]Addressableentityref `json:"variations,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Article) SetField(field string, fieldValue interface{}) {
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

func (o Article) MarshalJSON() ([]byte, error) {
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
	type Alias Article
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Uri *string `json:"uri,omitempty"`
		
		Snippets *[]string `json:"snippets,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		
		Metadata *map[string]Metadataattribute `json:"metadata,omitempty"`
		
		Version *Addressableentityref `json:"version,omitempty"`
		
		Variations *[]Addressableentityref `json:"variations,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Uri: o.Uri,
		
		Snippets: o.Snippets,
		
		Confidence: o.Confidence,
		
		Metadata: o.Metadata,
		
		Version: o.Version,
		
		Variations: o.Variations,
		Alias:    (Alias)(o),
	})
}

func (o *Article) UnmarshalJSON(b []byte) error {
	var ArticleMap map[string]interface{}
	err := json.Unmarshal(b, &ArticleMap)
	if err != nil {
		return err
	}
	
	if Title, ok := ArticleMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Uri, ok := ArticleMap["uri"].(string); ok {
		o.Uri = &Uri
	}
    
	if Snippets, ok := ArticleMap["snippets"].([]interface{}); ok {
		SnippetsString, _ := json.Marshal(Snippets)
		json.Unmarshal(SnippetsString, &o.Snippets)
	}
	
	if Confidence, ok := ArticleMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if Metadata, ok := ArticleMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	
	if Version, ok := ArticleMap["version"].(map[string]interface{}); ok {
		VersionString, _ := json.Marshal(Version)
		json.Unmarshal(VersionString, &o.Version)
	}
	
	if Variations, ok := ArticleMap["variations"].([]interface{}); ok {
		VariationsString, _ := json.Marshal(Variations)
		json.Unmarshal(VariationsString, &o.Variations)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Article) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
