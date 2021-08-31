package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Attribute
type Attribute struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The attribute name.
	Name *string `json:"name,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// ModifiedBy
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Attribute) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Attribute
	
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
		
		Version *int `json:"version,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Version: o.Version,
		
		Description: o.Description,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		ModifiedBy: o.ModifiedBy,
		
		DateModified: DateModified,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Attribute) UnmarshalJSON(b []byte) error {
	var AttributeMap map[string]interface{}
	err := json.Unmarshal(b, &AttributeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AttributeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := AttributeMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Version, ok := AttributeMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Description, ok := AttributeMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if CreatedBy, ok := AttributeMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := AttributeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if ModifiedBy, ok := AttributeMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if dateModifiedString, ok := AttributeMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if SelfUri, ok := AttributeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Attribute) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
