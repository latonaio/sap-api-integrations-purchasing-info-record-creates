package sap_api_input_reader

import (
	"sap-api-integrations-purchasing-info-record-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneralPurchasingOrganizationPlant() *requests.GeneralPurchasingOrganizationPlant {
	data := sdc.PurchasingInfoRecord
	results := make([]requests.PurchasingOrganizationPlant, 0, len(data.PurchasingOrganizationPlant))

	general := sdc.ConvertToGeneral()

	for i := range data.PurchasingOrganizationPlant {
		results = append(results, *sdc.ConvertToPurchasingOrganizationPlant(i))
	}

	return &requests.GeneralPurchasingOrganizationPlant{
		General: *general,
		ToPurchasingOrganizationPlant: requests.ToPurchasingOrganizationPlant{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.PurchasingInfoRecord
	return &requests.General{
		PurchasingInfoRecord: data.PurchasingInfoRecord,
		Supplier:             data.Supplier,
		Material:             data.Material,
		// MaterialGroup:        data.MaterialGroup,
		// PurgDocOrderQuantityUnit:    data.PurgDocOrderQuantityUnit,
		SupplierMaterialNumber:      data.SupplierMaterialNumber,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplierMaterialGroup:       data.SupplierMaterialGroup,
		IsRegularSupplier:           data.IsRegularSupplier,
		// AvailabilityStartDate:       data.AvailabilityStartDate,
		// AvailabilityEndDate:      data.AvailabilityEndDate,
		Manufacturer: data.Manufacturer,
		// CreationDate:             data.CreationDate,
		PurchasingInfoRecordDesc: data.PurchasingInfoRecordDesc,
		// LastChangeDateTime:       data.LastChangeDateTime,
		// IsDeleted: data.IsDeleted,
	}
}

func (sdc *SDC) ConvertToPurchasingOrganizationPlant(num int) *requests.PurchasingOrganizationPlant {
	dataPurchasingInfoRecord := sdc.PurchasingInfoRecord
	data := sdc.PurchasingInfoRecord.PurchasingOrganizationPlant[num]
	return &requests.PurchasingOrganizationPlant{
		PurchasingInfoRecord:         dataPurchasingInfoRecord.PurchasingInfoRecord,
		PurchasingInfoRecordCategory: data.PurchasingInfoRecordCategory,
		PurchasingOrganization:       data.PurchasingOrganization,
		Plant:                        data.Plant,
		Supplier:                     data.Supplier,
		Material:                     data.Material,
		MaterialGroup:                data.MaterialGroup,
		PurgDocOrderQuantityUnit:     data.PurgDocOrderQuantityUnit,
		PurchasingGroup:              data.PurchasingGroup,
		// MinimumPurchaseOrderQuantity:   data.MinimumPurchaseOrderQuantity,
		StandardPurchaseOrderQuantity: data.StandardPurchaseOrderQuantity,
		MaterialPlannedDeliveryDurn:   data.MaterialPlannedDeliveryDurn,
		// OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		// UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		LastReferencingPurchaseOrder:   data.LastReferencingPurchaseOrder,
		LastReferencingPurOrderItem:    data.LastReferencingPurOrderItem,
		// NetPriceAmount:                 data.NetPriceAmount,
		// MaterialPriceUnitQty:          data.MaterialPriceUnitQty,
		PurchaseOrderPriceUnit: data.PurchaseOrderPriceUnit,
		// PriceValidityEndDate:          data.PriceValidityEndDate,
		InvoiceIsGoodsReceiptBased: data.InvoiceIsGoodsReceiptBased,
		TaxCode:                    data.TaxCode,
		IncotermsClassification:    data.IncotermsClassification,
		// MaximumOrderQuantity:          data.MaximumOrderQuantity,
		IsRelevantForAutomSrcg:        data.IsRelevantForAutomSrcg,
		IsEvaluatedRcptSettlmtAllowed: data.IsEvaluatedRcptSettlmtAllowed,
		IsPurOrderAllwdForInbDeliv:    data.IsPurOrderAllwdForInbDeliv,
		IsOrderAcknRqd:                data.IsOrderAcknRqd,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}
