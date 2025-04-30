package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeployment - Details about a Web Deployment
type Webdeployment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The deployment ID
	Id *string `json:"id,omitempty"`

	// Name - The deployment name
	Name *string `json:"name,omitempty"`

	// Description - The description of the config
	Description *string `json:"description,omitempty"`

	// AllowAllDomains - Property indicates whether all domains are allowed or not. allowedDomains must be empty when this is set as true.
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`

	// AllowedDomains - The list of domains that are approved to use this deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`

	// SupportedContent - The supported content profile for a deployment
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`

	// Snippet - Javascript snippet used to load the config
	Snippet *string `json:"snippet,omitempty"`

	// DateCreated - The date the deployment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The date the deployment was most recently modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// LastModifiedUser - A reference to the user who most recently modified the deployment
	LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`

	// Flow - A reference to the inboundshortmessage flow used by this deployment
	Flow *Webdeploymentflowentityref `json:"flow,omitempty"`

	// Status - The current status of the deployment
	Status *string `json:"status,omitempty"`

	// PushIntegrations - The push integration objects associated with the deployment
	PushIntegrations *[]Pushintegration `json:"pushIntegrations,omitempty"`

	// Configuration - The config version this deployment uses
	Configuration *Webdeploymentconfigurationversionentityref `json:"configuration,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webdeployment) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Webdeployment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		
		AllowAllDomains *bool `json:"allowAllDomains,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		Snippet *string `json:"snippet,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		LastModifiedUser *Addressableentityref `json:"lastModifiedUser,omitempty"`
		
		Flow *Webdeploymentflowentityref `json:"flow,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		PushIntegrations *[]Pushintegration `json:"pushIntegrations,omitempty"`
		
		Configuration *Webdeploymentconfigurationversionentityref `json:"configuration,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		AllowAllDomains: o.AllowAllDomains,
		
		AllowedDomains: o.AllowedDomains,
		
		SupportedContent: o.SupportedContent,
		
		Snippet: o.Snippet,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		LastModifiedUser: o.LastModifiedUser,
		
		Flow: o.Flow,
		
		Status: o.Status,
		
		PushIntegrations: o.PushIntegrations,
		
		Configuration: o.Configuration,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
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
    
	if AllowAllDomains, ok := WebdeploymentMap["allowAllDomains"].(bool); ok {
		o.AllowAllDomains = &AllowAllDomains
	}
    
	if AllowedDomains, ok := WebdeploymentMap["allowedDomains"].([]interface{}); ok {
		AllowedDomainsString, _ := json.Marshal(AllowedDomains)
		json.Unmarshal(AllowedDomainsString, &o.AllowedDomains)
	}
	
	if SupportedContent, ok := WebdeploymentMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
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
    
	if PushIntegrations, ok := WebdeploymentMap["pushIntegrations"].([]interface{}); ok {
		PushIntegrationsString, _ := json.Marshal(PushIntegrations)
		json.Unmarshal(PushIntegrationsString, &o.PushIntegrations)
	}
	
	if Configuration, ok := WebdeploymentMap["configuration"].(map[string]interface{}); ok {
		ConfigurationString, _ := json.Marshal(Configuration)
		json.Unmarshal(ConfigurationString, &o.Configuration)
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
