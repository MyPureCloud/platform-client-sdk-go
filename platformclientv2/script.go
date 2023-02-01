package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Script
type Script struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// VersionId
	VersionId *string `json:"versionId,omitempty"`

	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`

	// VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	VersionDate *time.Time `json:"versionDate,omitempty"`

	// StartPageId
	StartPageId *string `json:"startPageId,omitempty"`

	// StartPageName
	StartPageName *string `json:"startPageName,omitempty"`

	// Features
	Features *interface{} `json:"features,omitempty"`

	// Variables
	Variables *interface{} `json:"variables,omitempty"`

	// CustomActions
	CustomActions *interface{} `json:"customActions,omitempty"`

	// Pages
	Pages *[]Page `json:"pages,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Script) SetField(field string, fieldValue interface{}) {
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

func (o Script) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate","PublishedDate","VersionDate", }
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
	type Alias Script
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	PublishedDate := new(string)
	if o.PublishedDate != nil {
		
		*PublishedDate = timeutil.Strftime(o.PublishedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishedDate = nil
	}
	
	VersionDate := new(string)
	if o.VersionDate != nil {
		
		*VersionDate = timeutil.Strftime(o.VersionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VersionDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		VersionDate *string `json:"versionDate,omitempty"`
		
		StartPageId *string `json:"startPageId,omitempty"`
		
		StartPageName *string `json:"startPageName,omitempty"`
		
		Features *interface{} `json:"features,omitempty"`
		
		Variables *interface{} `json:"variables,omitempty"`
		
		CustomActions *interface{} `json:"customActions,omitempty"`
		
		Pages *[]Page `json:"pages,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		VersionId: o.VersionId,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		PublishedDate: PublishedDate,
		
		VersionDate: VersionDate,
		
		StartPageId: o.StartPageId,
		
		StartPageName: o.StartPageName,
		
		Features: o.Features,
		
		Variables: o.Variables,
		
		CustomActions: o.CustomActions,
		
		Pages: o.Pages,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Script) UnmarshalJSON(b []byte) error {
	var ScriptMap map[string]interface{}
	err := json.Unmarshal(b, &ScriptMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScriptMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ScriptMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := ScriptMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if VersionId, ok := ScriptMap["versionId"].(string); ok {
		o.VersionId = &VersionId
	}
    
	if createdDateString, ok := ScriptMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := ScriptMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if publishedDateString, ok := ScriptMap["publishedDate"].(string); ok {
		PublishedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishedDateString)
		o.PublishedDate = &PublishedDate
	}
	
	if versionDateString, ok := ScriptMap["versionDate"].(string); ok {
		VersionDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", versionDateString)
		o.VersionDate = &VersionDate
	}
	
	if StartPageId, ok := ScriptMap["startPageId"].(string); ok {
		o.StartPageId = &StartPageId
	}
    
	if StartPageName, ok := ScriptMap["startPageName"].(string); ok {
		o.StartPageName = &StartPageName
	}
    
	if Features, ok := ScriptMap["features"].(map[string]interface{}); ok {
		FeaturesString, _ := json.Marshal(Features)
		json.Unmarshal(FeaturesString, &o.Features)
	}
	
	if Variables, ok := ScriptMap["variables"].(map[string]interface{}); ok {
		VariablesString, _ := json.Marshal(Variables)
		json.Unmarshal(VariablesString, &o.Variables)
	}
	
	if CustomActions, ok := ScriptMap["customActions"].(map[string]interface{}); ok {
		CustomActionsString, _ := json.Marshal(CustomActions)
		json.Unmarshal(CustomActionsString, &o.CustomActions)
	}
	
	if Pages, ok := ScriptMap["pages"].([]interface{}); ok {
		PagesString, _ := json.Marshal(Pages)
		json.Unmarshal(PagesString, &o.Pages)
	}
	
	if SelfUri, ok := ScriptMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Script) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
