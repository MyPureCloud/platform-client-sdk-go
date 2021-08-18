package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Siteconnection) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Siteconnection

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		MediaModel *string `json:"mediaModel,omitempty"`
		
		EdgeList *[]Connectededge `json:"edgeList,omitempty"`
		
		CoreSite *bool `json:"coreSite,omitempty"`
		
		PrimaryCoreSites *[]Domainentityref `json:"primaryCoreSites,omitempty"`
		
		SecondaryCoreSites *[]Domainentityref `json:"secondaryCoreSites,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SelfUri: u.SelfUri,
		
		Managed: u.Managed,
		
		VarType: u.VarType,
		
		Enabled: u.Enabled,
		
		MediaModel: u.MediaModel,
		
		EdgeList: u.EdgeList,
		
		CoreSite: u.CoreSite,
		
		PrimaryCoreSites: u.PrimaryCoreSites,
		
		SecondaryCoreSites: u.SecondaryCoreSites,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Siteconnection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
