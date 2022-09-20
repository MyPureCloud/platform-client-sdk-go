package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactiontemplate
type Patchactiontemplate struct { 
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
	ContentOffer *Patchcontentoffer `json:"contentOffer,omitempty"`

}

func (o *Patchactiontemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchactiontemplate
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ContentOffer *Patchcontentoffer `json:"contentOffer,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		MediaType: o.MediaType,
		
		State: o.State,
		
		ContentOffer: o.ContentOffer,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchactiontemplate) UnmarshalJSON(b []byte) error {
	var PatchactiontemplateMap map[string]interface{}
	err := json.Unmarshal(b, &PatchactiontemplateMap)
	if err != nil {
		return err
	}
	
	if Name, ok := PatchactiontemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := PatchactiontemplateMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := PatchactiontemplateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if MediaType, ok := PatchactiontemplateMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if State, ok := PatchactiontemplateMap["state"].(string); ok {
		o.State = &State
	}
    
	if ContentOffer, ok := PatchactiontemplateMap["contentOffer"].(map[string]interface{}); ok {
		ContentOfferString, _ := json.Marshal(ContentOffer)
		json.Unmarshal(ContentOfferString, &o.ContentOffer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchactiontemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
