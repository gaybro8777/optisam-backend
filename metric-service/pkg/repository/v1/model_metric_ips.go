// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 

package v1

// MetricIPS is a representation of IBM.pvu.standard
type MetricIPS struct {
	ID               string
	Name             string
	NumCoreAttrID    string
	CoreFactorAttrID string
	BaseEqTypeID     string
}

// MetricIPSComputed has all the information required to be computed
type MetricIPSComputed struct {
	Name           string
	BaseType       *EquipmentType
	CoreFactorAttr *Attribute
	NumCoresAttr   *Attribute
}

// EquipmentType for creating equipment type
type EquipmentType struct {
	ID         string
	Type       string
	SourceID   string
	SourceName string
	ParentID   string
	ParentType string
	Attributes []*Attribute
}

//MetricIPSConfig is a representation of IBM.pvu.standard metric configuration
type MetricIPSConfig struct {
	ID             string
	Name           string
	NumCoreAttr    string
	CoreFactorAttr string
	BaseEqType     string
}
