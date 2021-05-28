package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Siteconnection
type Siteconnection struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// Managed
	Managed *bool `json:"managed,omitempty"`


	// VarType - Connection method from site to site (Direct, Indirect, CloudProxy
	VarType *string `json:"type,omitempty"`


	// Enabled - Indicates if the current site is linked
	Enabled *bool `json:"enabled,omitempty"`


	// MediaModel - Media model for the current site.
	MediaModel *string `json:"mediaModel,omitempty"`


	// EdgeList - All of the edges to which the site connects
	EdgeList *[]Connectededge `json:"edgeList,omitempty"`


	// CoreSite - The core site
	CoreSite *bool `json:"coreSite,omitempty"`


	// PrimaryCoreSites - List of site ids names and selfUris for the primary core sites
	PrimaryCoreSites *[]Domainentityref `json:"primaryCoreSites,omitempty"`


	// SecondaryCoreSites - List of site ids names and selfUris for the secondary core sites
	SecondaryCoreSites *[]Domainentityref `json:"secondaryCoreSites,omitempty"`

}

// String returns a JSON representation of the model
func (o *Siteconnection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
