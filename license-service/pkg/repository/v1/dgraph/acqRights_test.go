// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package dgraph

import (
	"fmt"
	v1 "optisam-backend/license-service/pkg/repository/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compareAcquiredRightsAllNoOrder(t *testing.T, name string, exp, act []*v1.AcquiredRights) {
	if !assert.Lenf(t, act, len(exp), "expected number of elemnts are: %d", len(exp)) {
		return
	}

	for i := range exp {
		idx := acqRightsIndexForSKU(exp[i].SKU, act)
		if !assert.NotEqualf(t, -1, idx, "acqRights with SKU is not found", exp[i].SKU) {
			continue
		}
		compareAcquiredRights(t, fmt.Sprintf("%s[%d]", name, i), exp[i], act[idx])
	}
}

func acqRightsIndexForSKU(sku string, act []*v1.AcquiredRights) int {
	for i := range act {
		if sku == act[i].SKU {
			return i
		}
	}
	return -1
}

func compareAcquiredRights(t *testing.T, name string, exp *v1.AcquiredRights, act *v1.AcquiredRights) {
	if exp == nil && act == nil {
		return
	}
	if exp == nil {
		assert.Nil(t, act, "attribute is expected to be nil")
	}

	// if exp.ID != "" {
	// 	assert.Equalf(t, exp.ID, act.ID, "%s.ID are not same", name)
	// }

	assert.Equalf(t, exp.Entity, act.Entity, "%s.Entity are not same", name)
	assert.Equalf(t, exp.SKU, act.SKU, "%s.SKU are not same", name)
	assert.Equalf(t, exp.SwidTag, act.SwidTag, "%s.SwidTag are not same", name)
	assert.Equalf(t, exp.ProductName, act.ProductName, "%s.ProductName are not same", name)
	assert.Equalf(t, exp.Editor, act.Editor, "%s.Type are not same", name)
	assert.Equalf(t, exp.Metric, act.Metric, "%s.Metric are not same", name)
	assert.Equalf(t, exp.AcquiredLicensesNumber, act.AcquiredLicensesNumber, "%s.AcquiredLicensesNumber are not same", name)
	assert.Equalf(t, exp.LicensesUnderMaintenanceNumber, act.LicensesUnderMaintenanceNumber, "%s.LicensesUnderMaintenanceNumber are not same", name)
	assert.Equalf(t, exp.AvgLicenesUnitPrice, act.AvgLicenesUnitPrice, "%s.AvgLicenesUnitPrice are not same", name)
	assert.Equalf(t, exp.AvgMaintenanceUnitPrice, act.AvgMaintenanceUnitPrice, "%s.AvgMaintenanceUnitPrice are not same", name)
	assert.Equalf(t, exp.TotalPurchaseCost, act.TotalPurchaseCost, "%s.TotalPurchaseCost are not same", name)
	assert.Equalf(t, exp.TotalMaintenanceCost, act.TotalMaintenanceCost, "%s.TotalMaintenanceCost are not same", name)
	assert.Equalf(t, exp.TotalCost, act.TotalCost, "%s.TotalCost are not same", name)
}
