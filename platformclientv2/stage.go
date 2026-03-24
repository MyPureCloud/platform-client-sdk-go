package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Stage
type Stage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Stage.
	Name *string `json:"name,omitempty"`

	// Description - The description of the Stage.
	Description *string `json:"description,omitempty"`

	// DateCreated - The Stage creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The Stage modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DateCompleted - The Stage completion date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

	// DateStarted - The Stage start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// ModifiedBy - The ID of the User who modified the Stage.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// Version - The version of the Stage.
	Version *int `json:"version,omitempty"`

	// Status - The Status of the Stage.
	Status *string `json:"status,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// VarCase - The parent case of the Stage.
	VarCase *Casereference `json:"case,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Stage) SetField(field string, fieldValue interface{}) {
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

func (o Stage) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateCompleted","DateStarted", }
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
	type Alias Stage
	
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
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VarCase *Casereference `json:"case,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DateCompleted: DateCompleted,
		
		DateStarted: DateStarted,
		
		ModifiedBy: o.ModifiedBy,
		
		Version: o.Version,
		
		Status: o.Status,
		
		SelfUri: o.SelfUri,
		
		VarCase: o.VarCase,
		Alias:    (Alias)(o),
	})
}

func (o *Stage) UnmarshalJSON(b []byte) error {
	var StageMap map[string]interface{}
	err := json.Unmarshal(b, &StageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := StageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := StageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := StageMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := StageMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := StageMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if dateCompletedString, ok := StageMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if dateStartedString, ok := StageMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if ModifiedBy, ok := StageMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if Version, ok := StageMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Status, ok := StageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SelfUri, ok := StageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if VarCase, ok := StageMap["case"].(map[string]interface{}); ok {
		VarCaseString, _ := json.Marshal(VarCase)
		json.Unmarshal(VarCaseString, &o.VarCase)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Stage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
