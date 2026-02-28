package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Knowledgesettingsresponse
type Knowledgesettingsresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - Knowledge Setting Id.
	Id *string `json:"id,omitempty"`

	// Name - Knowledge Setting Name.
	Name *string `json:"name,omitempty"`

	// Description - Knowledge setting description.
	Description *string `json:"description,omitempty"`

	// Sources - Knowledge source information searched upon.
	Sources *[]V3sourceref `json:"sources,omitempty"`

	// GenerationSetting - Settings for answer generation.
	GenerationSetting *Knowledgegenerationsetting `json:"generationSetting,omitempty"`

	// Stateful - Indicates if stateful search and generation is enabled for the knowledge setting.
	Stateful *bool `json:"stateful,omitempty"`

	// DateCreated - Knowledge setting created date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Knowledge setting last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ModifiedBy - The user who modified the knowledge setting.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Knowledgesettingsresponse) SetField(field string, fieldValue interface{}) {
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

func (o Knowledgesettingsresponse) MarshalJSON() ([]byte, error) {
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
	type Alias Knowledgesettingsresponse
	
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
		
		Sources *[]V3sourceref `json:"sources,omitempty"`
		
		GenerationSetting *Knowledgegenerationsetting `json:"generationSetting,omitempty"`
		
		Stateful *bool `json:"stateful,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Sources: o.Sources,
		
		GenerationSetting: o.GenerationSetting,
		
		Stateful: o.Stateful,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Knowledgesettingsresponse) UnmarshalJSON(b []byte) error {
	var KnowledgesettingsresponseMap map[string]interface{}
	err := json.Unmarshal(b, &KnowledgesettingsresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KnowledgesettingsresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KnowledgesettingsresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := KnowledgesettingsresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Sources, ok := KnowledgesettingsresponseMap["sources"].([]interface{}); ok {
		SourcesString, _ := json.Marshal(Sources)
		json.Unmarshal(SourcesString, &o.Sources)
	}
	
	if GenerationSetting, ok := KnowledgesettingsresponseMap["generationSetting"].(map[string]interface{}); ok {
		GenerationSettingString, _ := json.Marshal(GenerationSetting)
		json.Unmarshal(GenerationSettingString, &o.GenerationSetting)
	}
	
	if Stateful, ok := KnowledgesettingsresponseMap["stateful"].(bool); ok {
		o.Stateful = &Stateful
	}
    
	if dateCreatedString, ok := KnowledgesettingsresponseMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := KnowledgesettingsresponseMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := KnowledgesettingsresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if SelfUri, ok := KnowledgesettingsresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Knowledgesettingsresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
