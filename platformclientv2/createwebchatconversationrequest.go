package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createwebchatconversationrequest
type Createwebchatconversationrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OrganizationId - The organization identifier.
	OrganizationId *string `json:"organizationId,omitempty"`

	// DeploymentId - The web chat Deployment ID which contains the appropriate settings for this chat conversation.
	DeploymentId *string `json:"deploymentId,omitempty"`

	// RoutingTarget - The routing information to use for the new chat conversation.
	RoutingTarget *Webchatroutingtarget `json:"routingTarget,omitempty"`

	// MemberInfo - The guest member info to use for the new chat conversation.
	MemberInfo *Guestmemberinfo `json:"memberInfo,omitempty"`

	// MemberAuthToken - If the guest member is an authenticated member (ie, not anonymous) his JWT is provided here. The token will have been previously generated with the \"POST /api/v2/signeddata\" resource.
	MemberAuthToken *string `json:"memberAuthToken,omitempty"`

	// JourneyContext - A subset of the Journey System's data relevant to this conversation/session request (for external linkage and internal usage/context).
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createwebchatconversationrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createwebchatconversationrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createwebchatconversationrequest
	
	return json.Marshal(&struct { 
		OrganizationId *string `json:"organizationId,omitempty"`
		
		DeploymentId *string `json:"deploymentId,omitempty"`
		
		RoutingTarget *Webchatroutingtarget `json:"routingTarget,omitempty"`
		
		MemberInfo *Guestmemberinfo `json:"memberInfo,omitempty"`
		
		MemberAuthToken *string `json:"memberAuthToken,omitempty"`
		
		JourneyContext *Journeycontext `json:"journeyContext,omitempty"`
		Alias
	}{ 
		OrganizationId: o.OrganizationId,
		
		DeploymentId: o.DeploymentId,
		
		RoutingTarget: o.RoutingTarget,
		
		MemberInfo: o.MemberInfo,
		
		MemberAuthToken: o.MemberAuthToken,
		
		JourneyContext: o.JourneyContext,
		Alias:    (Alias)(o),
	})
}

func (o *Createwebchatconversationrequest) UnmarshalJSON(b []byte) error {
	var CreatewebchatconversationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreatewebchatconversationrequestMap)
	if err != nil {
		return err
	}
	
	if OrganizationId, ok := CreatewebchatconversationrequestMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if DeploymentId, ok := CreatewebchatconversationrequestMap["deploymentId"].(string); ok {
		o.DeploymentId = &DeploymentId
	}
    
	if RoutingTarget, ok := CreatewebchatconversationrequestMap["routingTarget"].(map[string]interface{}); ok {
		RoutingTargetString, _ := json.Marshal(RoutingTarget)
		json.Unmarshal(RoutingTargetString, &o.RoutingTarget)
	}
	
	if MemberInfo, ok := CreatewebchatconversationrequestMap["memberInfo"].(map[string]interface{}); ok {
		MemberInfoString, _ := json.Marshal(MemberInfo)
		json.Unmarshal(MemberInfoString, &o.MemberInfo)
	}
	
	if MemberAuthToken, ok := CreatewebchatconversationrequestMap["memberAuthToken"].(string); ok {
		o.MemberAuthToken = &MemberAuthToken
	}
    
	if JourneyContext, ok := CreatewebchatconversationrequestMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createwebchatconversationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
