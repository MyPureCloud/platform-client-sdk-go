package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Station
type Station struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// Status
	Status *string `json:"status,omitempty"`

	// UserId - The Id of the user currently logged in and associated with the station.
	UserId *string `json:"userId,omitempty"`

	// WebRtcUserId - The Id of the user configured for the station if it is of type inin_webrtc_softphone. Empty if station type is not inin_webrtc_softphone.
	WebRtcUserId *string `json:"webRtcUserId,omitempty"`

	// PrimaryEdge
	PrimaryEdge *Domainentityref `json:"primaryEdge,omitempty"`

	// SecondaryEdge
	SecondaryEdge *Domainentityref `json:"secondaryEdge,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// LineAppearanceId
	LineAppearanceId *string `json:"lineAppearanceId,omitempty"`

	// WebRtcMediaDscp - The default or configured value of media dscp for the station. Empty if station type is not inin_webrtc_softphone.
	WebRtcMediaDscp *int `json:"webRtcMediaDscp,omitempty"`

	// WebRtcPersistentEnabled - The default or configured value of persistent connection setting for the station. Empty if station type is not inin_webrtc_softphone.
	WebRtcPersistentEnabled *bool `json:"webRtcPersistentEnabled,omitempty"`

	// WebRtcForceTurn - Whether the station is configured to require TURN for routing WebRTC calls. Empty if station type is not inin_webrtc_softphone.
	WebRtcForceTurn *bool `json:"webRtcForceTurn,omitempty"`

	// WebRtcCallAppearances - The number of call appearances on the station.
	WebRtcCallAppearances *int `json:"webRtcCallAppearances,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Station) SetField(field string, fieldValue interface{}) {
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

func (o Station) MarshalJSON() ([]byte, error) {
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
	type Alias Station
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		WebRtcUserId *string `json:"webRtcUserId,omitempty"`
		
		PrimaryEdge *Domainentityref `json:"primaryEdge,omitempty"`
		
		SecondaryEdge *Domainentityref `json:"secondaryEdge,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LineAppearanceId *string `json:"lineAppearanceId,omitempty"`
		
		WebRtcMediaDscp *int `json:"webRtcMediaDscp,omitempty"`
		
		WebRtcPersistentEnabled *bool `json:"webRtcPersistentEnabled,omitempty"`
		
		WebRtcForceTurn *bool `json:"webRtcForceTurn,omitempty"`
		
		WebRtcCallAppearances *int `json:"webRtcCallAppearances,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Status: o.Status,
		
		UserId: o.UserId,
		
		WebRtcUserId: o.WebRtcUserId,
		
		PrimaryEdge: o.PrimaryEdge,
		
		SecondaryEdge: o.SecondaryEdge,
		
		VarType: o.VarType,
		
		LineAppearanceId: o.LineAppearanceId,
		
		WebRtcMediaDscp: o.WebRtcMediaDscp,
		
		WebRtcPersistentEnabled: o.WebRtcPersistentEnabled,
		
		WebRtcForceTurn: o.WebRtcForceTurn,
		
		WebRtcCallAppearances: o.WebRtcCallAppearances,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Station) UnmarshalJSON(b []byte) error {
	var StationMap map[string]interface{}
	err := json.Unmarshal(b, &StationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := StationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := StationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := StationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Status, ok := StationMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if UserId, ok := StationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if WebRtcUserId, ok := StationMap["webRtcUserId"].(string); ok {
		o.WebRtcUserId = &WebRtcUserId
	}
    
	if PrimaryEdge, ok := StationMap["primaryEdge"].(map[string]interface{}); ok {
		PrimaryEdgeString, _ := json.Marshal(PrimaryEdge)
		json.Unmarshal(PrimaryEdgeString, &o.PrimaryEdge)
	}
	
	if SecondaryEdge, ok := StationMap["secondaryEdge"].(map[string]interface{}); ok {
		SecondaryEdgeString, _ := json.Marshal(SecondaryEdge)
		json.Unmarshal(SecondaryEdgeString, &o.SecondaryEdge)
	}
	
	if VarType, ok := StationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if LineAppearanceId, ok := StationMap["lineAppearanceId"].(string); ok {
		o.LineAppearanceId = &LineAppearanceId
	}
    
	if WebRtcMediaDscp, ok := StationMap["webRtcMediaDscp"].(float64); ok {
		WebRtcMediaDscpInt := int(WebRtcMediaDscp)
		o.WebRtcMediaDscp = &WebRtcMediaDscpInt
	}
	
	if WebRtcPersistentEnabled, ok := StationMap["webRtcPersistentEnabled"].(bool); ok {
		o.WebRtcPersistentEnabled = &WebRtcPersistentEnabled
	}
    
	if WebRtcForceTurn, ok := StationMap["webRtcForceTurn"].(bool); ok {
		o.WebRtcForceTurn = &WebRtcForceTurn
	}
    
	if WebRtcCallAppearances, ok := StationMap["webRtcCallAppearances"].(float64); ok {
		WebRtcCallAppearancesInt := int(WebRtcCallAppearances)
		o.WebRtcCallAppearances = &WebRtcCallAppearancesInt
	}
	
	if SelfUri, ok := StationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Station) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
