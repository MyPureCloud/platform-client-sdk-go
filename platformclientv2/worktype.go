package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Worktype
type Worktype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Worktype.
	Name *string `json:"name,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// Description - The description of the Worktype.
	Description *string `json:"description,omitempty"`

	// DateCreated - The creation date of the Worktype. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The modified date of the Worktype. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// DefaultWorkbin - The default Workbin for Workitems created from the Worktype.
	DefaultWorkbin *Workbinreference `json:"defaultWorkbin,omitempty"`

	// DefaultStatus - The default status for Workitems created from the Worktype.
	DefaultStatus *Workitemstatusreference `json:"defaultStatus,omitempty"`

	// Statuses - The list of possible statuses for Workitems created from the Worktype.
	Statuses *[]Workitemstatus `json:"statuses,omitempty"`

	// DefaultDurationSeconds - The default duration in seconds for Workitems created from the Worktype.
	DefaultDurationSeconds *int `json:"defaultDurationSeconds,omitempty"`

	// DefaultExpirationSeconds - The default expiration time in seconds for Workitems created from the Worktype.
	DefaultExpirationSeconds *int `json:"defaultExpirationSeconds,omitempty"`

	// DefaultDueDurationSeconds - The default due duration in seconds for Workitems created from the Worktype.
	DefaultDueDurationSeconds *int `json:"defaultDueDurationSeconds,omitempty"`

	// DefaultPriority - The default priority for Workitems created from the Worktype. The valid range is between -25,000,000 and 25,000,000.
	DefaultPriority *int `json:"defaultPriority,omitempty"`

	// DefaultLanguage - The default language for Workitems created from the Worktype.
	DefaultLanguage *Languagereference `json:"defaultLanguage,omitempty"`

	// DefaultTtlSeconds - The default time to time to live in seconds for Workitems created from the Worktype.
	DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`

	// ModifiedBy - The id of the User who modified the Worktype.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// DefaultQueue - The default queue for Workitems created from the Worktype.
	DefaultQueue *Workitemqueuereference `json:"defaultQueue,omitempty"`

	// DefaultSkills - The default skills for Workitems created from the Worktype.
	DefaultSkills *[]Routingskillreference `json:"defaultSkills,omitempty"`

	// AssignmentEnabled - When set to true, Workitems will be sent to the queue of the Worktype as they are created. Default value is false.
	AssignmentEnabled *bool `json:"assignmentEnabled,omitempty"`

	// Schema - The schema defining the custom attributes for Workitems created from the Worktype.
	Schema *Workitemschema `json:"schema,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Worktype) SetField(field string, fieldValue interface{}) {
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

func (o Worktype) MarshalJSON() ([]byte, error) {
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
	type Alias Worktype
	
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
		
		Division *Division `json:"division,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		DefaultWorkbin *Workbinreference `json:"defaultWorkbin,omitempty"`
		
		DefaultStatus *Workitemstatusreference `json:"defaultStatus,omitempty"`
		
		Statuses *[]Workitemstatus `json:"statuses,omitempty"`
		
		DefaultDurationSeconds *int `json:"defaultDurationSeconds,omitempty"`
		
		DefaultExpirationSeconds *int `json:"defaultExpirationSeconds,omitempty"`
		
		DefaultDueDurationSeconds *int `json:"defaultDueDurationSeconds,omitempty"`
		
		DefaultPriority *int `json:"defaultPriority,omitempty"`
		
		DefaultLanguage *Languagereference `json:"defaultLanguage,omitempty"`
		
		DefaultTtlSeconds *int `json:"defaultTtlSeconds,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		DefaultQueue *Workitemqueuereference `json:"defaultQueue,omitempty"`
		
		DefaultSkills *[]Routingskillreference `json:"defaultSkills,omitempty"`
		
		AssignmentEnabled *bool `json:"assignmentEnabled,omitempty"`
		
		Schema *Workitemschema `json:"schema,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Division: o.Division,
		
		Description: o.Description,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		DefaultWorkbin: o.DefaultWorkbin,
		
		DefaultStatus: o.DefaultStatus,
		
		Statuses: o.Statuses,
		
		DefaultDurationSeconds: o.DefaultDurationSeconds,
		
		DefaultExpirationSeconds: o.DefaultExpirationSeconds,
		
		DefaultDueDurationSeconds: o.DefaultDueDurationSeconds,
		
		DefaultPriority: o.DefaultPriority,
		
		DefaultLanguage: o.DefaultLanguage,
		
		DefaultTtlSeconds: o.DefaultTtlSeconds,
		
		ModifiedBy: o.ModifiedBy,
		
		DefaultQueue: o.DefaultQueue,
		
		DefaultSkills: o.DefaultSkills,
		
		AssignmentEnabled: o.AssignmentEnabled,
		
		Schema: o.Schema,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Worktype) UnmarshalJSON(b []byte) error {
	var WorktypeMap map[string]interface{}
	err := json.Unmarshal(b, &WorktypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorktypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorktypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Division, ok := WorktypeMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if Description, ok := WorktypeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if dateCreatedString, ok := WorktypeMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := WorktypeMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if DefaultWorkbin, ok := WorktypeMap["defaultWorkbin"].(map[string]interface{}); ok {
		DefaultWorkbinString, _ := json.Marshal(DefaultWorkbin)
		json.Unmarshal(DefaultWorkbinString, &o.DefaultWorkbin)
	}
	
	if DefaultStatus, ok := WorktypeMap["defaultStatus"].(map[string]interface{}); ok {
		DefaultStatusString, _ := json.Marshal(DefaultStatus)
		json.Unmarshal(DefaultStatusString, &o.DefaultStatus)
	}
	
	if Statuses, ok := WorktypeMap["statuses"].([]interface{}); ok {
		StatusesString, _ := json.Marshal(Statuses)
		json.Unmarshal(StatusesString, &o.Statuses)
	}
	
	if DefaultDurationSeconds, ok := WorktypeMap["defaultDurationSeconds"].(float64); ok {
		DefaultDurationSecondsInt := int(DefaultDurationSeconds)
		o.DefaultDurationSeconds = &DefaultDurationSecondsInt
	}
	
	if DefaultExpirationSeconds, ok := WorktypeMap["defaultExpirationSeconds"].(float64); ok {
		DefaultExpirationSecondsInt := int(DefaultExpirationSeconds)
		o.DefaultExpirationSeconds = &DefaultExpirationSecondsInt
	}
	
	if DefaultDueDurationSeconds, ok := WorktypeMap["defaultDueDurationSeconds"].(float64); ok {
		DefaultDueDurationSecondsInt := int(DefaultDueDurationSeconds)
		o.DefaultDueDurationSeconds = &DefaultDueDurationSecondsInt
	}
	
	if DefaultPriority, ok := WorktypeMap["defaultPriority"].(float64); ok {
		DefaultPriorityInt := int(DefaultPriority)
		o.DefaultPriority = &DefaultPriorityInt
	}
	
	if DefaultLanguage, ok := WorktypeMap["defaultLanguage"].(map[string]interface{}); ok {
		DefaultLanguageString, _ := json.Marshal(DefaultLanguage)
		json.Unmarshal(DefaultLanguageString, &o.DefaultLanguage)
	}
	
	if DefaultTtlSeconds, ok := WorktypeMap["defaultTtlSeconds"].(float64); ok {
		DefaultTtlSecondsInt := int(DefaultTtlSeconds)
		o.DefaultTtlSeconds = &DefaultTtlSecondsInt
	}
	
	if ModifiedBy, ok := WorktypeMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if DefaultQueue, ok := WorktypeMap["defaultQueue"].(map[string]interface{}); ok {
		DefaultQueueString, _ := json.Marshal(DefaultQueue)
		json.Unmarshal(DefaultQueueString, &o.DefaultQueue)
	}
	
	if DefaultSkills, ok := WorktypeMap["defaultSkills"].([]interface{}); ok {
		DefaultSkillsString, _ := json.Marshal(DefaultSkills)
		json.Unmarshal(DefaultSkillsString, &o.DefaultSkills)
	}
	
	if AssignmentEnabled, ok := WorktypeMap["assignmentEnabled"].(bool); ok {
		o.AssignmentEnabled = &AssignmentEnabled
	}
    
	if Schema, ok := WorktypeMap["schema"].(map[string]interface{}); ok {
		SchemaString, _ := json.Marshal(Schema)
		json.Unmarshal(SchemaString, &o.Schema)
	}
	
	if SelfUri, ok := WorktypeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Worktype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
