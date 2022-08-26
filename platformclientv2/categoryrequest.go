package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Categoryrequest
type Categoryrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the category.
	Name *string `json:"name,omitempty"`


	// Description - The description for the category.
	Description *string `json:"description,omitempty"`


	// ParentCategoryId - The category to which this category belongs.
	ParentCategoryId *string `json:"parentCategoryId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Categoryrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Categoryrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ParentCategoryId *string `json:"parentCategoryId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		ParentCategoryId: o.ParentCategoryId,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Categoryrequest) UnmarshalJSON(b []byte) error {
	var CategoryrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CategoryrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CategoryrequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CategoryrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := CategoryrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ParentCategoryId, ok := CategoryrequestMap["parentCategoryId"].(string); ok {
		o.ParentCategoryId = &ParentCategoryId
	}
    
	if SelfUri, ok := CategoryrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Categoryrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
