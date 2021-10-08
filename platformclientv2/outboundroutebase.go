package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outboundroutebase
type Outboundroutebase struct { 
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


	// ClassificationTypes - The site associated to the outbound route.
	ClassificationTypes *[]string `json:"classificationTypes,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// Distribution
	Distribution *string `json:"distribution,omitempty"`


	// ExternalTrunkBases - Trunk base settings of trunkType \"EXTERNAL\".  This base must also be set on an edge logical interface for correct routing.
	ExternalTrunkBases *[]Domainentityref `json:"externalTrunkBases,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Outboundroutebase) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outboundroutebase
	
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
		
		ClassificationTypes *[]string `json:"classificationTypes,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Distribution *string `json:"distribution,omitempty"`
		
		ExternalTrunkBases *[]Domainentityref `json:"externalTrunkBases,omitempty"`
		
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
		
		ClassificationTypes: o.ClassificationTypes,
		
		Enabled: o.Enabled,
		
		Distribution: o.Distribution,
		
		ExternalTrunkBases: o.ExternalTrunkBases,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Outboundroutebase) UnmarshalJSON(b []byte) error {
	var OutboundroutebaseMap map[string]interface{}
	err := json.Unmarshal(b, &OutboundroutebaseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OutboundroutebaseMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := OutboundroutebaseMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := OutboundroutebaseMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := OutboundroutebaseMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := OutboundroutebaseMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := OutboundroutebaseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := OutboundroutebaseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := OutboundroutebaseMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := OutboundroutebaseMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := OutboundroutebaseMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := OutboundroutebaseMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := OutboundroutebaseMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if ClassificationTypes, ok := OutboundroutebaseMap["classificationTypes"].([]interface{}); ok {
		ClassificationTypesString, _ := json.Marshal(ClassificationTypes)
		json.Unmarshal(ClassificationTypesString, &o.ClassificationTypes)
	}
	
	if Enabled, ok := OutboundroutebaseMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if Distribution, ok := OutboundroutebaseMap["distribution"].(string); ok {
		o.Distribution = &Distribution
	}
	
	if ExternalTrunkBases, ok := OutboundroutebaseMap["externalTrunkBases"].([]interface{}); ok {
		ExternalTrunkBasesString, _ := json.Marshal(ExternalTrunkBases)
		json.Unmarshal(ExternalTrunkBasesString, &o.ExternalTrunkBases)
	}
	
	if SelfUri, ok := OutboundroutebaseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outboundroutebase) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
