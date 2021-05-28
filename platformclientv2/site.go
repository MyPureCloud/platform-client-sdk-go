package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Site
type Site struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the entity.
	Name *string `json:"name,omitempty"`


	// Description - The resource's description.
	Description *string `json:"description,omitempty"`


	// Version - The current version of the resource.
	Version *int `json:"version,omitempty"`


	// DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ModifiedBy - The ID of the user that last modified the resource.
	ModifiedBy *string `json:"modifiedBy,omitempty"`


	// CreatedBy - The ID of the user that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`


	// State - Indicates if the resource is active, inactive, or deleted.
	State *string `json:"state,omitempty"`


	// ModifiedByApp - The application that last modified the resource.
	ModifiedByApp *string `json:"modifiedByApp,omitempty"`


	// CreatedByApp - The application that created the resource.
	CreatedByApp *string `json:"createdByApp,omitempty"`


	// PrimarySites
	PrimarySites *[]Domainentityref `json:"primarySites,omitempty"`


	// SecondarySites
	SecondarySites *[]Domainentityref `json:"secondarySites,omitempty"`


	// PrimaryEdges
	PrimaryEdges *[]Edge `json:"primaryEdges,omitempty"`


	// SecondaryEdges
	SecondaryEdges *[]Edge `json:"secondaryEdges,omitempty"`


	// Addresses
	Addresses *[]Contact `json:"addresses,omitempty"`


	// Edges
	Edges *[]Edge `json:"edges,omitempty"`


	// EdgeAutoUpdateConfig - Recurrance rule, time zone, and start/end settings for automatic edge updates for this site
	EdgeAutoUpdateConfig *Edgeautoupdateconfig `json:"edgeAutoUpdateConfig,omitempty"`


	// MediaRegionsUseLatencyBased
	MediaRegionsUseLatencyBased *bool `json:"mediaRegionsUseLatencyBased,omitempty"`


	// Location - Location
	Location *Locationdefinition `json:"location,omitempty"`


	// Managed
	Managed *bool `json:"managed,omitempty"`


	// NtpSettings - Network Time Protocol settings for the site
	NtpSettings *Ntpsettings `json:"ntpSettings,omitempty"`


	// MediaModel - Media model for the site
	MediaModel *string `json:"mediaModel,omitempty"`


	// CoreSite - Is this site a core site
	CoreSite *bool `json:"coreSite,omitempty"`


	// SiteConnections - The site connections
	SiteConnections *[]Siteconnection `json:"siteConnections,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Site) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
