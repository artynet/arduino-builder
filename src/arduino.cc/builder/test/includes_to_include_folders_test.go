/*
 * This file is part of Arduino Builder.
 *
 * Arduino Builder is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2015 Arduino LLC (http://www.arduino.cc/)
 */

package test

import (
	"arduino.cc/builder"
	"arduino.cc/builder/constants"
	"arduino.cc/builder/types"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestIncludesToIncludeFolders(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("downloaded_libraries", "Bridge", "examples", "Bridge", "Bridge.ino"),
		FQBN:                    "arduino:avr:leonardo",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	require.Equal(t, 1, len(importedLibraries))
	require.Equal(t, "Bridge", importedLibraries[0].Name)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.False(t, libraryResolutionResults["Bridge.h"].IsLibraryFromPlatform)
}

func TestIncludesToIncludeFoldersSketchWithIfDef(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("sketch2", "SketchWithIfDef.ino"),
		FQBN:                    "arduino:avr:leonardo",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	require.Equal(t, 0, len(importedLibraries))

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
}

func TestIncludesToIncludeFoldersIRremoteLibrary(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("sketch9", "sketch.ino"),
		FQBN:                    "arduino:avr:leonardo",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	sort.Sort(ByLibraryName(importedLibraries))
	require.Equal(t, 2, len(importedLibraries))
	require.Equal(t, "Bridge", importedLibraries[0].Name)
	require.Equal(t, "IRremote", importedLibraries[1].Name)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.False(t, libraryResolutionResults["Bridge.h"].IsLibraryFromPlatform)
	require.False(t, libraryResolutionResults["IRremote.h"].IsLibraryFromPlatform)
}

func TestIncludesToIncludeFoldersANewLibrary(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("sketch10", "sketch.ino"),
		FQBN:                    "arduino:avr:leonardo",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	sort.Sort(ByLibraryName(importedLibraries))
	require.Equal(t, 2, len(importedLibraries))
	require.Equal(t, "ANewLibrary-master", importedLibraries[0].Name)
	require.Equal(t, "IRremote", importedLibraries[1].Name)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.False(t, libraryResolutionResults["anewlibrary.h"].IsLibraryFromPlatform)
	require.False(t, libraryResolutionResults["IRremote.h"].IsLibraryFromPlatform)
}

func TestIncludesToIncludeFoldersDuplicateLibs(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware", "user_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		SketchLocation:          filepath.Join("user_hardware", "my_avr_platform", "avr", "libraries", "SPI", "examples", "BarometricPressureSensor", "BarometricPressureSensor.ino"),
		FQBN:                    "my_avr_platform:avr:custom_yun",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	sort.Sort(ByLibraryName(importedLibraries))
	require.Equal(t, 1, len(importedLibraries))
	require.Equal(t, "SPI", importedLibraries[0].Name)
	require.Equal(t, Abs(t, filepath.Join("user_hardware", "my_avr_platform", "avr", "libraries", "SPI")), importedLibraries[0].SrcFolder)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.True(t, libraryResolutionResults["SPI.h"].IsLibraryFromPlatform)
}

func TestIncludesToIncludeFoldersDuplicateLibsWithConflictingLibsOutsideOfPlatform(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware", "user_hardware"},
		ToolsFolders:            []string{"downloaded_tools"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("user_hardware", "my_avr_platform", "avr", "libraries", "SPI", "examples", "BarometricPressureSensor", "BarometricPressureSensor.ino"),
		FQBN:                    "my_avr_platform:avr:custom_yun",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	sort.Sort(ByLibraryName(importedLibraries))
	require.Equal(t, 1, len(importedLibraries))
	require.Equal(t, "SPI", importedLibraries[0].Name)
	require.Equal(t, Abs(t, filepath.Join("libraries", "SPI")), importedLibraries[0].SrcFolder)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.False(t, libraryResolutionResults["SPI.h"].IsLibraryFromPlatform)
}

func TestIncludesToIncludeFoldersDuplicateLibs2(t *testing.T) {
	DownloadCoresAndToolsAndLibraries(t)

	context := make(map[string]interface{})
	ctx := &types.Context{
		HardwareFolders:         []string{filepath.Join("..", "hardware"), "hardware", "downloaded_hardware", "downloaded_board_manager_stuff"},
		ToolsFolders:            []string{"downloaded_tools", "downloaded_board_manager_stuff"},
		BuiltInLibrariesFolders: []string{"downloaded_libraries"},
		OtherLibrariesFolders:   []string{"libraries"},
		SketchLocation:          filepath.Join("sketch_usbhost", "sketch_usbhost.ino"),
		FQBN:                    "arduino:samd:arduino_zero_native",
		ArduinoAPIVersion:       "10600",
	}

	buildPath := SetupBuildPath(t, context)
	defer os.RemoveAll(buildPath)

	context[constants.CTX_VERBOSE] = true

	commands := []types.Command{
		&builder.SetupHumanLoggerIfMissing{},

		&builder.ContainerSetupHardwareToolsLibsSketchAndProps{},

		&builder.ContainerMergeCopySketchFiles{},

		&builder.ContainerFindIncludes{},
	}

	for _, command := range commands {
		err := command.Run(context, ctx)
		NoError(t, err)
	}

	importedLibraries := context[constants.CTX_IMPORTED_LIBRARIES].([]*types.Library)
	sort.Sort(ByLibraryName(importedLibraries))
	require.Equal(t, 1, len(importedLibraries))
	require.Equal(t, "USBHost", importedLibraries[0].Name)
	require.Equal(t, Abs(t, filepath.Join("libraries", "USBHost", "src")), importedLibraries[0].SrcFolder)

	libraryResolutionResults := context[constants.CTX_LIBRARY_RESOLUTION_RESULTS].(map[string]types.LibraryResolutionResult)
	require.NotNil(t, libraryResolutionResults)
	require.False(t, libraryResolutionResults["Usb.h"].IsLibraryFromPlatform)
}
