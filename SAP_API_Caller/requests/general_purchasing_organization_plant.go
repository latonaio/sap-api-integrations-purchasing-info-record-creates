package requests

type GeneralPurchasingOrganizationPlant struct {
	General
	ToPurchasingOrganizationPlant `json:"to_PurgInfoRecdOrgPlantData"`
}

type ToPurchasingOrganizationPlant struct {
	Results []PurchasingOrganizationPlant `json:"results"`
}
