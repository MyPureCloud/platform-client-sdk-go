package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestedsearchchunk
type Suggestedsearchchunk struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - The document title.
	Title *string `json:"title,omitempty"`

	// Uri - The document uri.
	Uri *string `json:"uri,omitempty"`

	// ChunkId - The chunk ID.
	ChunkId *string `json:"chunkId,omitempty"`

	// Text - The text of the knowledge chunk.
	Text *string `json:"text,omitempty"`

	// Confidence - Value between 0 and 1. 1 corresponds to very confident, 0 to not confident at all.
	Confidence *float32 `json:"confidence,omitempty"`

	// Document - The article.
	Document *Addressableentityref `json:"document,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Suggestedsearchchunk) SetField(field string, fieldValue interface{}) {
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

func (o Suggestedsearchchunk) MarshalJSON() ([]byte, error) {
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
	type Alias Suggestedsearchchunk
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Uri *string `json:"uri,omitempty"`
		
		ChunkId *string `json:"chunkId,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Confidence *float32 `json:"confidence,omitempty"`
		
		Document *Addressableentityref `json:"document,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Uri: o.Uri,
		
		ChunkId: o.ChunkId,
		
		Text: o.Text,
		
		Confidence: o.Confidence,
		
		Document: o.Document,
		Alias:    (Alias)(o),
	})
}

func (o *Suggestedsearchchunk) UnmarshalJSON(b []byte) error {
	var SuggestedsearchchunkMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestedsearchchunkMap)
	if err != nil {
		return err
	}
	
	if Title, ok := SuggestedsearchchunkMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Uri, ok := SuggestedsearchchunkMap["uri"].(string); ok {
		o.Uri = &Uri
	}
    
	if ChunkId, ok := SuggestedsearchchunkMap["chunkId"].(string); ok {
		o.ChunkId = &ChunkId
	}
    
	if Text, ok := SuggestedsearchchunkMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Confidence, ok := SuggestedsearchchunkMap["confidence"].(float64); ok {
		ConfidenceFloat32 := float32(Confidence)
		o.Confidence = &ConfidenceFloat32
	}
	
	if Document, ok := SuggestedsearchchunkMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestedsearchchunk) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
