package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Developmentactivity - Development Activity object
type Developmentactivity struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DateCompleted - Date that activity was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

	// CreatedBy - User that created activity
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - Date activity was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// PercentageScore - The user's percentage score for this activity
	PercentageScore *float32 `json:"percentageScore,omitempty"`

	// IsPassed - True if the activity was passed
	IsPassed *bool `json:"isPassed,omitempty"`

	// IsLatest - True if this is the latest version of assignment assigned to the user
	IsLatest *bool `json:"isLatest,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// Name - The name of the activity
	Name *string `json:"name,omitempty"`

	// VarType - The type of activity
	VarType *string `json:"type,omitempty"`

	// Status - The status of the activity
	Status *string `json:"status,omitempty"`

	// DateDue - Due date for completion of the activity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDue *time.Time `json:"dateDue,omitempty"`

	// Facilitator - Facilitator of the activity
	Facilitator *Userreference `json:"facilitator,omitempty"`

	// Attendees - List of users attending the activity
	Attendees *[]Userreference `json:"attendees,omitempty"`

	// IsOverdue - Indicates if the activity is overdue
	IsOverdue *bool `json:"isOverdue,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Developmentactivity) SetField(field string, fieldValue interface{}) {
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

func (o Developmentactivity) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCompleted","DateCreated","DateDue", }
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
	type Alias Developmentactivity
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateDue := new(string)
	if o.DateDue != nil {
		
		*DateDue = timeutil.Strftime(o.DateDue, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDue = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		PercentageScore *float32 `json:"percentageScore,omitempty"`
		
		IsPassed *bool `json:"isPassed,omitempty"`
		
		IsLatest *bool `json:"isLatest,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateDue *string `json:"dateDue,omitempty"`
		
		Facilitator *Userreference `json:"facilitator,omitempty"`
		
		Attendees *[]Userreference `json:"attendees,omitempty"`
		
		IsOverdue *bool `json:"isOverdue,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DateCompleted: DateCompleted,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		PercentageScore: o.PercentageScore,
		
		IsPassed: o.IsPassed,
		
		IsLatest: o.IsLatest,
		
		SelfUri: o.SelfUri,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Status: o.Status,
		
		DateDue: DateDue,
		
		Facilitator: o.Facilitator,
		
		Attendees: o.Attendees,
		
		IsOverdue: o.IsOverdue,
		Alias:    (Alias)(o),
	})
}

func (o *Developmentactivity) UnmarshalJSON(b []byte) error {
	var DevelopmentactivityMap map[string]interface{}
	err := json.Unmarshal(b, &DevelopmentactivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DevelopmentactivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if dateCompletedString, ok := DevelopmentactivityMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if CreatedBy, ok := DevelopmentactivityMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := DevelopmentactivityMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if PercentageScore, ok := DevelopmentactivityMap["percentageScore"].(float64); ok {
		PercentageScoreFloat32 := float32(PercentageScore)
		o.PercentageScore = &PercentageScoreFloat32
	}
	
	if IsPassed, ok := DevelopmentactivityMap["isPassed"].(bool); ok {
		o.IsPassed = &IsPassed
	}
    
	if IsLatest, ok := DevelopmentactivityMap["isLatest"].(bool); ok {
		o.IsLatest = &IsLatest
	}
    
	if SelfUri, ok := DevelopmentactivityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if Name, ok := DevelopmentactivityMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := DevelopmentactivityMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Status, ok := DevelopmentactivityMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateDueString, ok := DevelopmentactivityMap["dateDue"].(string); ok {
		DateDue, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDueString)
		o.DateDue = &DateDue
	}
	
	if Facilitator, ok := DevelopmentactivityMap["facilitator"].(map[string]interface{}); ok {
		FacilitatorString, _ := json.Marshal(Facilitator)
		json.Unmarshal(FacilitatorString, &o.Facilitator)
	}
	
	if Attendees, ok := DevelopmentactivityMap["attendees"].([]interface{}); ok {
		AttendeesString, _ := json.Marshal(Attendees)
		json.Unmarshal(AttendeesString, &o.Attendees)
	}
	
	if IsOverdue, ok := DevelopmentactivityMap["isOverdue"].(bool); ok {
		o.IsOverdue = &IsOverdue
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Developmentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
