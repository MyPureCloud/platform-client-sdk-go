package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phone
type Phone struct { 
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


	// Site - The site associated to the phone.
	Site *Domainentityref `json:"site,omitempty"`


	// PhoneBaseSettings - Phone Base Settings
	PhoneBaseSettings *Domainentityref `json:"phoneBaseSettings,omitempty"`


	// LineBaseSettings
	LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`


	// PhoneMetaBase
	PhoneMetaBase *Domainentityref `json:"phoneMetaBase,omitempty"`


	// Lines - Lines
	Lines *[]Line `json:"lines,omitempty"`


	// Status - The status of the phone and lines from the primary Edge.
	Status *Phonestatus `json:"status,omitempty"`


	// SecondaryStatus - The status of the phone and lines from the secondary Edge.
	SecondaryStatus *Phonestatus `json:"secondaryStatus,omitempty"`


	// UserAgentInfo - User Agent Information for this phone. This includes model, firmware version, and manufacturer.
	UserAgentInfo *Useragentinfo `json:"userAgentInfo,omitempty"`


	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// Capabilities
	Capabilities *Phonecapabilities `json:"capabilities,omitempty"`


	// WebRtcUser - This is the user associated with a WebRTC type phone.  It is required for all WebRTC phones.
	WebRtcUser *Domainentityref `json:"webRtcUser,omitempty"`


	// PrimaryEdge
	PrimaryEdge *Edge `json:"primaryEdge,omitempty"`


	// SecondaryEdge
	SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Phone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phone
	
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
		
		Site *Domainentityref `json:"site,omitempty"`
		
		PhoneBaseSettings *Domainentityref `json:"phoneBaseSettings,omitempty"`
		
		LineBaseSettings *Domainentityref `json:"lineBaseSettings,omitempty"`
		
		PhoneMetaBase *Domainentityref `json:"phoneMetaBase,omitempty"`
		
		Lines *[]Line `json:"lines,omitempty"`
		
		Status *Phonestatus `json:"status,omitempty"`
		
		SecondaryStatus *Phonestatus `json:"secondaryStatus,omitempty"`
		
		UserAgentInfo *Useragentinfo `json:"userAgentInfo,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		Capabilities *Phonecapabilities `json:"capabilities,omitempty"`
		
		WebRtcUser *Domainentityref `json:"webRtcUser,omitempty"`
		
		PrimaryEdge *Edge `json:"primaryEdge,omitempty"`
		
		SecondaryEdge *Edge `json:"secondaryEdge,omitempty"`
		
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
		
		Site: o.Site,
		
		PhoneBaseSettings: o.PhoneBaseSettings,
		
		LineBaseSettings: o.LineBaseSettings,
		
		PhoneMetaBase: o.PhoneMetaBase,
		
		Lines: o.Lines,
		
		Status: o.Status,
		
		SecondaryStatus: o.SecondaryStatus,
		
		UserAgentInfo: o.UserAgentInfo,
		
		Properties: o.Properties,
		
		Capabilities: o.Capabilities,
		
		WebRtcUser: o.WebRtcUser,
		
		PrimaryEdge: o.PrimaryEdge,
		
		SecondaryEdge: o.SecondaryEdge,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Phone) UnmarshalJSON(b []byte) error {
	var PhoneMap map[string]interface{}
	err := json.Unmarshal(b, &PhoneMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhoneMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PhoneMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := PhoneMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := PhoneMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := PhoneMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := PhoneMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := PhoneMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := PhoneMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := PhoneMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := PhoneMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := PhoneMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := PhoneMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if Site, ok := PhoneMap["site"].(map[string]interface{}); ok {
		SiteString, _ := json.Marshal(Site)
		json.Unmarshal(SiteString, &o.Site)
	}
	
	if PhoneBaseSettings, ok := PhoneMap["phoneBaseSettings"].(map[string]interface{}); ok {
		PhoneBaseSettingsString, _ := json.Marshal(PhoneBaseSettings)
		json.Unmarshal(PhoneBaseSettingsString, &o.PhoneBaseSettings)
	}
	
	if LineBaseSettings, ok := PhoneMap["lineBaseSettings"].(map[string]interface{}); ok {
		LineBaseSettingsString, _ := json.Marshal(LineBaseSettings)
		json.Unmarshal(LineBaseSettingsString, &o.LineBaseSettings)
	}
	
	if PhoneMetaBase, ok := PhoneMap["phoneMetaBase"].(map[string]interface{}); ok {
		PhoneMetaBaseString, _ := json.Marshal(PhoneMetaBase)
		json.Unmarshal(PhoneMetaBaseString, &o.PhoneMetaBase)
	}
	
	if Lines, ok := PhoneMap["lines"].([]interface{}); ok {
		LinesString, _ := json.Marshal(Lines)
		json.Unmarshal(LinesString, &o.Lines)
	}
	
	if Status, ok := PhoneMap["status"].(map[string]interface{}); ok {
		StatusString, _ := json.Marshal(Status)
		json.Unmarshal(StatusString, &o.Status)
	}
	
	if SecondaryStatus, ok := PhoneMap["secondaryStatus"].(map[string]interface{}); ok {
		SecondaryStatusString, _ := json.Marshal(SecondaryStatus)
		json.Unmarshal(SecondaryStatusString, &o.SecondaryStatus)
	}
	
	if UserAgentInfo, ok := PhoneMap["userAgentInfo"].(map[string]interface{}); ok {
		UserAgentInfoString, _ := json.Marshal(UserAgentInfo)
		json.Unmarshal(UserAgentInfoString, &o.UserAgentInfo)
	}
	
	if Properties, ok := PhoneMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Capabilities, ok := PhoneMap["capabilities"].(map[string]interface{}); ok {
		CapabilitiesString, _ := json.Marshal(Capabilities)
		json.Unmarshal(CapabilitiesString, &o.Capabilities)
	}
	
	if WebRtcUser, ok := PhoneMap["webRtcUser"].(map[string]interface{}); ok {
		WebRtcUserString, _ := json.Marshal(WebRtcUser)
		json.Unmarshal(WebRtcUserString, &o.WebRtcUser)
	}
	
	if PrimaryEdge, ok := PhoneMap["primaryEdge"].(map[string]interface{}); ok {
		PrimaryEdgeString, _ := json.Marshal(PrimaryEdge)
		json.Unmarshal(PrimaryEdgeString, &o.PrimaryEdge)
	}
	
	if SecondaryEdge, ok := PhoneMap["secondaryEdge"].(map[string]interface{}); ok {
		SecondaryEdgeString, _ := json.Marshal(SecondaryEdge)
		json.Unmarshal(SecondaryEdgeString, &o.SecondaryEdge)
	}
	
	if SelfUri, ok := PhoneMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Phone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
