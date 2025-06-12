package model

import "gorm.io/gorm"

// Cluster is the model of 3fs cluster.
type Cluster struct {
	gorm.Model
	Name           string
	NetworkType    string
	DiskType       string
	ReplicationNum int
	DiskNumPerNode int
}

// Node is the model of 3fs node.
type Node struct {
	gorm.Model
	Name string
	Host string
}

// FdbCluster is the model of FoundationDB cluster.
type FdbCluster struct {
	gorm.Model
}

// FdbProcess is the model of FoundationDB process.
type FdbProcess struct {
	gorm.Model
	Name   string
	NodeID uint
}

// ChService is the model of ClickHouse service.
type ChService struct {
	gorm.Model
	Name   string
	NodeID uint
}

// GrafanaService is the model of Grafana service.
type GrafanaService struct {
	gorm.Model
	Name   string
	NodeID uint
}

// PgService is the model of PostgreSQL service.
type PgService struct {
	gorm.Model
	Name   string
	NodeID uint
}

// MgmtService is the model of 3fs management service.
type MgmtService struct {
	gorm.Model
	Name     string
	NodeID   uint
	FsNodeID string
}

// MonService is the model of 3fs monitoring service.
type MonService struct {
	gorm.Model
	Name     string
	NodeID   uint
	FsNodeID string
}

// MetaService is the model of 3fs metadata service.
type MetaService struct {
	gorm.Model
	Name     string
	NodeID   uint
	FsNodeID string
}

// StorageService is the model of 3fs storage service.
type StorageService struct {
	gorm.Model
	Name     string
	NodeID   uint
	FsNodeID string
}

// Disk is the model of node disk.
type Disk struct {
	gorm.Model
	Name             string
	NodeID           uint
	FsNodeID         string
	StorageServiceID uint
	Index            int
	SizeByte         int64
	SerialNum        string
}

// Target is the model of 3fs target.
type Target struct {
	gorm.Model
	DiskID  uint
	NodeID  uint
	ChainID uint
}

// Chain is the model of 3fs chain.
type Chain struct {
	gorm.Model
	Name string
}
