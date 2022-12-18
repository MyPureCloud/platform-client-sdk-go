package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingcampaignscheduleentitylisting
type Messagingcampaignscheduleentitylisting struct { 
	// Entities
	Entities *[]Messagingcampaignschedule `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total
	Total *int `json:"total,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

func (o *Messagingcampaignscheduleentitylisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingcampaignscheduleentitylisting
	
	return json.Marshal(&struct { 
		Entities *[]Messagingcampaignschedule `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		*Alias
	}{ 
		Entities: o.Entities,
		
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Total: o.Total,
		
		FirstUri: o.FirstUri,
		
		NextUri: o.NextUri,
		
		LastUri: o.LastUri,
		
		PreviousUri: o.PreviousUri,
		
		SelfUri: o.SelfUri,
		
		PageCount: o.PageCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingcampaignscheduleentitylisting) UnmarshalJSON(b []byte) error {
	var MessagingcampaignscheduleentitylistingMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingcampaignscheduleentitylistingMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := MessagingcampaignscheduleentitylistingMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if PageSize, ok := MessagingcampaignscheduleentitylistingMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := MessagingcampaignscheduleentitylistingMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Total, ok := MessagingcampaignscheduleentitylistingMap["total"].(float64); ok {
		TotalInt := int(Total)
		o.Total = &TotalInt
	}
	
	if FirstUri, ok := MessagingcampaignscheduleentitylistingMap["firstUri"].(string); ok {
		o.FirstUri = &FirstUri
	}
    
	if NextUri, ok := MessagingcampaignscheduleentitylistingMap["nextUri"].(string); ok {
		o.NextUri = &NextUri
	}
    
	if LastUri, ok := MessagingcampaignscheduleentitylistingMap["lastUri"].(string); ok {
		o.LastUri = &LastUri
	}
    
	if PreviousUri, ok := MessagingcampaignscheduleentitylistingMap["previousUri"].(string); ok {
		o.PreviousUri = &PreviousUri
	}
    
	if SelfUri, ok := MessagingcampaignscheduleentitylistingMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if PageCount, ok := MessagingcampaignscheduleentitylistingMap["pageCount"].(float64); ok {
		PageCountInt := int(PageCount)
		o.PageCount = &PageCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingcampaignscheduleentitylisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
