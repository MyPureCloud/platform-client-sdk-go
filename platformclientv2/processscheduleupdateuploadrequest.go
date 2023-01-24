package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Processscheduleupdateuploadrequest
type Processscheduleupdateuploadrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UploadKey - The uploadKey provided by the request to get an upload URL
	UploadKey *string `json:"uploadKey,omitempty"`

	// TeamIds - The list of teams to which the users being modified belong. Only required if the requesting user has conditional permission to wfm:schedule:edit
	TeamIds *[]string `json:"teamIds,omitempty"`

	// ManagementUnitIdsForAddedTeamUsers - The set of muIds to which agents belong if agents are being newly added to the schedule, if the requesting user has conditional permission to wfm:schedule:edit
	ManagementUnitIdsForAddedTeamUsers *[]string `json:"managementUnitIdsForAddedTeamUsers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Processscheduleupdateuploadrequest) SetField(field string, fieldValue interface{}) {
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

func (o Processscheduleupdateuploadrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Processscheduleupdateuploadrequest
	
	return json.Marshal(&struct { 
		UploadKey *string `json:"uploadKey,omitempty"`
		
		TeamIds *[]string `json:"teamIds,omitempty"`
		
		ManagementUnitIdsForAddedTeamUsers *[]string `json:"managementUnitIdsForAddedTeamUsers,omitempty"`
		Alias
	}{ 
		UploadKey: o.UploadKey,
		
		TeamIds: o.TeamIds,
		
		ManagementUnitIdsForAddedTeamUsers: o.ManagementUnitIdsForAddedTeamUsers,
		Alias:    (Alias)(o),
	})
}

func (o *Processscheduleupdateuploadrequest) UnmarshalJSON(b []byte) error {
	var ProcessscheduleupdateuploadrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ProcessscheduleupdateuploadrequestMap)
	if err != nil {
		return err
	}
	
	if UploadKey, ok := ProcessscheduleupdateuploadrequestMap["uploadKey"].(string); ok {
		o.UploadKey = &UploadKey
	}
    
	if TeamIds, ok := ProcessscheduleupdateuploadrequestMap["teamIds"].([]interface{}); ok {
		TeamIdsString, _ := json.Marshal(TeamIds)
		json.Unmarshal(TeamIdsString, &o.TeamIds)
	}
	
	if ManagementUnitIdsForAddedTeamUsers, ok := ProcessscheduleupdateuploadrequestMap["managementUnitIdsForAddedTeamUsers"].([]interface{}); ok {
		ManagementUnitIdsForAddedTeamUsersString, _ := json.Marshal(ManagementUnitIdsForAddedTeamUsers)
		json.Unmarshal(ManagementUnitIdsForAddedTeamUsersString, &o.ManagementUnitIdsForAddedTeamUsers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Processscheduleupdateuploadrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
