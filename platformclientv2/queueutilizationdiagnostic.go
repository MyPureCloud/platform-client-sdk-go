package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueutilizationdiagnostic
type Queueutilizationdiagnostic struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Queue - Identifier of the queue
	Queue *Domainentityref `json:"queue,omitempty"`

	// UsersInQueue - The number of users joined to the queue
	UsersInQueue *int `json:"usersInQueue,omitempty"`

	// ActiveUsersInQueue - The number of users active on the queue
	ActiveUsersInQueue *int `json:"activeUsersInQueue,omitempty"`

	// UsersOnQueue - The number of users with a status of on-queue
	UsersOnQueue *int `json:"usersOnQueue,omitempty"`

	// UsersNotUtilized - The number of users in the queue currently not engaged
	UsersNotUtilized *int `json:"usersNotUtilized,omitempty"`

	// UsersOnQueueWithStation - The number of users in the queue with a station
	UsersOnQueueWithStation *int `json:"usersOnQueueWithStation,omitempty"`

	// UsersOnACampaignCall - The number of users currently engaged in a campaign call
	UsersOnACampaignCall *int `json:"usersOnACampaignCall,omitempty"`

	// UsersOnDifferentEdgeGroup - The number of users whose station is homed to an edge different from the campaign
	UsersOnDifferentEdgeGroup *int `json:"usersOnDifferentEdgeGroup,omitempty"`

	// UsersOnANonCampaignCall - The number of users currently engaged in a communication that is not part of the campaign
	UsersOnANonCampaignCall *int `json:"usersOnANonCampaignCall,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Queueutilizationdiagnostic) SetField(field string, fieldValue interface{}) {
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

func (o Queueutilizationdiagnostic) MarshalJSON() ([]byte, error) {
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
	type Alias Queueutilizationdiagnostic
	
	return json.Marshal(&struct { 
		Queue *Domainentityref `json:"queue,omitempty"`
		
		UsersInQueue *int `json:"usersInQueue,omitempty"`
		
		ActiveUsersInQueue *int `json:"activeUsersInQueue,omitempty"`
		
		UsersOnQueue *int `json:"usersOnQueue,omitempty"`
		
		UsersNotUtilized *int `json:"usersNotUtilized,omitempty"`
		
		UsersOnQueueWithStation *int `json:"usersOnQueueWithStation,omitempty"`
		
		UsersOnACampaignCall *int `json:"usersOnACampaignCall,omitempty"`
		
		UsersOnDifferentEdgeGroup *int `json:"usersOnDifferentEdgeGroup,omitempty"`
		
		UsersOnANonCampaignCall *int `json:"usersOnANonCampaignCall,omitempty"`
		Alias
	}{ 
		Queue: o.Queue,
		
		UsersInQueue: o.UsersInQueue,
		
		ActiveUsersInQueue: o.ActiveUsersInQueue,
		
		UsersOnQueue: o.UsersOnQueue,
		
		UsersNotUtilized: o.UsersNotUtilized,
		
		UsersOnQueueWithStation: o.UsersOnQueueWithStation,
		
		UsersOnACampaignCall: o.UsersOnACampaignCall,
		
		UsersOnDifferentEdgeGroup: o.UsersOnDifferentEdgeGroup,
		
		UsersOnANonCampaignCall: o.UsersOnANonCampaignCall,
		Alias:    (Alias)(o),
	})
}

func (o *Queueutilizationdiagnostic) UnmarshalJSON(b []byte) error {
	var QueueutilizationdiagnosticMap map[string]interface{}
	err := json.Unmarshal(b, &QueueutilizationdiagnosticMap)
	if err != nil {
		return err
	}
	
	if Queue, ok := QueueutilizationdiagnosticMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if UsersInQueue, ok := QueueutilizationdiagnosticMap["usersInQueue"].(float64); ok {
		UsersInQueueInt := int(UsersInQueue)
		o.UsersInQueue = &UsersInQueueInt
	}
	
	if ActiveUsersInQueue, ok := QueueutilizationdiagnosticMap["activeUsersInQueue"].(float64); ok {
		ActiveUsersInQueueInt := int(ActiveUsersInQueue)
		o.ActiveUsersInQueue = &ActiveUsersInQueueInt
	}
	
	if UsersOnQueue, ok := QueueutilizationdiagnosticMap["usersOnQueue"].(float64); ok {
		UsersOnQueueInt := int(UsersOnQueue)
		o.UsersOnQueue = &UsersOnQueueInt
	}
	
	if UsersNotUtilized, ok := QueueutilizationdiagnosticMap["usersNotUtilized"].(float64); ok {
		UsersNotUtilizedInt := int(UsersNotUtilized)
		o.UsersNotUtilized = &UsersNotUtilizedInt
	}
	
	if UsersOnQueueWithStation, ok := QueueutilizationdiagnosticMap["usersOnQueueWithStation"].(float64); ok {
		UsersOnQueueWithStationInt := int(UsersOnQueueWithStation)
		o.UsersOnQueueWithStation = &UsersOnQueueWithStationInt
	}
	
	if UsersOnACampaignCall, ok := QueueutilizationdiagnosticMap["usersOnACampaignCall"].(float64); ok {
		UsersOnACampaignCallInt := int(UsersOnACampaignCall)
		o.UsersOnACampaignCall = &UsersOnACampaignCallInt
	}
	
	if UsersOnDifferentEdgeGroup, ok := QueueutilizationdiagnosticMap["usersOnDifferentEdgeGroup"].(float64); ok {
		UsersOnDifferentEdgeGroupInt := int(UsersOnDifferentEdgeGroup)
		o.UsersOnDifferentEdgeGroup = &UsersOnDifferentEdgeGroupInt
	}
	
	if UsersOnANonCampaignCall, ok := QueueutilizationdiagnosticMap["usersOnANonCampaignCall"].(float64); ok {
		UsersOnANonCampaignCallInt := int(UsersOnANonCampaignCall)
		o.UsersOnANonCampaignCall = &UsersOnANonCampaignCallInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueutilizationdiagnostic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
