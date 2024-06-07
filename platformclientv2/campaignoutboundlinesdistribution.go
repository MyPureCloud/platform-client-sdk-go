package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Campaignoutboundlinesdistribution - Lines distribution information or Campaign's Edge Group or Site
type Campaignoutboundlinesdistribution struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Campaign - The Campaign for which dialing group distribution information was requested
	Campaign *Addressableentityref `json:"campaign,omitempty"`

	// MaxOutboundLineCount - Maximum outbound calls that can be placed for Campaign's Edge Group or Site
	MaxOutboundLineCount *int `json:"maxOutboundLineCount,omitempty"`

	// MaxLineUtilization - Maximum ratio of dialer calls to Campaign's Edge Group or Site capacity
	MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`

	// AvailableOutboundLines - Number of available outbound lines in Campaign's Edge Group or Site
	AvailableOutboundLines *int `json:"availableOutboundLines,omitempty"`

	// ReservedLines - Number of reserved outbound lines in Campaign's Edge Group or Site
	ReservedLines *int `json:"reservedLines,omitempty"`

	// CampaignsWithReservedLines - Information about campaigns with reserving lines in Campaign's Edge Group or Site
	CampaignsWithReservedLines *[]Campaignoutboundlinesreservation `json:"campaignsWithReservedLines,omitempty"`

	// CampaignsWithDynamicallyAllocatedLines - Information about campaigns using dynamic lines allocation in Campaign's Edge Group or Site
	CampaignsWithDynamicallyAllocatedLines *[]Campaignoutboundlinesallocation `json:"campaignsWithDynamicallyAllocatedLines,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Campaignoutboundlinesdistribution) SetField(field string, fieldValue interface{}) {
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

func (o Campaignoutboundlinesdistribution) MarshalJSON() ([]byte, error) {
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
	type Alias Campaignoutboundlinesdistribution
	
	return json.Marshal(&struct { 
		Campaign *Addressableentityref `json:"campaign,omitempty"`
		
		MaxOutboundLineCount *int `json:"maxOutboundLineCount,omitempty"`
		
		MaxLineUtilization *float32 `json:"maxLineUtilization,omitempty"`
		
		AvailableOutboundLines *int `json:"availableOutboundLines,omitempty"`
		
		ReservedLines *int `json:"reservedLines,omitempty"`
		
		CampaignsWithReservedLines *[]Campaignoutboundlinesreservation `json:"campaignsWithReservedLines,omitempty"`
		
		CampaignsWithDynamicallyAllocatedLines *[]Campaignoutboundlinesallocation `json:"campaignsWithDynamicallyAllocatedLines,omitempty"`
		Alias
	}{ 
		Campaign: o.Campaign,
		
		MaxOutboundLineCount: o.MaxOutboundLineCount,
		
		MaxLineUtilization: o.MaxLineUtilization,
		
		AvailableOutboundLines: o.AvailableOutboundLines,
		
		ReservedLines: o.ReservedLines,
		
		CampaignsWithReservedLines: o.CampaignsWithReservedLines,
		
		CampaignsWithDynamicallyAllocatedLines: o.CampaignsWithDynamicallyAllocatedLines,
		Alias:    (Alias)(o),
	})
}

func (o *Campaignoutboundlinesdistribution) UnmarshalJSON(b []byte) error {
	var CampaignoutboundlinesdistributionMap map[string]interface{}
	err := json.Unmarshal(b, &CampaignoutboundlinesdistributionMap)
	if err != nil {
		return err
	}
	
	if Campaign, ok := CampaignoutboundlinesdistributionMap["campaign"].(map[string]interface{}); ok {
		CampaignString, _ := json.Marshal(Campaign)
		json.Unmarshal(CampaignString, &o.Campaign)
	}
	
	if MaxOutboundLineCount, ok := CampaignoutboundlinesdistributionMap["maxOutboundLineCount"].(float64); ok {
		MaxOutboundLineCountInt := int(MaxOutboundLineCount)
		o.MaxOutboundLineCount = &MaxOutboundLineCountInt
	}
	
	if MaxLineUtilization, ok := CampaignoutboundlinesdistributionMap["maxLineUtilization"].(float64); ok {
		MaxLineUtilizationFloat32 := float32(MaxLineUtilization)
		o.MaxLineUtilization = &MaxLineUtilizationFloat32
	}
    
	if AvailableOutboundLines, ok := CampaignoutboundlinesdistributionMap["availableOutboundLines"].(float64); ok {
		AvailableOutboundLinesInt := int(AvailableOutboundLines)
		o.AvailableOutboundLines = &AvailableOutboundLinesInt
	}
	
	if ReservedLines, ok := CampaignoutboundlinesdistributionMap["reservedLines"].(float64); ok {
		ReservedLinesInt := int(ReservedLines)
		o.ReservedLines = &ReservedLinesInt
	}
	
	if CampaignsWithReservedLines, ok := CampaignoutboundlinesdistributionMap["campaignsWithReservedLines"].([]interface{}); ok {
		CampaignsWithReservedLinesString, _ := json.Marshal(CampaignsWithReservedLines)
		json.Unmarshal(CampaignsWithReservedLinesString, &o.CampaignsWithReservedLines)
	}
	
	if CampaignsWithDynamicallyAllocatedLines, ok := CampaignoutboundlinesdistributionMap["campaignsWithDynamicallyAllocatedLines"].([]interface{}); ok {
		CampaignsWithDynamicallyAllocatedLinesString, _ := json.Marshal(CampaignsWithDynamicallyAllocatedLines)
		json.Unmarshal(CampaignsWithDynamicallyAllocatedLinesString, &o.CampaignsWithDynamicallyAllocatedLines)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Campaignoutboundlinesdistribution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
