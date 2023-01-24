package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerdnclistconfigchangednclist
type Dialerdnclistconfigchangednclist struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ImportStatus
	ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`

	// Size - the number of phone numbers in the do not call list
	Size *int `json:"size,omitempty"`

	// DncSourceType - the type of dnc list being created, rds (csv file), gryphon, or dnc.com
	DncSourceType *string `json:"dncSourceType,omitempty"`

	// LoginId - the loginId if the dncSourceType is dnc.com
	LoginId *string `json:"loginId,omitempty"`

	// DncCodes - the list of dnc.com codes to be treated as DNC
	DncCodes *[]string `json:"dncCodes,omitempty"`

	// LicenseId - the license number if the dncSourceType is gryphon
	LicenseId *string `json:"licenseId,omitempty"`

	// ContactMethod
	ContactMethod *string `json:"contactMethod,omitempty"`

	// Division
	Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The UI-visible name of the object
	Name *string `json:"name,omitempty"`

	// DateCreated - Creation time of the entity
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - Last modified time of the entity
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialerdnclistconfigchangednclist) SetField(field string, fieldValue interface{}) {
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

func (o Dialerdnclistconfigchangednclist) MarshalJSON() ([]byte, error) {
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
	type Alias Dialerdnclistconfigchangednclist
	
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
		ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		DncSourceType *string `json:"dncSourceType,omitempty"`
		
		LoginId *string `json:"loginId,omitempty"`
		
		DncCodes *[]string `json:"dncCodes,omitempty"`
		
		LicenseId *string `json:"licenseId,omitempty"`
		
		ContactMethod *string `json:"contactMethod,omitempty"`
		
		Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		Alias
	}{ 
		ImportStatus: o.ImportStatus,
		
		Size: o.Size,
		
		DncSourceType: o.DncSourceType,
		
		LoginId: o.LoginId,
		
		DncCodes: o.DncCodes,
		
		LicenseId: o.LicenseId,
		
		ContactMethod: o.ContactMethod,
		
		Division: o.Division,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		Alias:    (Alias)(o),
	})
}

func (o *Dialerdnclistconfigchangednclist) UnmarshalJSON(b []byte) error {
	var DialerdnclistconfigchangednclistMap map[string]interface{}
	err := json.Unmarshal(b, &DialerdnclistconfigchangednclistMap)
	if err != nil {
		return err
	}
	
	if ImportStatus, ok := DialerdnclistconfigchangednclistMap["importStatus"].(map[string]interface{}); ok {
		ImportStatusString, _ := json.Marshal(ImportStatus)
		json.Unmarshal(ImportStatusString, &o.ImportStatus)
	}
	
	if Size, ok := DialerdnclistconfigchangednclistMap["size"].(float64); ok {
		SizeInt := int(Size)
		o.Size = &SizeInt
	}
	
	if DncSourceType, ok := DialerdnclistconfigchangednclistMap["dncSourceType"].(string); ok {
		o.DncSourceType = &DncSourceType
	}
    
	if LoginId, ok := DialerdnclistconfigchangednclistMap["loginId"].(string); ok {
		o.LoginId = &LoginId
	}
    
	if DncCodes, ok := DialerdnclistconfigchangednclistMap["dncCodes"].([]interface{}); ok {
		DncCodesString, _ := json.Marshal(DncCodes)
		json.Unmarshal(DncCodesString, &o.DncCodes)
	}
	
	if LicenseId, ok := DialerdnclistconfigchangednclistMap["licenseId"].(string); ok {
		o.LicenseId = &LicenseId
	}
    
	if ContactMethod, ok := DialerdnclistconfigchangednclistMap["contactMethod"].(string); ok {
		o.ContactMethod = &ContactMethod
	}
    
	if Division, ok := DialerdnclistconfigchangednclistMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if AdditionalProperties, ok := DialerdnclistconfigchangednclistMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Id, ok := DialerdnclistconfigchangednclistMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DialerdnclistconfigchangednclistMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := DialerdnclistconfigchangednclistMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := DialerdnclistconfigchangednclistMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := DialerdnclistconfigchangednclistMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangednclist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
