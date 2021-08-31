package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgeextendedcategory
type Knowledgeextendedcategory struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Category name
	Name *string `json:"name,omitempty"`


	// Description - Category description
	Description *string `json:"description,omitempty"`


	// KnowledgeBase - Knowledge base which category does belong to
	KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`


	// LanguageCode - Actual language of the category
	LanguageCode *string `json:"languageCode,omitempty"`


	// DateCreated - Category creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Category last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Parent - Category parent
	Parent *Knowledgecategory `json:"parent,omitempty"`


	// Children - Category children
	Children *[]Knowledgecategory `json:"children,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Knowledgeextendedcategory) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Knowledgeextendedcategory
	
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
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		KnowledgeBase *Knowledgebase `json:"knowledgeBase,omitempty"`
		
		LanguageCode *string `json:"languageCode,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Parent *Knowledgecategory `json:"parent,omitempty"`
		
		Children *[]Knowledgecategory `json:"children,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		KnowledgeBase: o.KnowledgeBase,
		
		LanguageCode: o.LanguageCode,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Parent: o.Parent,
		
		Children: o.Children,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Knowledgeextendedcategory) UnmarshalJSON(b []byte) error {
	var KnowledgeextendedcategoryMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgeextendedcategoryMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgeextendedcategoryMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KnowledgeextendedcategoryMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := KnowledgeextendedcategoryMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if KnowledgeBase, ok := KnowledgeextendedcategoryMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if LanguageCode, ok := KnowledgeextendedcategoryMap["languageCode"].(string); ok {
		o.LanguageCode = &LanguageCode
	}
	
	if dateCreatedString, ok := KnowledgeextendedcategoryMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgeextendedcategoryMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Parent, ok := KnowledgeextendedcategoryMap["parent"].(map[string]interface{}); ok {
		ParentString, _ := json.Marshal(Parent)
		json.Unmarshal(ParentString, &o.Parent)
	}
	
	if Children, ok := KnowledgeextendedcategoryMap["children"].([]interface{}); ok {
		ChildrenString, _ := json.Marshal(Children)
		json.Unmarshal(ChildrenString, &o.Children)
	}
	
	if SelfUri, ok := KnowledgeextendedcategoryMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgeextendedcategory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
