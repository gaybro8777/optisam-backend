// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package v1

import "optisam-backend/dps-service/pkg/repository/v1/postgres/db"

//go:generate mockgen -destination=mock/mock.go -package=mock optisam-backend/common/optisam/Dps/repository/v1 Dps
type Dps interface {
	db.Querier
}
