package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Numberrouting
type Numberrouting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// NumberId - Phone number Id that has a disaster recovery linking
	NumberId *string `json:"numberId,omitempty"`

	// OwnerOrganizationId - Owner organization of numberId
	OwnerOrganizationId *string `json:"ownerOrganizationId,omitempty"`

	// CarrierCode - Code that indicates which carrier manages the number ie. VERIZON
	CarrierCode *string `json:"carrierCode,omitempty"`

	// PendingOrganizationId - OrganizationId where the number will be routed to during a change routing event
	PendingOrganizationId *string `json:"pendingOrganizationId,omitempty"`

	// Region - The current region where the number is located
	Region *string `json:"region,omitempty"`

	// Status - The current status of the number routing
	Status *string `json:"status,omitempty"`

	// ActiveOrganizationId - The orgId where the number is currently routing to
	ActiveOrganizationId *string `json:"activeOrganizationId,omitempty"`

	// LinkedOrganizationIds - List of linked organizations ids
	LinkedOrganizationIds *[]string `json:"linkedOrganizationIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Numberrouting) SetField(field string, fieldValue interface{}) {
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

func (o Numberrouting) MarshalJSON() ([]byte, error) {
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
	type Alias Numberrouting
	
	return json.Marshal(&struct { 
		NumberId *string `json:"numberId,omitempty"`
		
		OwnerOrganizationId *string `json:"ownerOrganizationId,omitempty"`
		
		CarrierCode *string `json:"carrierCode,omitempty"`
		
		PendingOrganizationId *string `json:"pendingOrganizationId,omitempty"`
		
		Region *string `json:"region,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ActiveOrganizationId *string `json:"activeOrganizationId,omitempty"`
		
		LinkedOrganizationIds *[]string `json:"linkedOrganizationIds,omitempty"`
		Alias
	}{ 
		NumberId: o.NumberId,
		
		OwnerOrganizationId: o.OwnerOrganizationId,
		
		CarrierCode: o.CarrierCode,
		
		PendingOrganizationId: o.PendingOrganizationId,
		
		Region: o.Region,
		
		Status: o.Status,
		
		ActiveOrganizationId: o.ActiveOrganizationId,
		
		LinkedOrganizationIds: o.LinkedOrganizationIds,
		Alias:    (Alias)(o),
	})
}

func (o *Numberrouting) UnmarshalJSON(b []byte) error {
	var NumberroutingMap map[string]interface{}
	err := json.Unmarshal(b, &NumberroutingMap)
	if err != nil {
		return err
	}
	
	if NumberId, ok := NumberroutingMap["numberId"].(string); ok {
		o.NumberId = &NumberId
	}
    
	if OwnerOrganizationId, ok := NumberroutingMap["ownerOrganizationId"].(string); ok {
		o.OwnerOrganizationId = &OwnerOrganizationId
	}
    
	if CarrierCode, ok := NumberroutingMap["carrierCode"].(string); ok {
		o.CarrierCode = &CarrierCode
	}
    
	if PendingOrganizationId, ok := NumberroutingMap["pendingOrganizationId"].(string); ok {
		o.PendingOrganizationId = &PendingOrganizationId
	}
    
	if Region, ok := NumberroutingMap["region"].(string); ok {
		o.Region = &Region
	}
    
	if Status, ok := NumberroutingMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ActiveOrganizationId, ok := NumberroutingMap["activeOrganizationId"].(string); ok {
		o.ActiveOrganizationId = &ActiveOrganizationId
	}
    
	if LinkedOrganizationIds, ok := NumberroutingMap["linkedOrganizationIds"].([]interface{}); ok {
		LinkedOrganizationIdsString, _ := json.Marshal(LinkedOrganizationIds)
		json.Unmarshal(LinkedOrganizationIdsString, &o.LinkedOrganizationIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Numberrouting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
