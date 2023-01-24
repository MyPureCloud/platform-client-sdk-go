package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetdeployment
type Widgetdeployment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description - A human-readable description of this Deployment.
	Description *string `json:"description,omitempty"`

	// AuthenticationRequired - When true, the customer members starting a chat must be authenticated by supplying their JWT to the create operation.
	AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`

	// Disabled - When true, all create chat operations using this Deployment will be rejected.
	Disabled *bool `json:"disabled,omitempty"`

	// Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
	Flow *Domainentityref `json:"flow,omitempty"`

	// AllowedDomains - The list of domains that are approved to use this Deployment; the list will be added to CORS headers for ease of web use.
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`

	// ClientType - The type of display widget for which this Deployment is configured, which controls the administrator settings shown.
	ClientType *string `json:"clientType,omitempty"`

	// ClientConfig - The client configuration options that should be made available to the clients of this Deployment.
	ClientConfig *Widgetclientconfig `json:"clientConfig,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Widgetdeployment) SetField(field string, fieldValue interface{}) {
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

func (o Widgetdeployment) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Widgetdeployment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		ClientType *string `json:"clientType,omitempty"`
		
		ClientConfig *Widgetclientconfig `json:"clientConfig,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		AuthenticationRequired: o.AuthenticationRequired,
		
		Disabled: o.Disabled,
		
		Flow: o.Flow,
		
		AllowedDomains: o.AllowedDomains,
		
		ClientType: o.ClientType,
		
		ClientConfig: o.ClientConfig,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Widgetdeployment) UnmarshalJSON(b []byte) error {
	var WidgetdeploymentMap map[string]interface{}
	err := json.Unmarshal(b, &WidgetdeploymentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WidgetdeploymentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WidgetdeploymentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := WidgetdeploymentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if AuthenticationRequired, ok := WidgetdeploymentMap["authenticationRequired"].(bool); ok {
		o.AuthenticationRequired = &AuthenticationRequired
	}
    
	if Disabled, ok := WidgetdeploymentMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if Flow, ok := WidgetdeploymentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if AllowedDomains, ok := WidgetdeploymentMap["allowedDomains"].([]interface{}); ok {
		AllowedDomainsString, _ := json.Marshal(AllowedDomains)
		json.Unmarshal(AllowedDomainsString, &o.AllowedDomains)
	}
	
	if ClientType, ok := WidgetdeploymentMap["clientType"].(string); ok {
		o.ClientType = &ClientType
	}
    
	if ClientConfig, ok := WidgetdeploymentMap["clientConfig"].(map[string]interface{}); ok {
		ClientConfigString, _ := json.Marshal(ClientConfig)
		json.Unmarshal(ClientConfigString, &o.ClientConfig)
	}
	
	if SelfUri, ok := WidgetdeploymentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Widgetdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
