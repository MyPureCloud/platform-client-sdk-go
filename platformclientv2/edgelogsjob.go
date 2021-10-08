package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgelogsjob
type Edgelogsjob struct { 
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


	// Files - The files available to upload from the Edge to the cloud.
	Files *[]Edgelogsjobfile `json:"files,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Edgelogsjob) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgelogsjob
	
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
		
		Files *[]Edgelogsjobfile `json:"files,omitempty"`
		
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
		
		Files: o.Files,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgelogsjob) UnmarshalJSON(b []byte) error {
	var EdgelogsjobMap map[string]interface{}
	err := json.Unmarshal(b, &EdgelogsjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EdgelogsjobMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := EdgelogsjobMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Division, ok := EdgelogsjobMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := EdgelogsjobMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := EdgelogsjobMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := EdgelogsjobMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := EdgelogsjobMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := EdgelogsjobMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := EdgelogsjobMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := EdgelogsjobMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := EdgelogsjobMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := EdgelogsjobMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if Files, ok := EdgelogsjobMap["files"].([]interface{}); ok {
		FilesString, _ := json.Marshal(Files)
		json.Unmarshal(FilesString, &o.Files)
	}
	
	if SelfUri, ok := EdgelogsjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgelogsjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
