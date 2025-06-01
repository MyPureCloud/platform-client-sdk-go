package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bot - A bot instance
type Bot struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - This is a string type that should denote a unique ID of the bot for calling purposes by the Genesys service (EG a UUID).
	Id *string `json:"id,omitempty"`

	// Name - This is the name that will be displayed to the user in Architect.
	Name *string `json:"name,omitempty"`

	// Description - An optional description of the bot.
	Description *string `json:"description,omitempty"`

	// Provider - The provider of the bot.
	Provider *string `json:"provider,omitempty"`

	// Versions - This bots version.
	Versions *[]Botversion `json:"versions,omitempty"`

	// BotCompositeTag - A system-generated string that contains metadata about this bot.
	BotCompositeTag *string `json:"botCompositeTag,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bot) SetField(field string, fieldValue interface{}) {
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

func (o Bot) MarshalJSON() ([]byte, error) {
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
	type Alias Bot
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Versions *[]Botversion `json:"versions,omitempty"`
		
		BotCompositeTag *string `json:"botCompositeTag,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Provider: o.Provider,
		
		Versions: o.Versions,
		
		BotCompositeTag: o.BotCompositeTag,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Bot) UnmarshalJSON(b []byte) error {
	var BotMap map[string]interface{}
	err := json.Unmarshal(b, &BotMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BotMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := BotMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := BotMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Provider, ok := BotMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if Versions, ok := BotMap["versions"].([]interface{}); ok {
		VersionsString, _ := json.Marshal(Versions)
		json.Unmarshal(VersionsString, &o.Versions)
	}
	
	if BotCompositeTag, ok := BotMap["botCompositeTag"].(string); ok {
		o.BotCompositeTag = &BotCompositeTag
	}
    
	if SelfUri, ok := BotMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
