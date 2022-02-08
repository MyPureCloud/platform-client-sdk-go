package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookintegrationupdaterequest
type Facebookintegrationupdaterequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Facebook Integration
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// PageAccessToken - The long-lived Page Access Token of Facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
	PageAccessToken *string `json:"pageAccessToken,omitempty"`


	// UserAccessToken - The short-lived User Access Token of the Facebook user logged into the Facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
	UserAccessToken *string `json:"userAccessToken,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Facebookintegrationupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookintegrationupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		PageAccessToken *string `json:"pageAccessToken,omitempty"`
		
		UserAccessToken *string `json:"userAccessToken,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		PageAccessToken: o.PageAccessToken,
		
		UserAccessToken: o.UserAccessToken,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookintegrationupdaterequest) UnmarshalJSON(b []byte) error {
	var FacebookintegrationupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookintegrationupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FacebookintegrationupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FacebookintegrationupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SupportedContent, ok := FacebookintegrationupdaterequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if PageAccessToken, ok := FacebookintegrationupdaterequestMap["pageAccessToken"].(string); ok {
		o.PageAccessToken = &PageAccessToken
	}
	
	if UserAccessToken, ok := FacebookintegrationupdaterequestMap["userAccessToken"].(string); ok {
		o.UserAccessToken = &UserAccessToken
	}
	
	if SelfUri, ok := FacebookintegrationupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
