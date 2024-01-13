package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemstatus
type Workitemstatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Category - The Category of the Status.
	Category *string `json:"category,omitempty"`

	// DestinationStatuses - The Statuses the Status can transition to.
	DestinationStatuses *[]Workitemstatusreference `json:"destinationStatuses,omitempty"`

	// Description - The description of the Status.
	Description *string `json:"description,omitempty"`

	// DefaultDestinationStatus - Default destination status to which this Status will transition to if auto status transition enabled.
	DefaultDestinationStatus *Workitemstatusreference `json:"defaultDestinationStatus,omitempty"`

	// StatusTransitionDelaySeconds - Delay in seconds for auto status transition
	StatusTransitionDelaySeconds *int `json:"statusTransitionDelaySeconds,omitempty"`

	// StatusTransitionTime - Time is represented as an ISO-8601 string without a timezone. For example: HH:mm:ss.SSS
	StatusTransitionTime *string `json:"statusTransitionTime,omitempty"`

	// Worktype - The Worktype containing the Status.
	Worktype *Worktypereference `json:"worktype,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemstatus) SetField(field string, fieldValue interface{}) {
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

func (o Workitemstatus) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemstatus
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		DestinationStatuses *[]Workitemstatusreference `json:"destinationStatuses,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		DefaultDestinationStatus *Workitemstatusreference `json:"defaultDestinationStatus,omitempty"`
		
		StatusTransitionDelaySeconds *int `json:"statusTransitionDelaySeconds,omitempty"`
		
		StatusTransitionTime *string `json:"statusTransitionTime,omitempty"`
		
		Worktype *Worktypereference `json:"worktype,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Category: o.Category,
		
		DestinationStatuses: o.DestinationStatuses,
		
		Description: o.Description,
		
		DefaultDestinationStatus: o.DefaultDestinationStatus,
		
		StatusTransitionDelaySeconds: o.StatusTransitionDelaySeconds,
		
		StatusTransitionTime: o.StatusTransitionTime,
		
		Worktype: o.Worktype,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemstatus) UnmarshalJSON(b []byte) error {
	var WorkitemstatusMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemstatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkitemstatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WorkitemstatusMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Category, ok := WorkitemstatusMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if DestinationStatuses, ok := WorkitemstatusMap["destinationStatuses"].([]interface{}); ok {
		DestinationStatusesString, _ := json.Marshal(DestinationStatuses)
		json.Unmarshal(DestinationStatusesString, &o.DestinationStatuses)
	}
	
	if Description, ok := WorkitemstatusMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if DefaultDestinationStatus, ok := WorkitemstatusMap["defaultDestinationStatus"].(map[string]interface{}); ok {
		DefaultDestinationStatusString, _ := json.Marshal(DefaultDestinationStatus)
		json.Unmarshal(DefaultDestinationStatusString, &o.DefaultDestinationStatus)
	}
	
	if StatusTransitionDelaySeconds, ok := WorkitemstatusMap["statusTransitionDelaySeconds"].(float64); ok {
		StatusTransitionDelaySecondsInt := int(StatusTransitionDelaySeconds)
		o.StatusTransitionDelaySeconds = &StatusTransitionDelaySecondsInt
	}
	
	if StatusTransitionTime, ok := WorkitemstatusMap["statusTransitionTime"].(string); ok {
		o.StatusTransitionTime = &StatusTransitionTime
	}
    
	if Worktype, ok := WorkitemstatusMap["worktype"].(map[string]interface{}); ok {
		WorktypeString, _ := json.Marshal(Worktype)
		json.Unmarshal(WorktypeString, &o.Worktype)
	}
	
	if SelfUri, ok := WorkitemstatusMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
