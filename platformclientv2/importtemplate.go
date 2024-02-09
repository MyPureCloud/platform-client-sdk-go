package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Importtemplate
type Importtemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the import template.
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`

	// ContactListTemplate - ContactListTemplate for this ImportTemplate.
	ContactListTemplate *Domainentityref `json:"contactListTemplate,omitempty"`

	// ContactListFilter - ContactListFilter for this ImportTemplate.
	ContactListFilter *Domainentityref `json:"contactListFilter,omitempty"`

	// UseSplittingCriteria - Whether or not to use splitting criteria. Default is false.
	UseSplittingCriteria *bool `json:"useSplittingCriteria,omitempty"`

	// SplittingInformation - How to split contact records, required if useSplittingCriteria is true.
	SplittingInformation *Splittinginformation `json:"splittingInformation,omitempty"`

	// ListNameFormat - The list name format for target ContactLists. When Custom is provided, customListNameFormatValue is required.
	ListNameFormat *string `json:"listNameFormat,omitempty"`

	// CustomListNameFormatValue - Custom value for the list name format, at least %N is required. Any character other than the specified tokens will be used as is. Available tokens: %N: ListNamePrefix; %P: Part number; %F: Filter name; %C: Column value; YYYY: year; MM: month; DD: day; hh: hour; mm: minute; ss: second.
	CustomListNameFormatValue *string `json:"customListNameFormatValue,omitempty"`

	// ImportStatus - The status of the import process.
	ImportStatus *Importstatus `json:"importStatus,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Importtemplate) SetField(field string, fieldValue interface{}) {
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

func (o Importtemplate) MarshalJSON() ([]byte, error) {
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
	type Alias Importtemplate
	
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
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ContactListTemplate *Domainentityref `json:"contactListTemplate,omitempty"`
		
		ContactListFilter *Domainentityref `json:"contactListFilter,omitempty"`
		
		UseSplittingCriteria *bool `json:"useSplittingCriteria,omitempty"`
		
		SplittingInformation *Splittinginformation `json:"splittingInformation,omitempty"`
		
		ListNameFormat *string `json:"listNameFormat,omitempty"`
		
		CustomListNameFormatValue *string `json:"customListNameFormatValue,omitempty"`
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ContactListTemplate: o.ContactListTemplate,
		
		ContactListFilter: o.ContactListFilter,
		
		UseSplittingCriteria: o.UseSplittingCriteria,
		
		SplittingInformation: o.SplittingInformation,
		
		ListNameFormat: o.ListNameFormat,
		
		CustomListNameFormatValue: o.CustomListNameFormatValue,
		
		ImportStatus: o.ImportStatus,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Importtemplate) UnmarshalJSON(b []byte) error {
	var ImporttemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ImporttemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ImporttemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ImporttemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := ImporttemplateMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ImporttemplateMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := ImporttemplateMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ContactListTemplate, ok := ImporttemplateMap["contactListTemplate"].(map[string]interface{}); ok {
		ContactListTemplateString, _ := json.Marshal(ContactListTemplate)
		json.Unmarshal(ContactListTemplateString, &o.ContactListTemplate)
	}
	
	if ContactListFilter, ok := ImporttemplateMap["contactListFilter"].(map[string]interface{}); ok {
		ContactListFilterString, _ := json.Marshal(ContactListFilter)
		json.Unmarshal(ContactListFilterString, &o.ContactListFilter)
	}
	
	if UseSplittingCriteria, ok := ImporttemplateMap["useSplittingCriteria"].(bool); ok {
		o.UseSplittingCriteria = &UseSplittingCriteria
	}
    
	if SplittingInformation, ok := ImporttemplateMap["splittingInformation"].(map[string]interface{}); ok {
		SplittingInformationString, _ := json.Marshal(SplittingInformation)
		json.Unmarshal(SplittingInformationString, &o.SplittingInformation)
	}
	
	if ListNameFormat, ok := ImporttemplateMap["listNameFormat"].(string); ok {
		o.ListNameFormat = &ListNameFormat
	}
    
	if CustomListNameFormatValue, ok := ImporttemplateMap["customListNameFormatValue"].(string); ok {
		o.CustomListNameFormatValue = &CustomListNameFormatValue
	}
    
	if ImportStatus, ok := ImporttemplateMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if SelfUri, ok := ImporttemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Importtemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
