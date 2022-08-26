package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Categoryreference
type Categoryreference struct { 
	// Id - The globally unique identifier for the category.
	Id *string `json:"id,omitempty"`


	// Name - Category name.
	Name *string `json:"name,omitempty"`


	// ParentCategory - The reference to category to which this category belongs to.
	ParentCategory **Categoryreference `json:"parentCategory,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Categoryreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Categoryreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ParentCategory **Categoryreference `json:"parentCategory,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ParentCategory: o.ParentCategory,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Categoryreference) UnmarshalJSON(b []byte) error {
	var CategoryreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &CategoryreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CategoryreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CategoryreferenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ParentCategory, ok := CategoryreferenceMap["parentCategory"].(map[string]interface{}); ok {
		ParentCategoryString, _ := json.Marshal(ParentCategory)
		json.Unmarshal(ParentCategoryString, &o.ParentCategory)
	}
	
	if SelfUri, ok := CategoryreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Categoryreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
