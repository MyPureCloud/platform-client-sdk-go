package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Coachingannotation
type Coachingannotation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// CreatedBy - The user who created the annotation.
	CreatedBy *Userreference `json:"createdBy,omitempty"`


	// DateCreated - The date/time the annotation was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy - The last user to modify the annotation.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`


	// DateModified - The date/time the annotation was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Text - The text of the annotation.
	Text *string `json:"text,omitempty"`


	// IsDeleted - Flag indicating whether the annotation is deleted.
	IsDeleted *bool `json:"isDeleted,omitempty"`


	// AccessType - Determines the permissions required to view this item.
	AccessType *string `json:"accessType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Coachingannotation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Coachingannotation
	
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
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		IsDeleted *bool `json:"isDeleted,omitempty"`
		
		AccessType *string `json:"accessType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		Text: o.Text,
		
		IsDeleted: o.IsDeleted,
		
		AccessType: o.AccessType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Coachingannotation) UnmarshalJSON(b []byte) error {
	var CoachingannotationMap map[string]interface{}
	err := json.Unmarshal(b, &CoachingannotationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CoachingannotationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if CreatedBy, ok := CoachingannotationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := CoachingannotationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := CoachingannotationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := CoachingannotationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Text, ok := CoachingannotationMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if IsDeleted, ok := CoachingannotationMap["isDeleted"].(bool); ok {
		o.IsDeleted = &IsDeleted
	}
    
	if AccessType, ok := CoachingannotationMap["accessType"].(string); ok {
		o.AccessType = &AccessType
	}
    
	if SelfUri, ok := CoachingannotationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Coachingannotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
