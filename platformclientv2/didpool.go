package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Didpool
type Didpool struct { 
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


	// StartPhoneNumber - The starting phone number for the range of this DID pool. Must be in E.164 format
	StartPhoneNumber *string `json:"startPhoneNumber,omitempty"`


	// EndPhoneNumber - The ending phone number for the range of this DID pool. Must be in E.164 format
	EndPhoneNumber *string `json:"endPhoneNumber,omitempty"`


	// Comments
	Comments *string `json:"comments,omitempty"`


	// Provider - The provider for this DID pool
	Provider *string `json:"provider,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Didpool) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Didpool
	
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
		
		Description *string `json:"description,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *string `json:"modifiedBy,omitempty"`
		
		CreatedBy *string `json:"createdBy,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		ModifiedByApp *string `json:"modifiedByApp,omitempty"`
		
		CreatedByApp *string `json:"createdByApp,omitempty"`
		
		StartPhoneNumber *string `json:"startPhoneNumber,omitempty"`
		
		EndPhoneNumber *string `json:"endPhoneNumber,omitempty"`
		
		Comments *string `json:"comments,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Version: o.Version,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		CreatedBy: o.CreatedBy,
		
		State: o.State,
		
		ModifiedByApp: o.ModifiedByApp,
		
		CreatedByApp: o.CreatedByApp,
		
		StartPhoneNumber: o.StartPhoneNumber,
		
		EndPhoneNumber: o.EndPhoneNumber,
		
		Comments: o.Comments,
		
		Provider: o.Provider,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Didpool) UnmarshalJSON(b []byte) error {
	var DidpoolMap map[string]interface{}
	err := json.Unmarshal(b, &DidpoolMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DidpoolMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DidpoolMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := DidpoolMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := DidpoolMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := DidpoolMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DidpoolMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := DidpoolMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := DidpoolMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := DidpoolMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := DidpoolMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := DidpoolMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if StartPhoneNumber, ok := DidpoolMap["startPhoneNumber"].(string); ok {
		o.StartPhoneNumber = &StartPhoneNumber
	}
	
	if EndPhoneNumber, ok := DidpoolMap["endPhoneNumber"].(string); ok {
		o.EndPhoneNumber = &EndPhoneNumber
	}
	
	if Comments, ok := DidpoolMap["comments"].(string); ok {
		o.Comments = &Comments
	}
	
	if Provider, ok := DidpoolMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if SelfUri, ok := DidpoolMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Didpool) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
