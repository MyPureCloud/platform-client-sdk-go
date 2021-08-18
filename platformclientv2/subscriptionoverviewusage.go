package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Subscriptionoverviewusage
type Subscriptionoverviewusage struct { 
	// Name - Product charge name
	Name *string `json:"name,omitempty"`


	// PartNumber - Product part number
	PartNumber *string `json:"partNumber,omitempty"`


	// Grouping - UI grouping key
	Grouping *string `json:"grouping,omitempty"`


	// UnitOfMeasureType - UI unit of measure
	UnitOfMeasureType *string `json:"unitOfMeasureType,omitempty"`


	// UsageQuantity - Usage count for specified period
	UsageQuantity *string `json:"usageQuantity,omitempty"`


	// OveragePrice - Price for usage / overage charge
	OveragePrice *string `json:"overagePrice,omitempty"`


	// PrepayQuantity - Items prepaid for specified period
	PrepayQuantity *string `json:"prepayQuantity,omitempty"`


	// PrepayPrice - Price for prepay charge
	PrepayPrice *string `json:"prepayPrice,omitempty"`


	// UsageNotes - Notes about the usage/charge item
	UsageNotes *string `json:"usageNotes,omitempty"`


	// IsCancellable - Indicates whether the item is cancellable
	IsCancellable *bool `json:"isCancellable,omitempty"`


	// BundleQuantity - Quantity multiplier for this charge
	BundleQuantity *string `json:"bundleQuantity,omitempty"`


	// IsThirdParty - A charge from a third party entity
	IsThirdParty *bool `json:"isThirdParty,omitempty"`

}

func (u *Subscriptionoverviewusage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Subscriptionoverviewusage

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		PartNumber *string `json:"partNumber,omitempty"`
		
		Grouping *string `json:"grouping,omitempty"`
		
		UnitOfMeasureType *string `json:"unitOfMeasureType,omitempty"`
		
		UsageQuantity *string `json:"usageQuantity,omitempty"`
		
		OveragePrice *string `json:"overagePrice,omitempty"`
		
		PrepayQuantity *string `json:"prepayQuantity,omitempty"`
		
		PrepayPrice *string `json:"prepayPrice,omitempty"`
		
		UsageNotes *string `json:"usageNotes,omitempty"`
		
		IsCancellable *bool `json:"isCancellable,omitempty"`
		
		BundleQuantity *string `json:"bundleQuantity,omitempty"`
		
		IsThirdParty *bool `json:"isThirdParty,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		PartNumber: u.PartNumber,
		
		Grouping: u.Grouping,
		
		UnitOfMeasureType: u.UnitOfMeasureType,
		
		UsageQuantity: u.UsageQuantity,
		
		OveragePrice: u.OveragePrice,
		
		PrepayQuantity: u.PrepayQuantity,
		
		PrepayPrice: u.PrepayPrice,
		
		UsageNotes: u.UsageNotes,
		
		IsCancellable: u.IsCancellable,
		
		BundleQuantity: u.BundleQuantity,
		
		IsThirdParty: u.IsThirdParty,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Subscriptionoverviewusage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
