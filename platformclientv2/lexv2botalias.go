package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexv2botalias
type Lexv2botalias struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Region - The Lex V2 bot region
	Region *string `json:"region,omitempty"`

	// AliasId - The Lex V2 bot alias Id
	AliasId *string `json:"aliasId,omitempty"`

	// Bot - The Lex V2 bot this is an alias for
	Bot *Lexv2bot `json:"bot,omitempty"`

	// BotVersion - The version of the Lex V2 bot this alias points at
	BotVersion *string `json:"botVersion,omitempty"`

	// Status - The status of the Lex V2 bot alias
	Status *string `json:"status,omitempty"`

	// Language - The target language of the Lex V2 bot
	Language *string `json:"language,omitempty"`

	// Intents - An array of Intents associated with this bot alias
	Intents *[]Lexv2intent `json:"intents,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Lexv2botalias) SetField(field string, fieldValue interface{}) {
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

func (o Lexv2botalias) MarshalJSON() ([]byte, error) {
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
	type Alias Lexv2botalias
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		AliasId *string `json:"aliasId,omitempty"`
		
		Bot *Lexv2bot `json:"bot,omitempty"`
		
		BotVersion *string `json:"botVersion,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Intents *[]Lexv2intent `json:"intents,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Region: o.Region,
		
		AliasId: o.AliasId,
		
		Bot: o.Bot,
		
		BotVersion: o.BotVersion,
		
		Status: o.Status,
		
		Language: o.Language,
		
		Intents: o.Intents,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Lexv2botalias) UnmarshalJSON(b []byte) error {
	var Lexv2botaliasMap map[string]interface{}
	err := json.Unmarshal(b, &Lexv2botaliasMap)
	if err != nil {
		return err
	}
	
	if Id, ok := Lexv2botaliasMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := Lexv2botaliasMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Region, ok := Lexv2botaliasMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if AliasId, ok := Lexv2botaliasMap["aliasId"].(string); ok {
		o.AliasId = &AliasId
	}
    
	if Bot, ok := Lexv2botaliasMap["bot"].(map[string]interface{}); ok {
		BotString, _ := json.Marshal(Bot)
		json.Unmarshal(BotString, &o.Bot)
	}
	
	if BotVersion, ok := Lexv2botaliasMap["botVersion"].(string); ok {
		o.BotVersion = &BotVersion
	}
    
	if Status, ok := Lexv2botaliasMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Language, ok := Lexv2botaliasMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Intents, ok := Lexv2botaliasMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if SelfUri, ok := Lexv2botaliasMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Lexv2botalias) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
