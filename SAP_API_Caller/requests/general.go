package requests

type General struct {
	PurchasingInfoRecord string `json:"PurchasingInfoRecord"`
	Supplier             string `json:"Supplier"`
	Material             string `json:"Material"`
	// MaterialGroup        *string `json:"MaterialGroup"`
	// PurgDocOrderQuantityUnit    *string `json:"PurgDocOrderQuantityUnit"`
	SupplierMaterialNumber      *string `json:"SupplierMaterialNumber"`
	SupplierRespSalesPersonName *string `json:"SupplierRespSalesPersonName"`
	SupplierPhoneNumber         *string `json:"SupplierPhoneNumber"`
	SupplierMaterialGroup       *string `json:"SupplierMaterialGroup"`
	IsRegularSupplier           *string `json:"IsRegularSupplier"`
	// AvailabilityStartDate			              *string `json:"AvailabilityStartDate"`
	// AvailabilityEndDate      *string `json:"AvailabilityEndDate"`
	Manufacturer *string `json:"Manufacturer"`
	// CreationDate             *string `json:"CreationDate"`
	PurchasingInfoRecordDesc *string `json:"PurchasingInfoRecordDesc"`
	// LastChangeDateTime       *string `json:"LastChangeDateTime"`
	// IsDeleted *string `json:"IsDeleted"`
}
