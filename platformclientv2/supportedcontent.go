package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportedcontent - Supported content profile for inbound and outbound messages
type Supportedcontent struct { 
	// Id - A unique supported content Id.
	Id *string `json:"id,omitempty"`


	// Name - The name of the supported content profile
	Name *string `json:"name,omitempty"`


	// DateCreated - Date this supported content profile was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this supported content profile was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this supported content profile
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that modified this supported content profile
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number
	Version *int `json:"version,omitempty"`


	// MediaTypes - Defines the allowable media that may be accepted for an inbound message or to be sent in an outbound message. The following is an example of allowing all inbound media, and for outbound all images and only mpeg video: {   \"mediaTypes\": {     \"allow\": {       \"inbound\": [{\"type\": \"*/*\"}],       \"outbound\": [{\"type\": \"image/*\"}, {\"type\": \"video/mpeg\"}]     }   } }
	MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Supportedcontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportedcontent
	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		MediaTypes: o.MediaTypes,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportedcontent) UnmarshalJSON(b []byte) error {
	var SupportedcontentMap map[string]interface{}
	err := json.Unmarshal(b, &SupportedcontentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SupportedcontentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SupportedcontentMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := SupportedcontentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SupportedcontentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := SupportedcontentMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := SupportedcontentMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := SupportedcontentMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MediaTypes, ok := SupportedcontentMap["mediaTypes"].(map[string]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	
	if SelfUri, ok := SupportedcontentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportedcontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
