// Code generated by "stringer -type=MMSmsCdmaServiceCategory"; DO NOT EDIT.

package go_modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmSmsCdmaServiceCategoryUnknown-0]
	_ = x[MmSmsCdmaServiceCategoryEmergencyBroadcast-1]
	_ = x[MmSmsCdmaServiceCategoryAdministrative-2]
	_ = x[MmSmsCdmaServiceCategoryMaintenance-3]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsLocal-4]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsRegional-5]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsNational-6]
	_ = x[MmSmsCdmaServiceCategoryGeneralNewsInternational-7]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsLocal-8]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsRegional-9]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsNational-10]
	_ = x[MmSmsCdmaServiceCategoryBusinessNewsInternational-11]
	_ = x[MmSmsCdmaServiceCategorySportsNewsLocal-12]
	_ = x[MmSmsCdmaServiceCategorySportsNewsRegional-13]
	_ = x[MmSmsCdmaServiceCategorySportsNewsNational-14]
	_ = x[MmSmsCdmaServiceCategorySportsNewsInternational-15]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsLocal-16]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsRegional-17]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsNational-18]
	_ = x[MmSmsCdmaServiceCategoryEntertainmentNewsInternational-19]
	_ = x[MmSmsCdmaServiceCategoryLocalWeather-20]
	_ = x[MmSmsCdmaServiceCategoryTrafficReport-21]
	_ = x[MmSmsCdmaServiceCategoryFlightSchedules-22]
	_ = x[MmSmsCdmaServiceCategoryRestaurants-23]
	_ = x[MmSmsCdmaServiceCategoryLodgings-24]
	_ = x[MmSmsCdmaServiceCategoryRetailDirectory-25]
	_ = x[MmSmsCdmaServiceCategoryAdvertisements-26]
	_ = x[MmSmsCdmaServiceCategoryStockQuotes-27]
	_ = x[MmSmsCdmaServiceCategoryEmployment-28]
	_ = x[MmSmsCdmaServiceCategoryHospitals-29]
	_ = x[MmSmsCdmaServiceCategoryTechnologyNews-30]
	_ = x[MmSmsCdmaServiceCategoryMulticategory-31]
	_ = x[MmSmsCdmaServiceCategoryCmasPresidentialAlert-4096]
	_ = x[MmSmsCdmaServiceCategoryCmasExtremeThreat-4097]
	_ = x[MmSmsCdmaServiceCategoryCmasSevereThreat-4098]
	_ = x[MmSmsCdmaServiceCategoryCmasChildAbductionEmergency-4099]
	_ = x[MmSmsCdmaServiceCategoryCmasTest-4100]
}

const (
	_MMSmsCdmaServiceCategory_name_0 = "MmSmsCdmaServiceCategoryUnknownMmSmsCdmaServiceCategoryEmergencyBroadcastMmSmsCdmaServiceCategoryAdministrativeMmSmsCdmaServiceCategoryMaintenanceMmSmsCdmaServiceCategoryGeneralNewsLocalMmSmsCdmaServiceCategoryGeneralNewsRegionalMmSmsCdmaServiceCategoryGeneralNewsNationalMmSmsCdmaServiceCategoryGeneralNewsInternationalMmSmsCdmaServiceCategoryBusinessNewsLocalMmSmsCdmaServiceCategoryBusinessNewsRegionalMmSmsCdmaServiceCategoryBusinessNewsNationalMmSmsCdmaServiceCategoryBusinessNewsInternationalMmSmsCdmaServiceCategorySportsNewsLocalMmSmsCdmaServiceCategorySportsNewsRegionalMmSmsCdmaServiceCategorySportsNewsNationalMmSmsCdmaServiceCategorySportsNewsInternationalMmSmsCdmaServiceCategoryEntertainmentNewsLocalMmSmsCdmaServiceCategoryEntertainmentNewsRegionalMmSmsCdmaServiceCategoryEntertainmentNewsNationalMmSmsCdmaServiceCategoryEntertainmentNewsInternationalMmSmsCdmaServiceCategoryLocalWeatherMmSmsCdmaServiceCategoryTrafficReportMmSmsCdmaServiceCategoryFlightSchedulesMmSmsCdmaServiceCategoryRestaurantsMmSmsCdmaServiceCategoryLodgingsMmSmsCdmaServiceCategoryRetailDirectoryMmSmsCdmaServiceCategoryAdvertisementsMmSmsCdmaServiceCategoryStockQuotesMmSmsCdmaServiceCategoryEmploymentMmSmsCdmaServiceCategoryHospitalsMmSmsCdmaServiceCategoryTechnologyNewsMmSmsCdmaServiceCategoryMulticategory"
	_MMSmsCdmaServiceCategory_name_1 = "MmSmsCdmaServiceCategoryCmasPresidentialAlertMmSmsCdmaServiceCategoryCmasExtremeThreatMmSmsCdmaServiceCategoryCmasSevereThreatMmSmsCdmaServiceCategoryCmasChildAbductionEmergencyMmSmsCdmaServiceCategoryCmasTest"
)

var (
	_MMSmsCdmaServiceCategory_index_0 = [...]uint16{0, 31, 73, 111, 146, 186, 229, 272, 320, 361, 405, 449, 498, 537, 579, 621, 668, 714, 763, 812, 866, 902, 939, 978, 1013, 1045, 1084, 1122, 1157, 1191, 1224, 1262, 1299}
	_MMSmsCdmaServiceCategory_index_1 = [...]uint8{0, 45, 86, 126, 177, 209}
)

func (i MMSmsCdmaServiceCategory) String() string {
	switch {
	case i <= 31:
		return _MMSmsCdmaServiceCategory_name_0[_MMSmsCdmaServiceCategory_index_0[i]:_MMSmsCdmaServiceCategory_index_0[i+1]]
	case 4096 <= i && i <= 4100:
		i -= 4096
		return _MMSmsCdmaServiceCategory_name_1[_MMSmsCdmaServiceCategory_index_1[i]:_MMSmsCdmaServiceCategory_index_1[i+1]]
	default:
		return "MMSmsCdmaServiceCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
