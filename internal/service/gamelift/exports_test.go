// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gamelift

// Exports for use in tests only.
var (
	ResourceAlias = resourceAlias
	ResourceBuild = resourceBuild
	ResourceFleet = resourceFleet

	DiffPortSettings = diffPortSettings
	FindAliasByID    = findAliasByID
	FindBuildByID    = findBuildByID
	FindFleetByID    = findFleetByID
)
