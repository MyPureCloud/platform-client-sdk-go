package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actiontemplate
type Actiontemplate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name of the action template.
	Name *string `json:"name,omitempty"`


	// Description - Description of the action template's functionality.
	Description *string `json:"description,omitempty"`


	// Version - The version of the action template.
	Version *int `json:"version,omitempty"`


	// MediaType - Media type of action described by the action template.
	MediaType *string `json:"mediaType,omitempty"`


	// State - Whether the action template is currently active, inactive or deleted.
	State *string `json:"state,omitempty"`


	// ContentOffer - Properties used to configure an action of type content offer
	ContentOffer *Contentoffer `json:"contentOffer,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// CreatedDate - Date when action template was created in ISO-8601 format.
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date when action template was last modified in ISO-8601 format.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

}

func (o *Actiontemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actiontemplate
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ContentOffer *Contentoffer `json:"contentOffer,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		MediaType: o.MediaType,
		
		State: o.State,
		
		ContentOffer: o.ContentOffer,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Actiontemplate) UnmarshalJSON(b []byte) error {
	var ActiontemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ActiontemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActiontemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActiontemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ActiontemplateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := ActiontemplateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MediaType, ok := ActiontemplateMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if State, ok := ActiontemplateMap["state"].(string); ok {
		o.State = &State
	}
    
	if ContentOffer, ok := ActiontemplateMap["contentOffer"].(map[string]interface{}); ok {
		ContentOfferString, _ := json.Marshal(ContentOffer)
		json.Unmarshal(ContentOfferString, &o.ContentOffer)
	}
	
	if SelfUri, ok := ActiontemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := ActiontemplateMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := ActiontemplateMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actiontemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
