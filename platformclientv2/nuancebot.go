package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nuancebot - Model for a Nuance bot
type Nuancebot struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Nuance bot Id
	Id *string `json:"id,omitempty"`

	// Name - Nuance bot name
	Name *string `json:"name,omitempty"`

	// IntegrationId - The Integration Id for this bot
	IntegrationId *string `json:"integrationId,omitempty"`

	// NuanceOrganization - The Nuance Organization for this bot
	NuanceOrganization *Nuanceorganization `json:"nuanceOrganization,omitempty"`

	// Application - The Application for this bot
	Application *Nuanceapplication `json:"application,omitempty"`

	// NuanceEnvironment - The environment of the Nuance bot
	NuanceEnvironment *Nuanceenvironment `json:"nuanceEnvironment,omitempty"`

	// Geography - The Geography of the Nuance bot
	Geography *Nuancegeography `json:"geography,omitempty"`

	// Credentials - client ID/Secret objects for the credentials that execute this Nuance bot
	Credentials *[]Nuancebotcredentials `json:"credentials,omitempty"`

	// Variables - List of available variables in this Nuance bot.  When querying, use the 'expand=variables' query param to populate this value
	Variables *[]Nuancebotvariable `json:"variables,omitempty"`

	// TransferNodes - List of transferNodes in this Nuance bot.  When querying, use the 'expand=transferNodes' query param to populate this value
	TransferNodes *[]Nuancebottransfernode `json:"transferNodes,omitempty"`

	// Locales - List of locales associated with this Nuance bot.  Generally in the ISO format such as 'en-US'
	Locales *[]string `json:"locales,omitempty"`

	// Channels - List of channels associated with this Nuance bot.
	Channels *[]Nuancechannel `json:"channels,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nuancebot) SetField(field string, fieldValue interface{}) {
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

func (o Nuancebot) MarshalJSON() ([]byte, error) {
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
	type Alias Nuancebot
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		NuanceOrganization *Nuanceorganization `json:"nuanceOrganization,omitempty"`
		
		Application *Nuanceapplication `json:"application,omitempty"`
		
		NuanceEnvironment *Nuanceenvironment `json:"nuanceEnvironment,omitempty"`
		
		Geography *Nuancegeography `json:"geography,omitempty"`
		
		Credentials *[]Nuancebotcredentials `json:"credentials,omitempty"`
		
		Variables *[]Nuancebotvariable `json:"variables,omitempty"`
		
		TransferNodes *[]Nuancebottransfernode `json:"transferNodes,omitempty"`
		
		Locales *[]string `json:"locales,omitempty"`
		
		Channels *[]Nuancechannel `json:"channels,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationId: o.IntegrationId,
		
		NuanceOrganization: o.NuanceOrganization,
		
		Application: o.Application,
		
		NuanceEnvironment: o.NuanceEnvironment,
		
		Geography: o.Geography,
		
		Credentials: o.Credentials,
		
		Variables: o.Variables,
		
		TransferNodes: o.TransferNodes,
		
		Locales: o.Locales,
		
		Channels: o.Channels,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Nuancebot) UnmarshalJSON(b []byte) error {
	var NuancebotMap map[string]interface{}
	err := json.Unmarshal(b, &NuancebotMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NuancebotMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := NuancebotMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IntegrationId, ok := NuancebotMap["integrationId"].(string); ok {
		o.IntegrationId = &IntegrationId
	}
    
	if NuanceOrganization, ok := NuancebotMap["nuanceOrganization"].(map[string]interface{}); ok {
		NuanceOrganizationString, _ := json.Marshal(NuanceOrganization)
		json.Unmarshal(NuanceOrganizationString, &o.NuanceOrganization)
	}
	
	if Application, ok := NuancebotMap["application"].(map[string]interface{}); ok {
		ApplicationString, _ := json.Marshal(Application)
		json.Unmarshal(ApplicationString, &o.Application)
	}
	
	if NuanceEnvironment, ok := NuancebotMap["nuanceEnvironment"].(map[string]interface{}); ok {
		NuanceEnvironmentString, _ := json.Marshal(NuanceEnvironment)
		json.Unmarshal(NuanceEnvironmentString, &o.NuanceEnvironment)
	}
	
	if Geography, ok := NuancebotMap["geography"].(map[string]interface{}); ok {
		GeographyString, _ := json.Marshal(Geography)
		json.Unmarshal(GeographyString, &o.Geography)
	}
	
	if Credentials, ok := NuancebotMap["credentials"].([]interface{}); ok {
		CredentialsString, _ := json.Marshal(Credentials)
		json.Unmarshal(CredentialsString, &o.Credentials)
	}
	
	if Variables, ok := NuancebotMap["variables"].([]interface{}); ok {
		VariablesString, _ := json.Marshal(Variables)
		json.Unmarshal(VariablesString, &o.Variables)
	}
	
	if TransferNodes, ok := NuancebotMap["transferNodes"].([]interface{}); ok {
		TransferNodesString, _ := json.Marshal(TransferNodes)
		json.Unmarshal(TransferNodesString, &o.TransferNodes)
	}
	
	if Locales, ok := NuancebotMap["locales"].([]interface{}); ok {
		LocalesString, _ := json.Marshal(Locales)
		json.Unmarshal(LocalesString, &o.Locales)
	}
	
	if Channels, ok := NuancebotMap["channels"].([]interface{}); ok {
		ChannelsString, _ := json.Marshal(Channels)
		json.Unmarshal(ChannelsString, &o.Channels)
	}
	
	if SelfUri, ok := NuancebotMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nuancebot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
