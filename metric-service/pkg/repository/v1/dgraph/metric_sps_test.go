// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package dgraph

import (
	"context"
	"errors"
	v1 "optisam-backend/metric-service/pkg/repository/v1"
	"testing"

	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/stretchr/testify/assert"
)

func TestMetricRepository_CreateMetricSPS(t *testing.T) {
	type args struct {
		ctx    context.Context
		scopes []string
	}
	tests := []struct {
		name    string
		l       *MetricRepository
		args    args
		setup   func() (*v1.MetricSPS, func() error, error)
		wantErr bool
	}{
		{name: "sucess",
			l: NewMetricRepository(dgClient),
			args: args{
				ctx: context.Background(),
			},
			setup: func() (retMat *v1.MetricSPS, cleanup func() error, retErr error) {

				baseID := "base"
				coreFactorAttrID := "coreFactor"
				numOfCoresAttrID := "cores"

				mu := &api.Mutation{
					CommitNow: true,
					Set: []*api.NQuad{

						&api.NQuad{
							Subject:     blankID(baseID),
							Predicate:   "type_name",
							ObjectValue: stringObjectValue("metadata"),
						},

						&api.NQuad{
							Subject:     blankID(coreFactorAttrID),
							Predicate:   "type_name",
							ObjectValue: stringObjectValue("metadata"),
						},
						&api.NQuad{
							Subject:     blankID(numOfCoresAttrID),
							Predicate:   "type_name",
							ObjectValue: stringObjectValue("metadata"),
						},
					},
				}
				assigned, err := dgClient.NewTxn().Mutate(context.Background(), mu)
				if err != nil {
					return nil, nil, err
				}

				baseID, ok := assigned.Uids[baseID]
				if !ok {
					return nil, nil, errors.New("baseID is not found in assigned map")
				}

				defer func() {
					if retErr != nil {
						if err := deleteNode(baseID); err != nil {
							t.Log(err)
						}
					}
				}()

				coreFactorAttrID, ok = assigned.Uids[coreFactorAttrID]
				if !ok {
					return nil, nil, errors.New("coreFactorAttrID is not found in assigned map")
				}

				defer func() {
					if retErr != nil {
						if err := deleteNode(coreFactorAttrID); err != nil {
							t.Log(err)
						}
					}
				}()

				numOfCoresAttrID, ok = assigned.Uids[numOfCoresAttrID]
				if !ok {
					return nil, nil, errors.New("numOfCoresAttrID is not found in assigned map")
				}

				defer func() {
					if retErr != nil {
						if err := deleteNode(numOfCoresAttrID); err != nil {
							t.Log(err)
						}
					}
				}()

				return &v1.MetricSPS{
						Name:             "sag.processor.standard",
						BaseEqTypeID:     baseID,
						CoreFactorAttrID: coreFactorAttrID,
						NumCoreAttrID:    numOfCoresAttrID,
					}, func() error {
						return deleteNodes(baseID, coreFactorAttrID, numOfCoresAttrID)
					}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mat, cleanup, err := tt.setup()
			if !assert.Empty(t, err, "not expecting error from setup") {
				return
			}
			defer func() {
				assert.Empty(t, cleanup(), "not expecting error in setup")
			}()
			gotRetMat, err := tt.l.CreateMetricSPS(tt.args.ctx, mat, tt.args.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricRepository.CreateMetricSPS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				defer func() {
					assert.Empty(t, deleteNode(gotRetMat.ID), "error not expected in deleting metric type")
				}()
				compareMetricSPS(t, "MetricOPS", mat, gotRetMat)
			}
		})
	}
}

func TestMetricRepository_GetMetricConfigSPS(t *testing.T) {
	type args struct {
		ctx     context.Context
		metName string
		scopes  []string
	}
	tests := []struct {
		name    string
		l       *MetricRepository
		args    args
		setup   func() (func() error, error)
		want    *v1.MetricSPSConfig
		wantErr bool
	}{
		{name: "SUCCESS",
			l: NewMetricRepository(dgClient),
			args: args{
				ctx:     context.Background(),
				metName: "sps1",
			},
			setup: func() (func() error, error) {
				ids, err := addMetricSPSConfig("sps1")
				if err != nil {
					t.Errorf("Failed to create config of SPS metric, err : %v", err)
				}
				return func() error {
					return deleteMetricConfig(ids)
				}, nil
			},
			want: &v1.MetricSPSConfig{
				Name:           "sps1",
				NumCoreAttr:    "sps_cores",
				CoreFactorAttr: "sps_corefactor",
				BaseEqType:     "server",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup, err := tt.setup()
			if !assert.Empty(t, err, "not expecting error from setup") {
				return
			}
			defer func() {
				assert.Empty(t, cleanup(), "not expecting error in setup")
			}()
			got, err := tt.l.GetMetricConfigSPS(tt.args.ctx, tt.args.metName, tt.args.scopes)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricRepository.GetMetricConfigSPS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				compareMetricSPSConfig(t, "MetricRepository.GetMetricConfigSPS", tt.want, got)
			}
		})
	}
}

func addMetricSPSConfig(metName string) (ids map[string]string, err error) {

	mu := &api.Mutation{
		CommitNow: true,
		Set: []*api.NQuad{
			&api.NQuad{
				Subject:     blankID("metric"),
				Predicate:   "metric.name",
				ObjectValue: stringObjectValue(metName),
			},
			&api.NQuad{
				Subject:     blankID("metric"),
				Predicate:   "dgraph.type",
				ObjectValue: stringObjectValue("Metric"),
			},
			&api.NQuad{
				Subject:   blankID("metric"),
				Predicate: "metric.sps.base",
				ObjectId:  "_:metadata1",
			},
			&api.NQuad{
				Subject:     blankID("metadata1"),
				Predicate:   "dgraph.type",
				ObjectValue: stringObjectValue("metadata"),
			},
			&api.NQuad{
				Subject:     blankID("metadata1"),
				Predicate:   "metadata.equipment.type",
				ObjectValue: stringObjectValue("server"),
			},
			&api.NQuad{
				Subject:   blankID("metric"),
				Predicate: "metric.sps.attr_core_factor",
				ObjectId:  "_:attribute1",
			},
			&api.NQuad{
				Subject:     blankID("attribute1"),
				Predicate:   "dgraph.type",
				ObjectValue: stringObjectValue("attr"),
			},
			&api.NQuad{
				Subject:     blankID("attribute1"),
				Predicate:   "attribute.name",
				ObjectValue: stringObjectValue("sps_corefactor"),
			},
			&api.NQuad{
				Subject:   blankID("metric"),
				Predicate: "metric.sps.attr_num_cores",
				ObjectId:  "_:attribute3",
			},
			&api.NQuad{
				Subject:     blankID("attribute3"),
				Predicate:   "dgraph.type",
				ObjectValue: stringObjectValue("attr"),
			},
			&api.NQuad{
				Subject:     blankID("attribute3"),
				Predicate:   "attribute.name",
				ObjectValue: stringObjectValue("sps_cores"),
			},
		},
	}
	assigned, err := dgClient.NewTxn().Mutate(context.Background(), mu)

	return assigned.Uids, err
}

func compareMetricSPS(t *testing.T, name string, exp, act *v1.MetricSPS) {
	if exp == nil && act == nil {
		return
	}
	if exp == nil {
		assert.Nil(t, act, "metadata is expected to be nil")
	}

	if exp.ID != "" {
		assert.Equalf(t, exp.ID, act.ID, "%s.ID should be same", name)
	}

	assert.Equalf(t, exp.Name, act.Name, "%s.Source should be same", name)
	assert.Equalf(t, exp.BaseEqTypeID, act.BaseEqTypeID, "%s.BaseEqTypeID should be same", name)
	assert.Equalf(t, exp.CoreFactorAttrID, act.CoreFactorAttrID, "%s.CoreFactorAttrID should be same", name)
	assert.Equalf(t, exp.NumCoreAttrID, act.NumCoreAttrID, "%s.NumCoreAttrID should be same", name)
}

func compareMetricSPSConfig(t *testing.T, name string, exp, act *v1.MetricSPSConfig) {
	if exp == nil && act == nil {
		return
	}
	if exp == nil {
		assert.Nil(t, act, "metadata is expected to be nil")
	}

	if exp.ID != "" {
		assert.Equalf(t, exp.ID, act.ID, "%s.ID should be same", name)
	}

	assert.Equalf(t, exp.Name, act.Name, "%s.Source should be same", name)
	assert.Equalf(t, exp.BaseEqType, act.BaseEqType, "%s.BaseEqType should be same", name)
	assert.Equalf(t, exp.CoreFactorAttr, act.CoreFactorAttr, "%s.CoreFactorAttr should be same", name)
	assert.Equalf(t, exp.NumCoreAttr, act.NumCoreAttr, "%s.NumCoreAttr should be same", name)
}
