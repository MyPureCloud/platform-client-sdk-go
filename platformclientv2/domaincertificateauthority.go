package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domaincertificateauthority - A certificate authority represents an organization that has issued a digital certificate for making secure connections with an edge device.
type Domaincertificateauthority struct { 
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


	// Certificate - The authorities signed X509 PEM encoded certificate.
	Certificate *string `json:"certificate,omitempty"`


	// VarType - The certificate authorities type.  Managed certificate authorities are generated and maintained by Interactive Intelligence.  These are read-only and not modifiable by clients.  Remote authorities are customer managed.
	VarType *string `json:"type,omitempty"`


	// Services - The service(s) that the authority can be used to authenticate.
	Services *[]string `json:"services,omitempty"`


	// CertificateDetails - The details of the parsed certificate(s).
	CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domaincertificateauthority) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domaincertificateauthority
	
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
		
		Certificate *string `json:"certificate,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Services *[]string `json:"services,omitempty"`
		
		CertificateDetails *[]Certificatedetails `json:"certificateDetails,omitempty"`
		
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
		
		Certificate: o.Certificate,
		
		VarType: o.VarType,
		
		Services: o.Services,
		
		CertificateDetails: o.CertificateDetails,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domaincertificateauthority) UnmarshalJSON(b []byte) error {
	var DomaincertificateauthorityMap map[string]interface{}
	err := json.Unmarshal(b, &DomaincertificateauthorityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomaincertificateauthorityMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomaincertificateauthorityMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := DomaincertificateauthorityMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if Version, ok := DomaincertificateauthorityMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if dateCreatedString, ok := DomaincertificateauthorityMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DomaincertificateauthorityMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := DomaincertificateauthorityMap["modifiedBy"].(string); ok {
		o.ModifiedBy = &ModifiedBy
	}
	
	if CreatedBy, ok := DomaincertificateauthorityMap["createdBy"].(string); ok {
		o.CreatedBy = &CreatedBy
	}
	
	if State, ok := DomaincertificateauthorityMap["state"].(string); ok {
		o.State = &State
	}
	
	if ModifiedByApp, ok := DomaincertificateauthorityMap["modifiedByApp"].(string); ok {
		o.ModifiedByApp = &ModifiedByApp
	}
	
	if CreatedByApp, ok := DomaincertificateauthorityMap["createdByApp"].(string); ok {
		o.CreatedByApp = &CreatedByApp
	}
	
	if Certificate, ok := DomaincertificateauthorityMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
	
	if VarType, ok := DomaincertificateauthorityMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Services, ok := DomaincertificateauthorityMap["services"].([]interface{}); ok {
		ServicesString, _ := json.Marshal(Services)
		json.Unmarshal(ServicesString, &o.Services)
	}
	
	if CertificateDetails, ok := DomaincertificateauthorityMap["certificateDetails"].([]interface{}); ok {
		CertificateDetailsString, _ := json.Marshal(CertificateDetails)
		json.Unmarshal(CertificateDetailsString, &o.CertificateDetails)
	}
	
	if SelfUri, ok := DomaincertificateauthorityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domaincertificateauthority) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
