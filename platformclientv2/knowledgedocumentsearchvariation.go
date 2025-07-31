package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentsearchvariation
type Knowledgedocumentsearchvariation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the variation.
	Id *string `json:"id,omitempty"`

	// DateCreated - The creation date-time for the document variation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The last modification date-time for the document variation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DocumentVersion - The version of the document.
	DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`

	// Contexts - The context values associated with the variation.
	Contexts *[]Documentvariationcontext `json:"contexts,omitempty"`

	// Document - The reference to document to which the variation is associated.
	Document *Knowledgedocumentreference `json:"document,omitempty"`

	// Priority - The priority of the variation.
	Priority *int `json:"priority,omitempty"`

	// Name - The name of the variation.
	Name *string `json:"name,omitempty"`

	// Body - The content for the variation.
	Body *Documentbodywithhighlight `json:"body,omitempty"`

	// Chunks - The chunk blocks associated with the variation.
	Chunks *[]Documentvariationsearchchunkblock `json:"chunks,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentsearchvariation) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentsearchvariation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
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
	type Alias Knowledgedocumentsearchvariation
	
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
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`
		
		Contexts *[]Documentvariationcontext `json:"contexts,omitempty"`
		
		Document *Knowledgedocumentreference `json:"document,omitempty"`
		
		Priority *int `json:"priority,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Body *Documentbodywithhighlight `json:"body,omitempty"`
		
		Chunks *[]Documentvariationsearchchunkblock `json:"chunks,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DocumentVersion: o.DocumentVersion,
		
		Contexts: o.Contexts,
		
		Document: o.Document,
		
		Priority: o.Priority,
		
		Name: o.Name,
		
		Body: o.Body,
		
		Chunks: o.Chunks,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentsearchvariation) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentsearchvariationMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentsearchvariationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentsearchvariationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateCreatedString, ok := KnowledgedocumentsearchvariationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgedocumentsearchvariationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if DocumentVersion, ok := KnowledgedocumentsearchvariationMap["documentVersion"].(map[string]interface{}); ok {
		DocumentVersionString, _ := json.Marshal(DocumentVersion)
		json.Unmarshal(DocumentVersionString, &o.DocumentVersion)
	}
	
	if Contexts, ok := KnowledgedocumentsearchvariationMap["contexts"].([]interface{}); ok {
		ContextsString, _ := json.Marshal(Contexts)
		json.Unmarshal(ContextsString, &o.Contexts)
	}
	
	if Document, ok := KnowledgedocumentsearchvariationMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if Priority, ok := KnowledgedocumentsearchvariationMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if Name, ok := KnowledgedocumentsearchvariationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Body, ok := KnowledgedocumentsearchvariationMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if Chunks, ok := KnowledgedocumentsearchvariationMap["chunks"].([]interface{}); ok {
		ChunksString, _ := json.Marshal(Chunks)
		json.Unmarshal(ChunksString, &o.Chunks)
	}
	
	if SelfUri, ok := KnowledgedocumentsearchvariationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchvariation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
