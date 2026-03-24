package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Stepplan
type Stepplan struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the Stepplan.
	Name *string `json:"name,omitempty"`

	// Description - The description of the Stepplan.
	Description *string `json:"description,omitempty"`

	// Caseplan - The Caseplan of the Stepplan.
	Caseplan *Caseplanreference `json:"caseplan,omitempty"`

	// Stageplan - The Stageplan of the Stepplan.
	Stageplan *Stageplanreference `json:"stageplan,omitempty"`

	// DateCreated - The Stepplan creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified - The Stepplan modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`

	// ModifiedBy - The ID of the User who modified the Stepplan.
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// ActivityType - The activityType of the Stepplan.
	ActivityType *string `json:"activityType,omitempty"`

	// WorkitemSettings - The workitemSettings of the Stepplan.
	WorkitemSettings *Workitemsettingsresponse `json:"workitemSettings,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Stepplan) SetField(field string, fieldValue interface{}) {
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

func (o Stepplan) MarshalJSON() ([]byte, error) {
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
	type Alias Stepplan
	
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
		
		Caseplan *Caseplanreference `json:"caseplan,omitempty"`
		
		Stageplan *Stageplanreference `json:"stageplan,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		ActivityType *string `json:"activityType,omitempty"`
		
		WorkitemSettings *Workitemsettingsresponse `json:"workitemSettings,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Caseplan: o.Caseplan,
		
		Stageplan: o.Stageplan,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ModifiedBy: o.ModifiedBy,
		
		ActivityType: o.ActivityType,
		
		WorkitemSettings: o.WorkitemSettings,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Stepplan) UnmarshalJSON(b []byte) error {
	var StepplanMap map[string]interface{}
	err := json.Unmarshal(b, &StepplanMap)
	if err != nil {
		return err
	}
	
	if Id, ok := StepplanMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := StepplanMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := StepplanMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Caseplan, ok := StepplanMap["caseplan"].(map[string]interface{}); ok {
		CaseplanString, _ := json.Marshal(Caseplan)
		json.Unmarshal(CaseplanString, &o.Caseplan)
	}
	
	if Stageplan, ok := StepplanMap["stageplan"].(map[string]interface{}); ok {
		StageplanString, _ := json.Marshal(Stageplan)
		json.Unmarshal(StageplanString, &o.Stageplan)
	}
	
	if dateCreatedString, ok := StepplanMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := StepplanMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if ModifiedBy, ok := StepplanMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if ActivityType, ok := StepplanMap["activityType"].(string); ok {
		o.ActivityType = &ActivityType
	}
    
	if WorkitemSettings, ok := StepplanMap["workitemSettings"].(map[string]interface{}); ok {
		WorkitemSettingsString, _ := json.Marshal(WorkitemSettings)
		json.Unmarshal(WorkitemSettingsString, &o.WorkitemSettings)
	}
	
	if SelfUri, ok := StepplanMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Stepplan) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
