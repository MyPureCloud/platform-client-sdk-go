package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Line
type Line struct { 
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


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// EdgeGroup
	EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`


	// Template
	Template *Domainentityref `json:"template,omitempty"`


	// Site
	Site *Domainentityref `json:"site,omitempty"`


	// LineBaseSettings
	LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`


	// PrimaryEdge - The primary edge associated to the line. (Deprecated)
	PrimaryEdge *Edge `json:"primaryEdge,omitempty"`


	// SecondaryEdge - The secondary edge associated to the line. (Deprecated)
	SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`


	// LoggedInUser
	LoggedInUser *Domainentityref `json:"loggedInUser,omitempty"`


	// DefaultForUser
	DefaultForUser *Domainentityref `json:"defaultForUser,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Line) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Line
	
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
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		EdgeGroup *Domainentityref `json:"edgeGroup,omitempty"`
		
		Template *Domainentityref `json:"template,omitempty"`
		
		Site *Domainentityref `json:"site,omitempty"`
		
		LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`
		
		PrimaryEdge *Edge `json:"primaryEdge,omitempty"`
		
		SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`
		
		LoggedInUser *Domainentityref `json:"loggedInUser,omitempty"`
		
		DefaultForUser *Domainentityref `json:"defaultForUser,omitempty"`
		
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
		
		Properties: o.Properties,
		
		EdgeGroup: o.EdgeGroup,
		
		Template: o.Template,
		
		Site: o.Site,
		
		LineBaseSettings: o.LineBaseSettings,
		
		PrimaryEdge: o.PrimaryEdge,
		
		SecondaryEdge: o.SecondaryEdge,
		
		LoggedInUser: o.LoggedInUser,
		
		DefaultForUser: o.DefaultForUser,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Line) UnmarshalJSON(b []byte) error {
	var LineMap map[string]interface{}
	err := json.Unmarshal(b, &LineMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LineMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LineMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := LineMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := LineMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Version, ok := LineMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := LineMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := LineMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := LineMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
    
	if CreatedBy, ok := LineMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
    
	if State, ok := LineMap["state"].(string); ok {
		o.State = &State
	}
    
	if ModifiedByApp, ok := LineMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
    
	if CreatedByApp, ok := LineMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
    
	if Properties, ok := LineMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if EdgeGroup, ok := LineMap["edgeGroup"].(map[string]interface{}); ok {
		EdgeGroupString, _ := json.Marshal(EdgeGroup)
		json.Unmarshal(EdgeGroupString, &o.EdgeGroup)
	}
	
	if Template, ok := LineMap["template"].(map[string]interface{}); ok {
		TemplateString, _ := json.Marshal(Template)
		json.Unmarshal(TemplateString, &o.Template)
	}
	
	if Site, ok := LineMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if LineBaseSettings, ok := LineMap["lineBaseSettings"].(map[string]interface{}); ok {
		LineBaseSettingsString, _ := json.Marshal(LineBaseSettings)
		json.Unmarshal(LineBaseSettingsString, &o.LineBaseSettings)
	}
	
	if PrimaryEdge, ok := LineMap["primaryEdge"].(map[string]interface{}); ok {
		PrimaryEdgeString, _ := json.Marshal(PrimaryEdge)
		json.Unmarshal(PrimaryEdgeString, &o.PrimaryEdge)
	}
	
	if SecondaryEdge, ok := LineMap["secondaryEdge"].(map[string]interface{}); ok {
		SecondaryEdgeString, _ := json.Marshal(SecondaryEdge)
		json.Unmarshal(SecondaryEdgeString, &o.SecondaryEdge)
	}
	
	if LoggedInUser, ok := LineMap["loggedInUser"].(map[string]interface{}); ok {
		LoggedInUserString, _ := json.Marshal(LoggedInUser)
		json.Unmarshal(LoggedInUserString, &o.LoggedInUser)
	}
	
	if DefaultForUser, ok := LineMap["defaultForUser"].(map[string]interface{}); ok {
		DefaultForUserString, _ := json.Marshal(DefaultForUser)
		json.Unmarshal(DefaultForUserString, &o.DefaultForUser)
	}
	
	if SelfUri, ok := LineMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Line) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
