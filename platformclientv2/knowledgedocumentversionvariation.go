package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgedocumentversionvariation
type Knowledgedocumentversionvariation struct { 
	// Id - The globally unique identifier for the variation.
	Id *string `json:"id,omitempty"`


	// Body - The content for the variation.
	Body *Documentbody `json:"body,omitempty"`


	// DateCreated - The creation date-time for the document variation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The last modification date-time for the document variation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Contexts - The context values associated with the variation.
	Contexts *[]Documentvariationcontext `json:"contexts,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// DocumentVersion - Reference to the document version to which the variation is associated with.
	DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`

}

func (o *Knowledgedocumentversionvariation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgedocumentversionvariation
	
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
		
		Body *Documentbody `json:"body,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Contexts *[]Documentvariationcontext `json:"contexts,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DocumentVersion *Addressableentityref `json:"documentVersion,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Body: o.Body,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Contexts: o.Contexts,
		
		SelfUri: o.SelfUri,
		
		DocumentVersion: o.DocumentVersion,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgedocumentversionvariation) UnmarshalJSON(b []byte) error {
	var KnowledgedocumentversionvariationMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgedocumentversionvariationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgedocumentversionvariationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Body, ok := KnowledgedocumentversionvariationMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	
	if dateCreatedString, ok := KnowledgedocumentversionvariationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgedocumentversionvariationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Contexts, ok := KnowledgedocumentversionvariationMap["contexts"].([]interface{}); ok {
		ContextsString, _ := json.Marshal(Contexts)
		json.Unmarshal(ContextsString, &o.Contexts)
	}
	
	if SelfUri, ok := KnowledgedocumentversionvariationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if DocumentVersion, ok := KnowledgedocumentversionvariationMap["documentVersion"].(map[string]interface{}); ok {
		DocumentVersionString, _ := json.Marshal(DocumentVersion)
		json.Unmarshal(DocumentVersionString, &o.DocumentVersion)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionvariation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
