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
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER AND CONTRIBUTOR, GERASIMOS MAROPOULOS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package typescript

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/cli/npm"
	"github.com/kataras/iris/cli/system"
	"github.com/kataras/iris/plugin/editor"
)

/* Notes

The editor is working when the typescript plugin finds a typescript project (tsconfig.json),
also working only if one typescript project found (normaly is one for client-side).

*/

// Name the name of the plugin, is "TypescriptPlugin"
const Name = "TypescriptPlugin"

var nodeModules = system.PathSeparator + "node_modules" + system.PathSeparator

type (
	// Options the struct which holds the TypescriptPlugin options
	// Has five (5) fields
	//
	// 1. Bin: 	string, the typescript installation directory/typescript/lib/tsc.js, if empty it will search inside global npm modules
	// 2. Dir:     string, Dir set the root, where to search for typescript files/project. Default "./"
	// 3. Ignore:  string, comma separated ignore typescript files/project from these directories. Default "" (node_modules are always ignored)
	// 4. Tsconfig:  &typescript.Tsconfig{}, here you can set all compilerOptions if no tsconfig.json exists inside the 'Dir'
	// 5. Editor: 	typescript.Editor("username","password"), if setted then alm-tools browser-based typescript IDE will be available. Defailt is nil
	Options struct {
		Bin      string
		Dir      string
		Ignore   string
		Tsconfig *Tsconfig
		Editor   *editor.Plugin // the editor is just a plugin also
	}
	// Plugin the struct of the Typescript Plugin, holds all necessary fields & methods
	Plugin struct {
		options Options
		// taken from Activate
		pluginContainer iris.IPluginContainer
		// taken at the PreListen
		logger *iris.Logger
	}
)

// Editor is just a shortcut for github.com/kataras/iris/plugin/editor.New()
// returns a new (Editor)Plugin, it's exists here because the typescript plugin has direct interest with the EditorPlugin
func Editor(username, password string) *editor.Plugin {
	return editor.New(username, password)
}

// DefaultOptions returns the default Options of the Plugin
func DefaultOptions() Options {
	root, err := os.Getwd()
	if err != nil {
		panic("Typescript Plugin: Cannot get the Current Working Directory !!! [os.getwd()]")
	}
	opt := Options{Dir: root + system.PathSeparator, Ignore: nodeModules, Tsconfig: DefaultTsconfig()}
	opt.Bin = npm.Abs("typescript/lib/tsc.js")
	return opt

}

// Plugin

// New creates & returns a new instnace typescript plugin
func New(_opt ...Options) *Plugin {
	var options = DefaultOptions()

	if _opt != nil && len(_opt) > 0 { //not nil always but I like this way :)
		opt := _opt[0]

		if opt.Bin != "" {
			options.Bin = opt.Bin
		}
		if opt.Dir != "" {
			options.Dir = opt.Dir
		}

		if !strings.Contains(opt.Ignore, nodeModules) {
			opt.Ignore += "," + nodeModules
		}

		if opt.Tsconfig != nil {
			options.Tsconfig = opt.Tsconfig
		}

		options.Ignore = opt.Ignore
	}

	return &Plugin{options: options}
}

// implement the IPlugin & IPluginPreListen

// Activate ...
func (t *Plugin) Activate(container iris.IPluginContainer) error {
	t.pluginContainer = container
	return nil
}

// GetName ...
func (t *Plugin) GetName() string {
	return Name + "[" + iris.RandomString(10) + "]" // this allows the specific plugin to be registed more than one time
}

// GetDescription TypescriptPlugin scans and compile typescript files with ease
func (t *Plugin) GetDescription() string {
	return Name + " scans and compile typescript files with ease. \n"
}

// PreListen ...
func (t *Plugin) PreListen(s *iris.Station) {
	t.logger = s.Logger()
	t.start()
}

//

// implementation

func (t *Plugin) start() {
	defaultCompilerArgs := t.options.Tsconfig.CompilerArgs() //these will be used if no .tsconfig found.
	if t.hasTypescriptFiles() {
		//Can't check if permission denied returns always exists = true....
		//typescriptModule := out + string(os.PathSeparator) + "typescript" + string(os.PathSeparator) + "bin"
		if !npm.Exists(t.options.Bin) {
			t.logger.Println("Installing typescript, please wait...")
			res := npm.Install("typescript")
			if res.Error != nil {
				t.logger.Print(res.Error.Error())
				return
			}
			t.logger.Print(res.Message)

		}

		projects := t.getTypescriptProjects()
		if len(projects) > 0 {
			watchedProjects := 0
			//typescript project (.tsconfig) found
			for _, project := range projects {
				cmd := system.CommandBuilder("node", t.options.Bin, "-p", project[0:strings.LastIndex(project, system.PathSeparator)]) //remove the /tsconfig.json)
				projectConfig := FromFile(project)

				if projectConfig.CompilerOptions.Watch {
					watchedProjects++
					// if has watch : true then we have to wrap the command to a goroutine (I don't want to use the .Start here)
					go func() {
						_, err := cmd.Output()
						if err != nil {
							t.logger.Println(err.Error())
							return
						}
					}()
				} else {

					_, err := cmd.Output()
					if err != nil {
						t.logger.Println(err.Error())
						return
					}

				}

			}
			t.logger.Printf("%d Typescript project(s) compiled ( %d monitored by a background file watcher ) ", len(projects), watchedProjects)
		} else {
			//search for standalone typescript (.ts) files and compile them
			files := t.getTypescriptFiles()

			if len(files) > 0 {
				watchedFiles := 0
				if t.options.Tsconfig.CompilerOptions.Watch {
					watchedFiles = len(files)
				}
				//it must be always > 0 if we came here, because of if hasTypescriptFiles == true.
				for _, file := range files {
					cmd := system.CommandBuilder("node", t.options.Bin)
					cmd.AppendArguments(defaultCompilerArgs...)
					cmd.AppendArguments(file)
					_, err := cmd.Output()
					cmd.Args = cmd.Args[0 : len(cmd.Args)-1] //remove the last, which is the file
					if err != nil {
						t.logger.Println(err.Error())
						return
					}

				}
				t.logger.Printf("%d Typescript file(s) compiled ( %d monitored by a background file watcher )", len(files), watchedFiles)
			}

		}

		//editor activation
		if len(projects) == 1 && t.options.Editor != nil {
			dir := projects[0][0:strings.LastIndex(projects[0], system.PathSeparator)]
			t.options.Editor.Dir(dir)
			t.pluginContainer.Plugin(t.options.Editor)
		}

	}
}

func (t *Plugin) hasTypescriptFiles() bool {
	root := t.options.Dir
	ignoreFolders := strings.Split(t.options.Ignore, ",")
	hasTs := false

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
		}
		for i := range ignoreFolders {
			if strings.Contains(path, ignoreFolders[i]) {
				return nil
			}
		}
		if strings.HasSuffix(path, ".ts") {
			hasTs = true
			return errors.New("Typescript found, hope that will stop here")
		}

		return nil
	})
	return hasTs
}

func (t *Plugin) getTypescriptProjects() []string {
	var projects []string
	ignoreFolders := strings.Split(t.options.Ignore, ",")

	root := t.options.Dir
	//t.logger.Printf("\nSearching for typescript projects in %s", root)

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		for i := range ignoreFolders {
			if strings.Contains(path, ignoreFolders[i]) {
				//t.logger.Println(path + " ignored")
				return nil
			}
		}

		if strings.HasSuffix(path, system.PathSeparator+"tsconfig.json") {
			//t.logger.Printf("\nTypescript project found in %s", path)
			projects = append(projects, path)
		}

		return nil
	})
	return projects
}

// this is being called if getTypescriptProjects return 0 len, then we are searching for files using that:
func (t *Plugin) getTypescriptFiles() []string {
	var files []string
	ignoreFolders := strings.Split(t.options.Ignore, ",")

	root := t.options.Dir
	//t.logger.Printf("\nSearching for typescript files in %s", root)

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		for i := range ignoreFolders {
			if strings.Contains(path, ignoreFolders[i]) {
				//t.logger.Println(path + " ignored")
				return nil
			}
		}

		if strings.HasSuffix(path, ".ts") {
			//t.logger.Printf("\nTypescript file found in %s", path)
			files = append(files, path)
		}

		return nil
	})
	return files
}

//
//
