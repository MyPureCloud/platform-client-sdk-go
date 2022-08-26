package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Categoryresponse
type Categoryresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the category.
	Name *string `json:"name,omitempty"`


	// Description - The description for the category.
	Description *string `json:"description,omitempty"`


	// DateCreated - The creation date-time for the category. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The last modification date-time for the category. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ParentCategory - The reference to category to which this category belongs to.
	ParentCategory **Categoryreference `json:"parentCategory,omitempty"`


	// DocumentCount - Number of documents assigned to this category.
	DocumentCount *int `json:"documentCount,omitempty"`


	// KnowledgeBase - The reference to knowledge base to which the category belongs to.
	KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Categoryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Categoryresponse
	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ParentCategory **Categoryreference `json:"parentCategory,omitempty"`
		
		DocumentCount *int `json:"documentCount,omitempty"`
		
		KnowledgeBase *Knowledgebasereference `json:"knowledgeBase,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ParentCategory: o.ParentCategory,
		
		DocumentCount: o.DocumentCount,
		
		KnowledgeBase: o.KnowledgeBase,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Categoryresponse) UnmarshalJSON(b []byte) error {
	var CategoryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CategoryresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CategoryresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CategoryresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CategoryresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := CategoryresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := CategoryresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ParentCategory, ok := CategoryresponseMap["parentCategory"].(map[string]interface{}); ok {
		ParentCategoryString, _ := json.Marshal(ParentCategory)
		json.Unmarshal(ParentCategoryString, &o.ParentCategory)
	}
	
	if DocumentCount, ok := CategoryresponseMap["documentCount"].(float64); ok {
		DocumentCountInt := int(DocumentCount)
		o.DocumentCount = &DocumentCountInt
	}
	
	if KnowledgeBase, ok := CategoryresponseMap["knowledgeBase"].(map[string]interface{}); ok {
		KnowledgeBaseString, _ := json.Marshal(KnowledgeBase)
		json.Unmarshal(KnowledgeBaseString, &o.KnowledgeBase)
	}
	
	if SelfUri, ok := CategoryresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Categoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
