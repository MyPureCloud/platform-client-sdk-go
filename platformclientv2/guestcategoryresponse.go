package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Guestcategoryresponse
type Guestcategoryresponse struct { 
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


	// ParentCategory - The reference to category to which this category belongs.
	ParentCategory *Guestcategoryreference `json:"parentCategory,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Guestcategoryresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Guestcategoryresponse
	
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
		
		ParentCategory *Guestcategoryreference `json:"parentCategory,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ParentCategory: o.ParentCategory,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Guestcategoryresponse) UnmarshalJSON(b []byte) error {
	var GuestcategoryresponseMap map[string]interface{}
	err := json.Unmarshal(b, &GuestcategoryresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GuestcategoryresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := GuestcategoryresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := GuestcategoryresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := GuestcategoryresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := GuestcategoryresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ParentCategory, ok := GuestcategoryresponseMap["parentCategory"].(map[string]interface{}); ok {
		ParentCategoryString, _ := json.Marshal(ParentCategory)
		json.Unmarshal(ParentCategoryString, &o.ParentCategory)
	}
	
	if SelfUri, ok := GuestcategoryresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Guestcategoryresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
