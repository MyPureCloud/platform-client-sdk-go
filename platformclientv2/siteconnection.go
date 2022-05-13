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

func (o *Siteconnection) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		Managed: o.Managed,
		
		VarType: o.VarType,
		
		Enabled: o.Enabled,
		
		MediaModel: o.MediaModel,
		
		EdgeList: o.EdgeList,
		
		CoreSite: o.CoreSite,
		
		PrimaryCoreSites: o.PrimaryCoreSites,
		
		SecondaryCoreSites: o.SecondaryCoreSites,
		Alias:    (*Alias)(o),
	})
}

func (o *Siteconnection) UnmarshalJSON(b []byte) error {
	var SiteconnectionMap map[string]interface{}
	err := json.Unmarshal(b, &SiteconnectionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SiteconnectionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SiteconnectionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := SiteconnectionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Managed, ok := SiteconnectionMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if VarType, ok := SiteconnectionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Enabled, ok := SiteconnectionMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if MediaModel, ok := SiteconnectionMap["mediaModel"].(string); ok {
		o.MediaModel = &MediaModel
	}
    
	if EdgeList, ok := SiteconnectionMap["edgeList"].([]interface{}); ok {
		EdgeListString, _ := json.Marshal(EdgeList)
		json.Unmarshal(EdgeListString, &o.EdgeList)
	}
	
	if CoreSite, ok := SiteconnectionMap["coreSite"].(bool); ok {
		o.CoreSite = &CoreSite
	}
    
	if PrimaryCoreSites, ok := SiteconnectionMap["primaryCoreSites"].([]interface{}); ok {
		PrimaryCoreSitesString, _ := json.Marshal(PrimaryCoreSites)
		json.Unmarshal(PrimaryCoreSitesString, &o.PrimaryCoreSites)
	}
	
	if SecondaryCoreSites, ok := SiteconnectionMap["secondaryCoreSites"].([]interface{}); ok {
		SecondaryCoreSitesString, _ := json.Marshal(SecondaryCoreSites)
		json.Unmarshal(SecondaryCoreSitesString, &o.SecondaryCoreSites)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Siteconnection) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
