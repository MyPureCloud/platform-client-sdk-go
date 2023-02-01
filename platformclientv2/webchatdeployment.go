package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatdeployment
type Webchatdeployment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// AuthenticationRequired
	AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`

	// AuthenticationUrl - URL for third party service authenticating web chat clients. See https://github.com/MyPureCloud/authenticated-web-chat-server-examples
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`

	// Disabled
	Disabled *bool `json:"disabled,omitempty"`

	// WebChatConfig
	WebChatConfig *Webchatconfig `json:"webChatConfig,omitempty"`

	// AllowedDomains
	AllowedDomains *[]string `json:"allowedDomains,omitempty"`

	// Flow - The URI of the Inbound Chat Flow to run when new chats are initiated under this Deployment.
	Flow *Domainentityref `json:"flow,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Webchatdeployment) SetField(field string, fieldValue interface{}) {
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

func (o Webchatdeployment) MarshalJSON() ([]byte, error) {
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
	type Alias Webchatdeployment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		AuthenticationRequired *bool `json:"authenticationRequired,omitempty"`
		
		AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		WebChatConfig *Webchatconfig `json:"webChatConfig,omitempty"`
		
		AllowedDomains *[]string `json:"allowedDomains,omitempty"`
		
		Flow *Domainentityref `json:"flow,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		AuthenticationRequired: o.AuthenticationRequired,
		
		AuthenticationUrl: o.AuthenticationUrl,
		
		Disabled: o.Disabled,
		
		WebChatConfig: o.WebChatConfig,
		
		AllowedDomains: o.AllowedDomains,
		
		Flow: o.Flow,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Webchatdeployment) UnmarshalJSON(b []byte) error {
	var WebchatdeploymentMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatdeploymentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatdeploymentMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebchatdeploymentMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := WebchatdeploymentMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if AuthenticationRequired, ok := WebchatdeploymentMap["authenticationRequired"].(bool); ok {
		o.AuthenticationRequired = &AuthenticationRequired
	}
    
	if AuthenticationUrl, ok := WebchatdeploymentMap["authenticationUrl"].(string); ok {
		o.AuthenticationUrl = &AuthenticationUrl
	}
    
	if Disabled, ok := WebchatdeploymentMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
    
	if WebChatConfig, ok := WebchatdeploymentMap["webChatConfig"].(map[string]interface{}); ok {
		WebChatConfigString, _ := json.Marshal(WebChatConfig)
		json.Unmarshal(WebChatConfigString, &o.WebChatConfig)
	}
	
	if AllowedDomains, ok := WebchatdeploymentMap["allowedDomains"].([]interface{}); ok {
		AllowedDomainsString, _ := json.Marshal(AllowedDomains)
		json.Unmarshal(AllowedDomainsString, &o.AllowedDomains)
	}
	
	if Flow, ok := WebchatdeploymentMap["flow"].(map[string]interface{}); ok {
		FlowString, _ := json.Marshal(Flow)
		json.Unmarshal(FlowString, &o.Flow)
	}
	
	if SelfUri, ok := WebchatdeploymentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatdeployment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
