// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"encoding/json"
)

type Querier interface {
	DownloadReport(ctx context.Context, arg DownloadReportParams) (json.RawMessage, error)
	GetReport(ctx context.Context, arg GetReportParams) ([]GetReportRow, error)
	GetReportType(ctx context.Context, reportTypeID int32) (ReportType, error)
	GetReportTypes(ctx context.Context) ([]ReportType, error)
	InsertReportData(ctx context.Context, arg InsertReportDataParams) error
	SubmitReport(ctx context.Context, arg SubmitReportParams) (int32, error)
	UpdateReportStatus(ctx context.Context, arg UpdateReportStatusParams) error
}

var _ Querier = (*Queries)(nil)
