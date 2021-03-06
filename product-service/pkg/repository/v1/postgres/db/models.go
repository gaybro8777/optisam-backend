// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type JobStatus string

const (
	JobStatusPENDING   JobStatus = "PENDING"
	JobStatusCOMPLETED JobStatus = "COMPLETED"
	JobStatusFAILED    JobStatus = "FAILED"
	JobStatusRETRY     JobStatus = "RETRY"
	JobStatusRUNNING   JobStatus = "RUNNING"
)

func (e *JobStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobStatus(s)
	case string:
		*e = JobStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for JobStatus: %T", src)
	}
	return nil
}

type Job struct {
	JobID     int32           `json:"job_id"`
	Type      string          `json:"type"`
	Status    JobStatus       `json:"status"`
	Data      json.RawMessage `json:"data"`
	Comments  sql.NullString  `json:"comments"`
	StartTime sql.NullTime    `json:"start_time"`
	EndTime   sql.NullTime    `json:"end_time"`
	CreatedAt time.Time       `json:"created_at"`
}

type Product struct {
	Swidtag         string         `json:"swidtag"`
	ProductName     string         `json:"product_name"`
	ProductVersion  string         `json:"product_version"`
	ProductEdition  string         `json:"product_edition"`
	ProductCategory string         `json:"product_category"`
	ProductEditor   string         `json:"product_editor"`
	Scope           string         `json:"scope"`
	OptionOf        string         `json:"option_of"`
	Cost            int32          `json:"cost"`
	AggregationID   int32          `json:"aggregation_id"`
	AggregationName string         `json:"aggregation_name"`
	CreatedOn       time.Time      `json:"created_on"`
	CreatedBy       string         `json:"created_by"`
	UpdatedOn       time.Time      `json:"updated_on"`
	UpdatedBy       sql.NullString `json:"updated_by"`
}

type ProductsApplication struct {
	Swidtag       string `json:"swidtag"`
	ApplicationID string `json:"application_id"`
}

type ProductsEquipment struct {
	Swidtag     string        `json:"swidtag"`
	EquipmentID string        `json:"equipment_id"`
	NumOfUsers  sql.NullInt32 `json:"num_of_users"`
}
