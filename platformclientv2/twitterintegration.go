package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Twitterintegration
type Twitterintegration struct { 
	// Id - A unique Integration Id
	Id *string `json:"id,omitempty"`


	// Name - The name of the Twitter Integration
	Name *string `json:"name,omitempty"`


	// AccessTokenKey - The Access Token Key from Twitter messenger
	AccessTokenKey *string `json:"accessTokenKey,omitempty"`


	// ConsumerKey - The Consumer Key from Twitter messenger
	ConsumerKey *string `json:"consumerKey,omitempty"`


	// Username - The Username from Twitter
	Username *string `json:"username,omitempty"`


	// UserId - The UserId from Twitter
	UserId *string `json:"userId,omitempty"`


	// Status - The status of the Twitter Integration
	Status *string `json:"status,omitempty"`


	// Tier - The type of twitter account to be used for the integration
	Tier *string `json:"tier,omitempty"`


	// EnvName - The Twitter environment name, e.g.: env-beta (required for premium tier)
	EnvName *string `json:"envName,omitempty"`


	// Recipient - The recipient associated to the Twitter Integration. This recipient is used to associate a flow to an integration
	Recipient *Domainentityref `json:"recipient,omitempty"`


	// DateCreated - Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// CreatedBy - User reference that created this Integration
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// ModifiedBy - User reference that last modified this Integration
	ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`


	// Version - Version number required for updates.
	Version *int `json:"version,omitempty"`


	// CreateStatus - Status of asynchronous create operation
	CreateStatus *string `json:"createStatus,omitempty"`


	// CreateError - Error information returned, if createStatus is set to Error
	CreateError *Errorbody `json:"createError,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Twitterintegration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Twitterintegration

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AccessTokenKey *string `json:"accessTokenKey,omitempty"`
		
		ConsumerKey *string `json:"consumerKey,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Tier *string `json:"tier,omitempty"`
		
		EnvName *string `json:"envName,omitempty"`
		
		Recipient *Domainentityref `json:"recipient,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		ModifiedBy *Domainentityref `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		CreateStatus *string `json:"createStatus,omitempty"`
		
		CreateError *Errorbody `json:"createError,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		AccessTokenKey: u.AccessTokenKey,
		
		ConsumerKey: u.ConsumerKey,
		
		Username: u.Username,
		
		UserId: u.UserId,
		
		Status: u.Status,
		
		Tier: u.Tier,
		
		EnvName: u.EnvName,
		
		Recipient: u.Recipient,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		CreatedBy: u.CreatedBy,
		
		ModifiedBy: u.ModifiedBy,
		
		Version: u.Version,
		
		CreateStatus: u.CreateStatus,
		
		CreateError: u.CreateError,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Twitterintegration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
