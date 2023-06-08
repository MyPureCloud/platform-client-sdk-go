package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingactivityentitydata
type Routingactivityentitydata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityDate - The time at which the activity was observed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActivityDate *time.Time `json:"activityDate,omitempty"`

	// OrganizationPresenceId - Organization presence identifier
	OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`

	// PresenceDate - Date of the latest presence change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PresenceDate *time.Time `json:"presenceDate,omitempty"`

	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`

	// QueueMembershipStatus - Queue membership status (e.g. active or inactive)
	QueueMembershipStatus *string `json:"queueMembershipStatus,omitempty"`

	// RoutingStatus - Agent routing status
	RoutingStatus *string `json:"routingStatus,omitempty"`

	// RoutingStatusDate - Date of the latest routing status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	RoutingStatusDate *time.Time `json:"routingStatusDate,omitempty"`

	// SystemPresence - System presence
	SystemPresence *string `json:"systemPresence,omitempty"`

	// TeamId - The team ID the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Routingactivityentitydata) SetField(field string, fieldValue interface{}) {
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

func (o Routingactivityentitydata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ActivityDate","PresenceDate","RoutingStatusDate", }
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
	type Alias Routingactivityentitydata
	
	ActivityDate := new(string)
	if o.ActivityDate != nil {
		
		*ActivityDate = timeutil.Strftime(o.ActivityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActivityDate = nil
	}
	
	PresenceDate := new(string)
	if o.PresenceDate != nil {
		
		*PresenceDate = timeutil.Strftime(o.PresenceDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PresenceDate = nil
	}
	
	RoutingStatusDate := new(string)
	if o.RoutingStatusDate != nil {
		
		*RoutingStatusDate = timeutil.Strftime(o.RoutingStatusDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RoutingStatusDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityDate *string `json:"activityDate,omitempty"`
		
		OrganizationPresenceId *string `json:"organizationPresenceId,omitempty"`
		
		PresenceDate *string `json:"presenceDate,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		QueueMembershipStatus *string `json:"queueMembershipStatus,omitempty"`
		
		RoutingStatus *string `json:"routingStatus,omitempty"`
		
		RoutingStatusDate *string `json:"routingStatusDate,omitempty"`
		
		SystemPresence *string `json:"systemPresence,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		Alias
	}{ 
		ActivityDate: ActivityDate,
		
		OrganizationPresenceId: o.OrganizationPresenceId,
		
		PresenceDate: PresenceDate,
		
		QueueId: o.QueueId,
		
		QueueMembershipStatus: o.QueueMembershipStatus,
		
		RoutingStatus: o.RoutingStatus,
		
		RoutingStatusDate: RoutingStatusDate,
		
		SystemPresence: o.SystemPresence,
		
		TeamId: o.TeamId,
		
		UserId: o.UserId,
		Alias:    (Alias)(o),
	})
}

func (o *Routingactivityentitydata) UnmarshalJSON(b []byte) error {
	var RoutingactivityentitydataMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingactivityentitydataMap)
	if err != nil {
		return err
	}
	
	if activityDateString, ok := RoutingactivityentitydataMap["activityDate"].(string); ok {
		ActivityDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", activityDateString)
		o.ActivityDate = &ActivityDate
	}
	
	if OrganizationPresenceId, ok := RoutingactivityentitydataMap["organizationPresenceId"].(string); ok {
		o.OrganizationPresenceId = &OrganizationPresenceId
	}
    
	if presenceDateString, ok := RoutingactivityentitydataMap["presenceDate"].(string); ok {
		PresenceDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", presenceDateString)
		o.PresenceDate = &PresenceDate
	}
	
	if QueueId, ok := RoutingactivityentitydataMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if QueueMembershipStatus, ok := RoutingactivityentitydataMap["queueMembershipStatus"].(string); ok {
		o.QueueMembershipStatus = &QueueMembershipStatus
	}
    
	if RoutingStatus, ok := RoutingactivityentitydataMap["routingStatus"].(string); ok {
		o.RoutingStatus = &RoutingStatus
	}
    
	if routingStatusDateString, ok := RoutingactivityentitydataMap["routingStatusDate"].(string); ok {
		RoutingStatusDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", routingStatusDateString)
		o.RoutingStatusDate = &RoutingStatusDate
	}
	
	if SystemPresence, ok := RoutingactivityentitydataMap["systemPresence"].(string); ok {
		o.SystemPresence = &SystemPresence
	}
    
	if TeamId, ok := RoutingactivityentitydataMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if UserId, ok := RoutingactivityentitydataMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Routingactivityentitydata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
