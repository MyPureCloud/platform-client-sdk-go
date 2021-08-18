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

func (u *Actiontemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actiontemplate

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ContentOffer *Contentoffer `json:"contentOffer,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		MediaType: u.MediaType,
		
		State: u.State,
		
		ContentOffer: u.ContentOffer,
		
		SelfUri: u.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actiontemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
