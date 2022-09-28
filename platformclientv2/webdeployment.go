package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeployment - Details about a Web Deployment
type Webdeployment struct { 
	// Id - The deployment ID
	Id *string `json:"id,omitempty"`


	// Name - The deployment name
	Name *string `json:"name,omitempty"`


	// Description - The description of the config
	Description *string `json:"description,omitempty"`


	// Configuration - The config version this deployment uses
	Configuration *Webdeploymentconfigurationversion `json:"configuration,omitempty"`


	// AllowAllDomains - Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`


	// AllowedDomains - The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`


	// Snippet - Javascript snippet used to load the config
	Snippet *string `json:"snippet,omitempty"`


	// DateCreated - The date the deployment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date the deployment was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// LastModifiedUser - A reference to the user who most recently modified the deployment
	LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`


	// Flow - A reference to the inboundshortmessage flow used by this deployment
	Flow *Domainentityref `json:"flow,omitempty"`


	// Status - The current status of the deployment
	Status *string `json:"status,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Webdeployment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeployment
	
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
		
		Configuration *Webdeploymentconfigurationversion `json:"configuration,omitempty"`
		
		AllowAllDomains *bool `json:"allowAllDomains,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		Snippet *string `json:"snippet,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Configuration: o.Configuration,
		
		AllowAllDomains: o.AllowAllDomains,
		
		AllowedDomains: o.AllowedDomains,
		
		Snippet: o.Snippet,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		LastModifiedUser: o.LastModifiedUser,
		
		Flow: o.Flow,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeployment) UnmarshalJSON(b []byte) error {
	var WebdeploymentMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebdeploymentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := WebdeploymentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Configuration, ok := WebdeploymentMap["configuration"].(map[string]interface{}); ok {
		ConfigurationString, _ := json.Marshal(Configuration)
		json.Unmarshal(ConfigurationString, &o.Configuration)
	}
	
	if AllowAllDomains, ok := WebdeploymentMap["allowAllDomains"].(bool); ok {
		o.AllowAllDomains = &AllowAllDomains
	}
    
	if AllowedDomains, ok := WebdeploymentMap["allowedDomains"].([]interface{}); ok {
		AllowedDomainsString, _ := json.Marshal(AllowedDomains)
		json.Unmarshal(AllowedDomainsString, &o.AllowedDomains)
	}
	
	if Snippet, ok := WebdeploymentMap["snippet"].(string); ok {
		o.Snippet = &Snippet
	}
    
	if dateCreatedString, ok := WebdeploymentMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WebdeploymentMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if LastModifiedUser, ok := WebdeploymentMap["lastModifiedUser"].(map[string]interface{}); ok {
		LastModifiedUserString, _ := json.Marshal(LastModifiedUser)
		json.Unmarshal(LastModifiedUserString, &o.LastModifiedUser)
	}
	
	if Flow, ok := WebdeploymentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if Status, ok := WebdeploymentMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := WebdeploymentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
