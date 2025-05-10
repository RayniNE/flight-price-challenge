package models

type ResponsePriceline struct {
	Data    DataPriceline     `json:"data"`
	Meta    MetaPriceline     `json:"meta"`
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

type DataPriceline struct {
	IsFireFly bool               `json:"isFireFly"`
	Listings  []ListingPriceline `json:"listings"`
}

type ListingPriceline struct {
	TotalPriceWithDecimal  PriceWithDecimalPriceline       `json:"totalPriceWithDecimal"`
	ID                     string                          `json:"id"`
	GroupID                *string                         `json:"groupId"`
	RefID                  *string                         `json:"refId"`
	IsSaleEligible         bool                            `json:"isSaleEligible"`
	SeatsAvailable         int                             `json:"seatsAvailable"`
	Slices                 []SlicePriceline                `json:"slices"`
	VoidWindowInfo         VoidWindowInfoPriceline         `json:"voidWindowInfo"`
	Airlines               []AirlinePriceline              `json:"airlines"`
	AllFareBrandAttributes [][]FareBrandAttributePriceline `json:"allFareBrandAttributes"`
	AllFareBrandNames      []string                        `json:"allFareBrandNames"`
	CandidateID            *string                         `json:"candidateId"`
	CandidateKey           *string                         `json:"candidateKey"`
	Merchandising          []string                        `json:"merchandising"`
	SaleSavings            *string                         `json:"saleSavings"`
}

type PriceWithDecimalPriceline struct {
	Price float64 `json:"price"`
}

type SlicePriceline struct {
	UniqueSliceID     *string                   `json:"uniqueSliceId"`
	SliceKey          string                    `json:"sliceKey"`
	SliceRefID        int                       `json:"sliceRefId"`
	IsOvernight       bool                      `json:"isOvernight"`
	Segments          []SegmentPriceline        `json:"segments"`
	DurationInMinutes string                    `json:"durationInMinutes"`
	DisplayText       SliceDisplayTextPriceline `json:"displayText"`
	Merchandising     []string                  `json:"merchandising"`
	ID                int                       `json:"id"`
	IsSelected        *bool                     `json:"isSelected"`
	Departing         *string                   `json:"departing"`
	Arrival           *string                   `json:"arrival"`
}

type SegmentPriceline struct {
	ID                      int                         `json:"id"`
	CabinClass              string                      `json:"cabinClass"`
	UniqueSegID             *string                     `json:"uniqueSegId"`
	DepartInfo              AirportInfoPriceline        `json:"departInfo"`
	ArrivalInfo             AirportInfoPriceline        `json:"arrivalInfo"`
	OperatingAirline        string                      `json:"operatingAirline"`
	MarketingAirline        string                      `json:"marketingAirline"`
	EquipmentName           string                      `json:"equipmentName"`
	Equipment               string                      `json:"equipment"`
	Duration                int                         `json:"duration"`
	StopQuantity            int                         `json:"stopQuantity"`
	FlightNumber            string                      `json:"flightNumber"`
	IsSubjectToGovtApproval bool                        `json:"isSubjectToGovtApproval"`
	IsOvernight             bool                        `json:"isOvernight"`
	Brand                   BrandPriceline              `json:"brand"`
	DisplayText             SegmentDisplayTextPriceline `json:"displayText"`
	SegmentNote             *string                     `json:"segmentNote"`
	BkgClass                *string                     `json:"bkgClass"`
	BrandID                 string                      `json:"brandId"`
}

type AirportInfoPriceline struct {
	Airport AirportPriceline `json:"airport"`
	Time    TimePriceline    `json:"time"`
}

type AirportPriceline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type TimePriceline struct {
	DateTime string `json:"dateTime"`
}

type BrandPriceline struct {
	BrandAttributes []FareBrandAttributePriceline `json:"brandAttributes"`
}

type FareBrandAttributePriceline struct {
	Inclusion   string  `json:"inclusion"`
	Type        string  `json:"type"`
	Description *string `json:"description,omitempty"`
}

type SegmentDisplayTextPriceline struct {
	OperatedByText       *string `json:"operatedByText"`
	FlightNumber         *string `json:"flightNumber"`
	DisplayCabinName     *string `json:"displayCabinName"`
	PlaneChangeText      *string `json:"planeChangeText"`
	DifferentAirportText *string `json:"differentAirportText"`
	LayoverText          *string `json:"layoverText"`
	EquipmentText        string  `json:"equipmentText"`
}

type SliceDisplayTextPriceline struct {
	OperatedByText *string `json:"operatedByText"`
}

type VoidWindowInfoPriceline struct {
	TimeStamp    string  `json:"timeStamp"`
	TzDesignator *string `json:"tzDesignator"`
	HoursLeft    string  `json:"hoursLeft"`
}

type AirlinePriceline struct {
	MarketingAirline *string `json:"marketingAirline"`
	Name             string  `json:"name"`
	Image            string  `json:"image"`
}

type MetaPriceline struct {
	CurrentPage  int `json:"currentPage"`
	Limit        int `json:"limit"`
	TotalRecords int `json:"totalRecords"`
	TotalPage    int `json:"totalPage"`
}
