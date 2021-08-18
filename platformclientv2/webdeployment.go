package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeployment - Details about a Web Deployment
type Webdeployment struct { 
	// Id - The deployment ID
	Id *string `json:"id,omitempty"`


	// Name - The deployment name
	Name *string `json:"name,omitempty"`


	// Description - The description of the config
	Description *string `json:"description,omitempty"`


	// Configuration - The config version this deployment uses
	Configuration *Webdeploymentconfigurationversionentityref `json:"configuration,omitempty"`


	// AllowAllDomains - Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`


	// AllowedDomains - The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`


	// Snippet - Javascript snippet used to load the config
	Snippet *string `json:"snippet,omitempty"`


	// DateCreated - The date the deployment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date the deployment was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// LastModifiedUser - A reference to the user who most recently modified the deployment
	LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`


	// Flow - A reference to the inboundshortmessage flow used by this deployment
	Flow *Domainentityref `json:"flow,omitempty"`


	// Status - The current status of the deployment
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Webdeployment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeployment

	
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
		
		Description *string `json:"description,omitempty"`
		
		Configuration *Webdeploymentconfigurationversionentityref `json:"configuration,omitempty"`
		
		AllowAllDomains *bool `json:"allowAllDomains,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		Snippet *string `json:"snippet,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Configuration: u.Configuration,
		
		AllowAllDomains: u.AllowAllDomains,
		
		AllowedDomains: u.AllowedDomains,
		
		Snippet: u.Snippet,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		LastModifiedUser: u.LastModifiedUser,
		
		Flow: u.Flow,
		
		Status: u.Status,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Webdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
