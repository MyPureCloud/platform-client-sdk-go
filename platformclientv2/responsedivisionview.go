package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Responsedivisionview - Division view of a response management response.
type Responsedivisionview struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ResponseType - The response type represented by the response.
	ResponseType *string `json:"responseType,omitempty"`

	// Libraries - One or more libraries response is associated with.
	Libraries *[]Librarydivisionview `json:"libraries,omitempty"`

	// Substitutions - Details about any text substitutions used in the texts for this response.
	Substitutions *[]Responsesubstitution `json:"substitutions,omitempty"`

	// SubstitutionsSchema - Metadata about the text substitutions in json schema format.
	SubstitutionsSchema *Jsonschemadocument `json:"substitutionsSchema,omitempty"`

	// MessagingTemplate - An optional messaging template definition for responseType.MessagingTemplate.
	MessagingTemplate *Messagingtemplate `json:"messagingTemplate,omitempty"`

	// Form - Form template definition for responseType.Form.
	Form *Form `json:"form,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Responsedivisionview) SetField(field string, fieldValue interface{}) {
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

func (o Responsedivisionview) MarshalJSON() ([]byte, error) {
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
	type Alias Responsedivisionview
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ResponseType *string `json:"responseType,omitempty"`
		
		Libraries *[]Librarydivisionview `json:"libraries,omitempty"`
		
		Substitutions *[]Responsesubstitution `json:"substitutions,omitempty"`
		
		SubstitutionsSchema *Jsonschemadocument `json:"substitutionsSchema,omitempty"`
		
		MessagingTemplate *Messagingtemplate `json:"messagingTemplate,omitempty"`
		
		Form *Form `json:"form,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ResponseType: o.ResponseType,
		
		Libraries: o.Libraries,
		
		Substitutions: o.Substitutions,
		
		SubstitutionsSchema: o.SubstitutionsSchema,
		
		MessagingTemplate: o.MessagingTemplate,
		
		Form: o.Form,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Responsedivisionview) UnmarshalJSON(b []byte) error {
	var ResponsedivisionviewMap map[string]interface{}
	err := json.Unmarshal(b, &ResponsedivisionviewMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ResponsedivisionviewMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ResponsedivisionviewMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ResponseType, ok := ResponsedivisionviewMap["responseType"].(string); ok {
		o.ResponseType = &ResponseType
	}
    
	if Libraries, ok := ResponsedivisionviewMap["libraries"].([]interface{}); ok {
		LibrariesString, _ := json.Marshal(Libraries)
		json.Unmarshal(LibrariesString, &o.Libraries)
	}
	
	if Substitutions, ok := ResponsedivisionviewMap["substitutions"].([]interface{}); ok {
		SubstitutionsString, _ := json.Marshal(Substitutions)
		json.Unmarshal(SubstitutionsString, &o.Substitutions)
	}
	
	if SubstitutionsSchema, ok := ResponsedivisionviewMap["substitutionsSchema"].(map[string]interface{}); ok {
		SubstitutionsSchemaString, _ := json.Marshal(SubstitutionsSchema)
		json.Unmarshal(SubstitutionsSchemaString, &o.SubstitutionsSchema)
	}
	
	if MessagingTemplate, ok := ResponsedivisionviewMap["messagingTemplate"].(map[string]interface{}); ok {
		MessagingTemplateString, _ := json.Marshal(MessagingTemplate)
		json.Unmarshal(MessagingTemplateString, &o.MessagingTemplate)
	}
	
	if Form, ok := ResponsedivisionviewMap["form"].(map[string]interface{}); ok {
		FormString, _ := json.Marshal(Form)
		json.Unmarshal(FormString, &o.Form)
	}
	
	if SelfUri, ok := ResponsedivisionviewMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responsedivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
