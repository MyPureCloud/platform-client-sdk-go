package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


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

func (o *Site) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Site
	
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
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		PrimarySites *[]Domainentityref `json:"primarySites,omitempty"`
		
		SecondarySites *[]Domainentityref `json:"secondarySites,omitempty"`
		
		PrimaryEdges *[]Edge `json:"primaryEdges,omitempty"`
		
		SecondaryEdges *[]Edge `json:"secondaryEdges,omitempty"`
		
		Addresses *[]Contact `json:"addresses,omitempty"`
		
		Edges *[]Edge `json:"edges,omitempty"`
		
		EdgeAutoUpdateConfig *Edgeautoupdateconfig `json:"edgeAutoUpdateConfig,omitempty"`
		
		MediaRegionsUseLatencyBased *bool `json:"mediaRegionsUseLatencyBased,omitempty"`
		
		Location *Locationdefinition `json:"location,omitempty"`
		
		Managed *bool `json:"managed,omitempty"`
		
		NtpSettings *Ntpsettings `json:"ntpSettings,omitempty"`
		
		MediaModel *string `json:"mediaModel,omitempty"`
		
		CoreSite *bool `json:"coreSite,omitempty"`
		
		SiteConnections *[]Siteconnection `json:"siteConnections,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		PrimarySites: o.PrimarySites,
		
		SecondarySites: o.SecondarySites,
		
		PrimaryEdges: o.PrimaryEdges,
		
		SecondaryEdges: o.SecondaryEdges,
		
		Addresses: o.Addresses,
		
		Edges: o.Edges,
		
		EdgeAutoUpdateConfig: o.EdgeAutoUpdateConfig,
		
		MediaRegionsUseLatencyBased: o.MediaRegionsUseLatencyBased,
		
		Location: o.Location,
		
		Managed: o.Managed,
		
		NtpSettings: o.NtpSettings,
		
		MediaModel: o.MediaModel,
		
		CoreSite: o.CoreSite,
		
		SiteConnections: o.SiteConnections,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Site) UnmarshalJSON(b []byte) error {
	var SiteMap map[string]interface{}
	err := json.Unmarshal(b, &SiteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SiteMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SiteMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := SiteMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := SiteMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := SiteMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := SiteMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := SiteMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := SiteMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := SiteMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := SiteMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := SiteMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := SiteMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if PrimarySites, ok := SiteMap["primarySites"].([]interface{}); ok {
		PrimarySitesString, _ := json.Marshal(PrimarySites)
		json.Unmarshal(PrimarySitesString, &o.PrimarySites)
	}
	
	if SecondarySites, ok := SiteMap["secondarySites"].([]interface{}); ok {
		SecondarySitesString, _ := json.Marshal(SecondarySites)
		json.Unmarshal(SecondarySitesString, &o.SecondarySites)
	}
	
	if PrimaryEdges, ok := SiteMap["primaryEdges"].([]interface{}); ok {
		PrimaryEdgesString, _ := json.Marshal(PrimaryEdges)
		json.Unmarshal(PrimaryEdgesString, &o.PrimaryEdges)
	}
	
	if SecondaryEdges, ok := SiteMap["secondaryEdges"].([]interface{}); ok {
		SecondaryEdgesString, _ := json.Marshal(SecondaryEdges)
		json.Unmarshal(SecondaryEdgesString, &o.SecondaryEdges)
	}
	
	if Addresses, ok := SiteMap["addresses"].([]interface{}); ok {
		AddressesString, _ := json.Marshal(Addresses)
		json.Unmarshal(AddressesString, &o.Addresses)
	}
	
	if Edges, ok := SiteMap["edges"].([]interface{}); ok {
		EdgesString, _ := json.Marshal(Edges)
		json.Unmarshal(EdgesString, &o.Edges)
	}
	
	if EdgeAutoUpdateConfig, ok := SiteMap["edgeAutoUpdateConfig"].(map[string]interface{}); ok {
		EdgeAutoUpdateConfigString, _ := json.Marshal(EdgeAutoUpdateConfig)
		json.Unmarshal(EdgeAutoUpdateConfigString, &o.EdgeAutoUpdateConfig)
	}
	
	if MediaRegionsUseLatencyBased, ok := SiteMap["mediaRegionsUseLatencyBased"].(bool); ok {
		o.MediaRegionsUseLatencyBased = &MediaRegionsUseLatencyBased
	}
    
	if Location, ok := SiteMap["location"].(map[string]interface{}); ok {
		LocationString, _ := json.Marshal(Location)
		json.Unmarshal(LocationString, &o.Location)
	}
	
	if Managed, ok := SiteMap["managed"].(bool); ok {
		o.Managed = &Managed
	}
    
	if NtpSettings, ok := SiteMap["ntpSettings"].(map[string]interface{}); ok {
		NtpSettingsString, _ := json.Marshal(NtpSettings)
		json.Unmarshal(NtpSettingsString, &o.NtpSettings)
	}
	
	if MediaModel, ok := SiteMap["mediaModel"].(string); ok {
		o.MediaModel = &MediaModel
	}
    
	if CoreSite, ok := SiteMap["coreSite"].(bool); ok {
		o.CoreSite = &CoreSite
	}
    
	if SiteConnections, ok := SiteMap["siteConnections"].([]interface{}); ok {
		SiteConnectionsString, _ := json.Marshal(SiteConnections)
		json.Unmarshal(SiteConnectionsString, &o.SiteConnections)
	}
	
	if SelfUri, ok := SiteMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Site) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
