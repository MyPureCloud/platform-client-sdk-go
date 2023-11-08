package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Deploymentwebaction
type Deploymentwebaction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - System-generated UUID for the action.
	Id *string `json:"id,omitempty"`

	// MediaType - Action media type used to deliver the action.
	MediaType *string `json:"mediaType,omitempty"`

	// CustomerId - ID string of the customer that the action was triggered for.
	CustomerId *string `json:"customerId,omitempty"`

	// CustomerIdType - Type of the customer ID that the action was triggered for.
	CustomerIdType *string `json:"customerIdType,omitempty"`

	// ActionMapId - ID of the action map that triggered the action.
	ActionMapId *string `json:"actionMapId,omitempty"`

	// ActionMapVersion - Version of the action map that triggered the action.
	ActionMapVersion *int `json:"actionMapVersion,omitempty"`

	// SessionId - ID of the session that the action was triggered for.
	SessionId *string `json:"sessionId,omitempty"`

	// WebMessagingOfferProperties - Web messaging offer specific properties.
	WebMessagingOfferProperties *Webmessagingofferproperties `json:"webMessagingOfferProperties,omitempty"`

	// ContentOfferProperties - Content offer specific properties.
	ContentOfferProperties *Contentoffer `json:"contentOfferProperties,omitempty"`

	// OpenActionProperties - Open action specific properties.
	OpenActionProperties *Openactionproperties `json:"openActionProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Deploymentwebaction) SetField(field string, fieldValue interface{}) {
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

func (o Deploymentwebaction) MarshalJSON() ([]byte, error) {
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
	type Alias Deploymentwebaction
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		CustomerId *string `json:"customerId,omitempty"`
		
		CustomerIdType *string `json:"customerIdType,omitempty"`
		
		ActionMapId *string `json:"actionMapId,omitempty"`
		
		ActionMapVersion *int `json:"actionMapVersion,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		WebMessagingOfferProperties *Webmessagingofferproperties `json:"webMessagingOfferProperties,omitempty"`
		
		ContentOfferProperties *Contentoffer `json:"contentOfferProperties,omitempty"`
		
		OpenActionProperties *Openactionproperties `json:"openActionProperties,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		MediaType: o.MediaType,
		
		CustomerId: o.CustomerId,
		
		CustomerIdType: o.CustomerIdType,
		
		ActionMapId: o.ActionMapId,
		
		ActionMapVersion: o.ActionMapVersion,
		
		SessionId: o.SessionId,
		
		WebMessagingOfferProperties: o.WebMessagingOfferProperties,
		
		ContentOfferProperties: o.ContentOfferProperties,
		
		OpenActionProperties: o.OpenActionProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Deploymentwebaction) UnmarshalJSON(b []byte) error {
	var DeploymentwebactionMap map[string]interface{}
	err := json.Unmarshal(b, &DeploymentwebactionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DeploymentwebactionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if MediaType, ok := DeploymentwebactionMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if CustomerId, ok := DeploymentwebactionMap["customerId"].(string); ok {
		o.CustomerId = &CustomerId
	}
    
	if CustomerIdType, ok := DeploymentwebactionMap["customerIdType"].(string); ok {
		o.CustomerIdType = &CustomerIdType
	}
    
	if ActionMapId, ok := DeploymentwebactionMap["actionMapId"].(string); ok {
		o.ActionMapId = &ActionMapId
	}
    
	if ActionMapVersion, ok := DeploymentwebactionMap["actionMapVersion"].(float64); ok {
		ActionMapVersionInt := int(ActionMapVersion)
		o.ActionMapVersion = &ActionMapVersionInt
	}
	
	if SessionId, ok := DeploymentwebactionMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if WebMessagingOfferProperties, ok := DeploymentwebactionMap["webMessagingOfferProperties"].(map[string]interface{}); ok {
		WebMessagingOfferPropertiesString, _ := json.Marshal(WebMessagingOfferProperties)
		json.Unmarshal(WebMessagingOfferPropertiesString, &o.WebMessagingOfferProperties)
	}
	
	if ContentOfferProperties, ok := DeploymentwebactionMap["contentOfferProperties"].(map[string]interface{}); ok {
		ContentOfferPropertiesString, _ := json.Marshal(ContentOfferProperties)
		json.Unmarshal(ContentOfferPropertiesString, &o.ContentOfferProperties)
	}
	
	if OpenActionProperties, ok := DeploymentwebactionMap["openActionProperties"].(map[string]interface{}); ok {
		OpenActionPropertiesString, _ := json.Marshal(OpenActionProperties)
		json.Unmarshal(OpenActionPropertiesString, &o.OpenActionProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Deploymentwebaction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
