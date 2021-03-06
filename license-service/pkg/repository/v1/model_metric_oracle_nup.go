// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package v1

// MetricNUPOracle is a representation of oracle.nup.standard
type MetricNUPOracle struct {
	ID                    string
	Name                  string
	NumCoreAttrID         string
	NumCPUAttrID          string
	CoreFactorAttrID      string
	StartEqTypeID         string
	BaseEqTypeID          string
	AggerateLevelEqTypeID string
	EndEqTypeID           string
	NumberOfUsers         uint32
}

// MetricOPS return metric ops
func (m *MetricNUPOracle) MetricOPS() *MetricOPS {
	return &MetricOPS{
		ID:                    m.ID,
		Name:                  m.Name,
		NumCoreAttrID:         m.NumCoreAttrID,
		NumCPUAttrID:          m.NumCPUAttrID,
		CoreFactorAttrID:      m.CoreFactorAttrID,
		StartEqTypeID:         m.StartEqTypeID,
		BaseEqTypeID:          m.BaseEqTypeID,
		AggerateLevelEqTypeID: m.AggerateLevelEqTypeID,
		EndEqTypeID:           m.EndEqTypeID,
	}
}

// MetricNUPComputed has all the information required to be computed
type MetricNUPComputed struct {
	Name           string
	EqTypeTree     []*EquipmentType
	BaseType       *EquipmentType
	AggregateLevel *EquipmentType
	CoreFactorAttr *Attribute
	NumCoresAttr   *Attribute
	NumCPUAttr     *Attribute
	NumOfUsers     uint32
}

// NewMetricNUPComputed  returns NewMetricNUPComputed from MetricOPSComputed and num of users node
func NewMetricNUPComputed(m *MetricOPSComputed, numOfUsers uint32) *MetricNUPComputed {
	return &MetricNUPComputed{
		EqTypeTree:     m.EqTypeTree,
		BaseType:       m.BaseType,
		AggregateLevel: m.AggregateLevel,
		CoreFactorAttr: m.CoreFactorAttr,
		NumCoresAttr:   m.NumCoresAttr,
		NumCPUAttr:     m.NumCPUAttr,
		NumOfUsers:     numOfUsers,
	}
}

// MetricOPSComputed returns MetricOPSComputed for nup metic
func (m MetricNUPComputed) MetricOPSComputed() *MetricOPSComputed {
	return &MetricOPSComputed{
		EqTypeTree:     m.EqTypeTree,
		BaseType:       m.BaseType,
		AggregateLevel: m.AggregateLevel,
		CoreFactorAttr: m.CoreFactorAttr,
		NumCoresAttr:   m.NumCoresAttr,
		NumCPUAttr:     m.NumCPUAttr,
	}
}

type User struct {
	ID        string
	UserID    string
	UserCount int64
}
