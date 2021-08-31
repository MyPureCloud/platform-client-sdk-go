package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facebookintegrationrequest
type Facebookintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the Facebook Integration
	Name *string `json:"name,omitempty"`


	// PageAccessToken - The long-lived Page Access Token of a facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When a pageAccessToken is provided, pageId and userAccessToken are not required.
	PageAccessToken *string `json:"pageAccessToken,omitempty"`


	// UserAccessToken - The short-lived User Access Token of the facebook user logged into the facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When userAccessToken is provided, pageId is mandatory.  When userAccessToken/pageId combination is provided, pageAccessToken is not required.
	UserAccessToken *string `json:"userAccessToken,omitempty"`


	// PageId - The page Id of a facebook page. The pageId is required when userAccessToken is provided.
	PageId *string `json:"pageId,omitempty"`


	// AppId - The app Id of a facebook app. The appId is required when a customer wants to use their own approved facebook app.
	AppId *string `json:"appId,omitempty"`


	// AppSecret - The app Secret of a facebook app. The appSecret is required when appId is provided.
	AppSecret *string `json:"appSecret,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Facebookintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facebookintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		PageAccessToken *string `json:"pageAccessToken,omitempty"`
		
		UserAccessToken *string `json:"userAccessToken,omitempty"`
		
		PageId *string `json:"pageId,omitempty"`
		
		AppId *string `json:"appId,omitempty"`
		
		AppSecret *string `json:"appSecret,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		PageAccessToken: o.PageAccessToken,
		
		UserAccessToken: o.UserAccessToken,
		
		PageId: o.PageId,
		
		AppId: o.AppId,
		
		AppSecret: o.AppSecret,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Facebookintegrationrequest) UnmarshalJSON(b []byte) error {
	var FacebookintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &FacebookintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FacebookintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := FacebookintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if PageAccessToken, ok := FacebookintegrationrequestMap["pageAccessToken"].(string); ok {
		o.PageAccessToken = &PageAccessToken
	}
	
	if UserAccessToken, ok := FacebookintegrationrequestMap["userAccessToken"].(string); ok {
		o.UserAccessToken = &UserAccessToken
	}
	
	if PageId, ok := FacebookintegrationrequestMap["pageId"].(string); ok {
		o.PageId = &PageId
	}
	
	if AppId, ok := FacebookintegrationrequestMap["appId"].(string); ok {
		o.AppId = &AppId
	}
	
	if AppSecret, ok := FacebookintegrationrequestMap["appSecret"].(string); ok {
		o.AppSecret = &AppSecret
	}
	
	if SelfUri, ok := FacebookintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facebookintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
