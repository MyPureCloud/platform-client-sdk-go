package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentresponse
type Knowledgedocumentresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Title - Document title, having a limit of 500 words.
	Title *string `json:"title,omitempty"`

	// Visible - Indicates if the knowledge document should be included in search results.
	Visible *bool `json:"visible,omitempty"`

	// Alternatives - List of alternate phrases related to the title which improves search results.
	Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`

	// State - State of the document.
	State *string `json:"state,omitempty"`

	// DateCreated - Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateImported - Document import date-time, or null if was not imported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateImported *time.Time `json:"dateImported,omitempty"`

	// LastPublishedVersionNumber - The last published version number of the document.
	LastPublishedVersionNumber *int `json:"lastPublishedVersionNumber,omitempty"`

	// DatePublished - The date on which the document was last published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`

	// CreatedBy - The user who created the document.
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// ModifiedBy - The user who modified the document.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DocumentVersion - The version of the document.
	DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`

	// Category - The reference to category associated with the document.
	Category *Categoryresponse `json:"category,omitempty"`

	// Labels - The references to labels associated with the document.
	Labels *[]Labelresponse `json:"labels,omitempty"`

	// KnowledgeBase - Knowledge base to which the document belongs to.
	KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`

	// ExternalId - The reference to external id associated with the document.
	ExternalId *string `json:"externalId,omitempty"`

	// Source - The reference to source associated with the document.
	Source *Addressableentityref `json:"source,omitempty"`

	// Readonly - Whether the document is read-only.
	Readonly *bool `json:"readonly,omitempty"`

	// Variations - Variations of the document.
	Variations *[]Documentvariation `json:"variations,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgedocumentresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgedocumentresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateImported","DatePublished", }
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
	type Alias Knowledgedocumentresponse
	
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
	
	DateImported := new(string)
	if o.DateImported != nil {
		
		*DateImported = timeutil.Strftime(o.DateImported, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateImported = nil
	}
	
	DatePublished := new(string)
	if o.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(o.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Visible *bool `json:"visible,omitempty"`
		
		Alternatives *[]Knowledgedocumentalternative `json:"alternatives,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateImported *string `json:"dateImported,omitempty"`
		
		LastPublishedVersionNumber *int `json:"lastPublishedVersionNumber,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`
		
		Category *Categoryresponse `json:"category,omitempty"`
		
		Labels *[]Labelresponse `json:"labels,omitempty"`
		
		KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`
		
		ExternalId *string `json:"externalId,omitempty"`
		
		Source *Addressableentityref `json:"source,omitempty"`
		
		Readonly *bool `json:"readonly,omitempty"`
		
		Variations *[]Documentvariation `json:"variations,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Title: o.Title,
		
		Visible: o.Visible,
		
		Alternatives: o.Alternatives,
		
		State: o.State,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateImported: DateImported,
		
		LastPublishedVersionNumber: o.LastPublishedVersionNumber,
		
		DatePublished: DatePublished,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		DocumentVersion: o.DocumentVersion,
		
		Category: o.Category,
		
		Labels: o.Labels,
		
		KnowledgeBase: o.KnowledgeBase,
		
		ExternalId: o.ExternalId,
		
		Source: o.Source,
		
		Readonly: o.Readonly,
		
		Variations: o.Variations,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgedocumentresponse) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Title, ok := KnowledgedocumentresponseMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Visible, ok := KnowledgedocumentresponseMap["visible"].(bool); ok {
		o.Visible = &Visible
	}
    
	if Alternatives, ok := KnowledgedocumentresponseMap["alternatives"].([]interface{}); ok {
		AlternativesString, _ := json.Marshal(Alternatives)
		json.Unmarshal(AlternativesString, &o.Alternatives)
	}
	
	if State, ok := KnowledgedocumentresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if dateCreatedString, ok := KnowledgedocumentresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgedocumentresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateImportedString, ok := KnowledgedocumentresponseMap["dateImported"].(string); ok {
		DateImported, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateImportedString)
		o.DateImported = &DateImported
	}
	
	if LastPublishedVersionNumber, ok := KnowledgedocumentresponseMap["lastPublishedVersionNumber"].(float64); ok {
		LastPublishedVersionNumberInt := int(LastPublishedVersionNumber)
		o.LastPublishedVersionNumber = &LastPublishedVersionNumberInt
	}
	
	if datePublishedString, ok := KnowledgedocumentresponseMap["datePublished"].(string); ok {
		DatePublished, _ := time.Parse("2006-01-02T15:04:05.999999Z", datePublishedString)
		o.DatePublished = &DatePublished
	}
	
	if CreatedBy, ok := KnowledgedocumentresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := KnowledgedocumentresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if DocumentVersion, ok := KnowledgedocumentresponseMap["documentVersion"].(map[string]interface{}); ok {
		DocumentVersionString, _ := json.Marshal(DocumentVersion)
		json.Unmarshal(DocumentVersionString, &o.DocumentVersion)
	}
	
	if Category, ok := KnowledgedocumentresponseMap["category"].(map[string]interface{}); ok {
		CategoryString, _ := json.Marshal(Category)
		json.Unmarshal(CategoryString, &o.Category)
	}
	
	if Labels, ok := KnowledgedocumentresponseMap["labels"].([]interface{}); ok {
		LabelsString, _ := json.Marshal(Labels)
		json.Unmarshal(LabelsString, &o.Labels)
	}
	
	if KnowledgeBase, ok := KnowledgedocumentresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if ExternalId, ok := KnowledgedocumentresponseMap["externalId"].(string); ok {
		o.ExternalId = &ExternalId
	}
    
	if Source, ok := KnowledgedocumentresponseMap["source"].(map[string]interface{}); ok {
		SourceString, _ := json.Marshal(Source)
		json.Unmarshal(SourceString, &o.Source)
	}
	
	if Readonly, ok := KnowledgedocumentresponseMap["readonly"].(bool); ok {
		o.Readonly = &Readonly
	}
    
	if Variations, ok := KnowledgedocumentresponseMap["variations"].([]interface{}); ok {
		VariationsString, _ := json.Marshal(Variations)
		json.Unmarshal(VariationsString, &o.Variations)
	}
	
	if SelfUri, ok := KnowledgedocumentresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
