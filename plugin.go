// Copyright (c) 2016, Gerasimos Maropoulos
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//	  this list of conditions and the following disclaimer
//    in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse
//    or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER AND CONTRIBUTOR, GERASIMOS MAROPOULOS BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package iris

import (
	"fmt"

	"github.com/kataras/iris/utils"
)

type (
	// IPlugin just an empty base for plugins
	IPlugin interface {
	}

	// IPluginGetName implements the GetName() string method
	IPluginGetName interface {
		// GetName has to returns the name of the plugin, a name is unique
		// name has to be not dependent from other methods of the plugin,
		// because it is being called even before the Activate
		GetName() string
	}

	// IPluginGetDescription implements the GetDescription() string method
	IPluginGetDescription interface {
		// GetDescription has to returns the description of what the plugins is used for
		GetDescription() string
	}

	// IPluginGetDescription implements the Activate(IPluginContainer) error method
	IPluginActivate interface {
		// Activate called BEFORE the plugin being added to the plugins list,
		// if Activate returns none nil error then the plugin is not being added to the list
		// it is being called only one time
		//
		// PluginContainer parameter used to add other plugins if that's necessary by the plugin
		Activate(IPluginContainer) error
	}

	// IPluginPreHandle implements the PreHandle(IRoute) method
	IPluginPreHandle interface {
		// PreHandle it's being called every time BEFORE a Route is registed to the Router
		//
		//  parameter is the Route
		PreHandle(IRoute)
	}
	// IPluginPostHandle implements the PostHandle(IRoute) method
	IPluginPostHandle interface {
		// PostHandle it's being called every time AFTER a Route successfully registed to the Router
		//
		// parameter is the Route
		PostHandle(IRoute)
	}
	// IPluginPreListen implements the PreListen(*Iris) method
	IPluginPreListen interface {
		// PreListen it's being called only one time, BEFORE the Server is started (if .Listen called)
		// is used to do work at the time all other things are ready to go
		//  parameter is the station
		PreListen(*Iris)
	}
	// IPluginPostListen implements the PostListen(*Iris) method
	IPluginPostListen interface {
		// PostListen it's being called only one time, AFTER the Server is started (if .Listen called)
		// parameter is the station
		PostListen(*Iris)
	}
	// IPluginPreClose implements the PreClose(*Iris) method
	IPluginPreClose interface {
		// PreClose it's being called only one time, BEFORE the Iris .Close method
		// any plugin cleanup/clear memory happens here
		//
		// The plugin is deactivated after this state
		PreClose(*Iris)
	}

	// IPluginPreDownload It's for the future, not being used, I need to create
	// and return an ActivatedPlugin type which will have it's methods, and pass it on .Activate
	// but now we return the whole pluginContainer, which I can't determinate which plugin tries to
	// download something, so we will leave it here for the future.
	IPluginPreDownload interface {
		// PreDownload it's being called every time a plugin tries to download something
		//
		// first parameter is the plugin
		// second parameter is the download url
		// must return a boolean, if false then the plugin is not permmited to download this file
		PreDownload(plugin IPlugin, downloadURL string) // bool
	}

	// IPluginContainer is the interface which the PluginContainer should implements
	IPluginContainer interface {
		Add(plugin IPlugin) error
		Remove(pluginName string) error
		GetName(plugin IPlugin) string
		GetDescription(plugin IPlugin) string
		GetByName(pluginName string) IPlugin
		Printf(format string, a ...interface{})
		DoPreHandle(route IRoute)
		DoPostHandle(route IRoute)
		DoPreListen(station *Iris)
		DoPostListen(station *Iris)
		DoPreClose(station *Iris)
		DoPreDownload(pluginTryToDownload IPlugin, downloadURL string)
		GetAll() []IPlugin
		// GetDownloader is the only one module that is used and fire listeners at the same time in this file
		GetDownloader() IDownloadManager
	}
	// IDownloadManager is the interface which the DownloadManager should implements
	IDownloadManager interface {
		DirectoryExists(dir string) bool
		DownloadZip(zipURL string, targetDir string) (string, error)
		Unzip(archive string, target string) (string, error)
		Remove(filePath string) error
		// install is just the flow of: downloadZip -> unzip -> removeFile(zippedFile)
		// accepts 2 parameters
		//
		// first parameter is the remote url file zip
		// second parameter is the target directory
		// returns a string(installedDirectory) and an error
		//
		// (string) installedDirectory is the directory which the zip file had, this is the real installation path, you don't need to know what it's because these things maybe change to the future let's keep it to return the correct path.
		// the installedDirectory is not empty when the installation is succed, the targetDirectory is not already exists and no error happens
		// the installedDirectory is empty when the installation is already done by previous time or an error happens
		Install(remoteFileZip string, targetDirectory string) (string, error)
	}

	// DownloadManager is just a struch which exports the util's downloadZip, directoryExists, unzip methods, used by the plugins via the PluginContainer
	DownloadManager struct {
	}
)

var _ IDownloadManager = &DownloadManager{}
var _ IPluginContainer = &PluginContainer{}

// DirectoryExists returns true if a given local directory exists
func (d *DownloadManager) DirectoryExists(dir string) bool {
	return utils.DirectoryExists(dir)
}

// DownloadZip downlodas a zip to the given local path location
func (d *DownloadManager) DownloadZip(zipURL string, targetDir string) (string, error) {
	return utils.DownloadZip(zipURL, targetDir)
}

// Unzip unzips a zip to the given local path location
func (d *DownloadManager) Unzip(archive string, target string) (string, error) {
	return utils.Unzip(archive, target)
}

// Remove deletes/removes/rm a file
func (d *DownloadManager) Remove(filePath string) error {
	return utils.RemoveFile(filePath)
}

// Install is just the flow of the: DownloadZip->Unzip->Remove the zip
func (d *DownloadManager) Install(remoteFileZip string, targetDirectory string) (string, error) {
	return utils.Install(remoteFileZip, targetDirectory)
}

// PluginContainer is the base container of all Iris, registed plugins
type PluginContainer struct {
	activatedPlugins []IPlugin
	downloader       *DownloadManager
}

// Add activates the plugins and if succeed then adds it to the activated plugins list
func (p *PluginContainer) Add(plugin IPlugin) error {
	if p.activatedPlugins == nil {
		p.activatedPlugins = make([]IPlugin, 0)
	}

	// Check if it's a plugin first, has Activate GetName

	// Check if the plugin already exists
	pName := p.GetName(plugin)
	if pName != "" && p.GetByName(pName) != nil {
		return ErrPluginAlreadyExists.Format(pName, p.GetDescription(plugin))
	}
	// Activate the plugin, if no error then add it to the plugins
	if pluginObj, ok := plugin.(IPluginActivate); ok {
		err := pluginObj.Activate(p)
		if err != nil {
			return ErrPluginActivate.Format(pName, err.Error())
		}
	}

	// All ok, add it to the plugins list
	p.activatedPlugins = append(p.activatedPlugins, plugin)
	return nil
}

// RemovePlugin it removes a plugin by it's name, if pluginName is empty "" or no plugin found with this name, then nothing is removed and a specific error is returned.
// This doesn't calls the PreClose method
func (p *PluginContainer) Remove(pluginName string) error {
	if p.activatedPlugins == nil {
		return ErrPluginRemoveNoPlugins.Return()
	}

	if pluginName == "" {
		//return error: cannot delete an unamed plugin
		return ErrPluginRemoveEmptyName.Return()
	}

	indexToRemove := -1
	for i := range p.activatedPlugins {
		if p.GetName(p.activatedPlugins[i]) == pluginName { // Note: if GetName is not implemented then the name is "" which is != with the plugiName, we checked this before.
			indexToRemove = i
		}
	}
	if indexToRemove == -1 { //if index stills -1 then no plugin was found with this name, just return an error. it is not a critical error.
		return ErrPluginRemoveNotFound.Return()
	} else {
		p.activatedPlugins = append(p.activatedPlugins[:indexToRemove], p.activatedPlugins[indexToRemove+1:]...)
	}
	return nil
}

// GetName returns the name of a plugin, if no GetName() implemented it returns an empty string ""
func (p *PluginContainer) GetName(plugin IPlugin) string {
	if pluginObj, ok := plugin.(IPluginGetName); ok {
		return pluginObj.GetName()
	}
	return ""
}

// GetDescription returns the name of a plugin, if no GetDescription() implemented it returns an empty string ""
func (p *PluginContainer) GetDescription(plugin IPlugin) string {
	if pluginObj, ok := plugin.(IPluginGetDescription); ok {
		return pluginObj.GetDescription()
	}
	return ""
}

// GetByName returns a plugin instance by it's name
func (p *PluginContainer) GetByName(pluginName string) IPlugin {
	if p.activatedPlugins == nil {
		return nil
	}

	for i := range p.activatedPlugins {
		if pluginObj, ok := p.activatedPlugins[i].(IPluginGetName); ok {
			if pluginObj.GetName() == pluginName {
				return pluginObj
			}
		}
	}

	return nil
}

// GetAll returns all activated plugins
func (p *PluginContainer) GetAll() []IPlugin {
	return p.activatedPlugins
}

// GetDownloader returns the download manager
func (p *PluginContainer) GetDownloader() IDownloadManager {
	// create it if and only if it used somewhere
	if p.downloader == nil {
		p.downloader = &DownloadManager{}
	}
	return p.downloader
}

// Printf sends plain text to any registed logger (future), some plugins maybe want use this method
// maybe at the future I change it, instead of sync even-driven to async channels...
func (p *PluginContainer) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...) //for now just this.
}

// DoPreHandle raise all plugins which has the PreHandle method
func (p *PluginContainer) DoPreHandle(route IRoute) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPreHandle); ok {
			pluginObj.PreHandle(route)
		}
	}
}

// DoPostHandle raise all plugins which has the DoPostHandle method
func (p *PluginContainer) DoPostHandle(route IRoute) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPostHandle); ok {
			pluginObj.PostHandle(route)
		}
	}
}

// DoPreListen raise all plugins which has the DoPreListen method
func (p *PluginContainer) DoPreListen(station *Iris) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPreListen); ok {
			pluginObj.PreListen(station)
		}
	}
}

// DoPostListen raise all plugins which has the DoPostListen method
func (p *PluginContainer) DoPostListen(station *Iris) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPostListen); ok {
			pluginObj.PostListen(station)
		}
	}
}

// DoPreClose raise all plugins which has the DoPreClose method
func (p *PluginContainer) DoPreClose(station *Iris) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPreClose); ok {
			pluginObj.PreClose(station)
		}
	}
}

// DoPreDownload raise all plugins which has the DoPreDownload method
func (p *PluginContainer) DoPreDownload(pluginTryToDownload IPlugin, downloadURL string) {
	for i := range p.activatedPlugins {
		// check if this method exists on our plugin obj, these are optionaly and call it
		if pluginObj, ok := p.activatedPlugins[i].(IPluginPreDownload); ok {
			pluginObj.PreDownload(pluginTryToDownload, downloadURL)
		}
	}
}
