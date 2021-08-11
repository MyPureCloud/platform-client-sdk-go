package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Webdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
