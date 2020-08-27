package types

import (
	log "github.com/sirupsen/logrus"
)

// State defination
type State struct {
	// RootArgs
	RootArgs *RootArgs
	// Logger
	Logger *log.Logger
}

// RootArgs defination
type RootArgs struct {
	// Verbose
	Verbose bool
	// LogLevel
	LogLevel string
}

// SyncArgs defination
type SyncArgs struct {
	// TargetPath
	TargetPath   string
	DepsFilePath string
	Arch         string
	OS           string
}

// VerifyArgs defination
type VerifyArgs struct {
	// TargetPath
	TargetPath   string
	DepsFilePath string
}

// GetTerraform defination
type GetTerraform struct {
	OS      string
	Arch    string
	Version string
	Path    string
}

// GetProvider defination
type GetProvider struct {
	OS      string
	Arch    string
	Version string
	Name    string
	Path    string
}
