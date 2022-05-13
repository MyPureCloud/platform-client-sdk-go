package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthauthorization
type Oauthauthorization struct { 
	// Client
	Client *Oauthclient `json:"client,omitempty"`


	// Scope
	Scope *[]string `json:"scope,omitempty"`


	// Roles
	Roles *[]string `json:"roles,omitempty"`


	// ResourceOwner
	ResourceOwner *Domainentityref `json:"resourceOwner,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Pending
	Pending *bool `json:"pending,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Oauthauthorization) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthauthorization
	
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
		Client *Oauthclient `json:"client,omitempty"`
		
		Scope *[]string `json:"scope,omitempty"`
		
		Roles *[]string `json:"roles,omitempty"`
		
		ResourceOwner *Domainentityref `json:"resourceOwner,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Pending *bool `json:"pending,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Client: o.Client,
		
		Scope: o.Scope,
		
		Roles: o.Roles,
		
		ResourceOwner: o.ResourceOwner,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedBy: o.ModifiedBy,
		
		Pending: o.Pending,
		
		State: o.State,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Oauthauthorization) UnmarshalJSON(b []byte) error {
	var OauthauthorizationMap map[string]interface{}
	err := json.Unmarshal(b, &OauthauthorizationMap)
	if err != nil {
		return err
	}
	
	if Client, ok := OauthauthorizationMap["client"].(map[string]interface{}); ok {
		ClientString, _ := json.Marshal(Client)
		json.Unmarshal(ClientString, &o.Client)
	}
	
	if Scope, ok := OauthauthorizationMap["scope"].([]interface{}); ok {
		ScopeString, _ := json.Marshal(Scope)
		json.Unmarshal(ScopeString, &o.Scope)
	}
	
	if Roles, ok := OauthauthorizationMap["roles"].([]interface{}); ok {
		RolesString, _ := json.Marshal(Roles)
		json.Unmarshal(RolesString, &o.Roles)
	}
	
	if ResourceOwner, ok := OauthauthorizationMap["resourceOwner"].(map[string]interface{}); ok {
		ResourceOwnerString, _ := json.Marshal(ResourceOwner)
		json.Unmarshal(ResourceOwnerString, &o.ResourceOwner)
	}
	
	if dateCreatedString, ok := OauthauthorizationMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OauthauthorizationMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if CreatedBy, ok := OauthauthorizationMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ModifiedBy, ok := OauthauthorizationMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Pending, ok := OauthauthorizationMap["pending"].(bool); ok {
		o.Pending = &Pending
	}
    
	if State, ok := OauthauthorizationMap["state"].(string); ok {
		o.State = &State
	}
    
	if SelfUri, ok := OauthauthorizationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthauthorization) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
