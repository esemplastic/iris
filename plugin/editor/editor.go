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

package editor

/* Notes for Auth
The Authorization specifies the authentication mechanism (in this case Basic) followed by the username and password.
Although, the string aHR0cHdhdGNoOmY= may look encrypted it is simply a base64 encoded version of <username>:<password>.
Would be readily available to anyone who could intercept the HTTP request.
*/

import (
	"os"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/logger"
	"github.com/kataras/iris/npm"
	"github.com/kataras/iris/utils"
)

const (
	// Name the name of the Plugin, which is "EditorPlugin"
	Name = "EditorPlugin"
)

type (
	// Plugin is an Editor Plugin the struct which implements the iris.IPlugin
	// it holds a logger from the iris' station
	// username,password for basic auth
	// directory which the client side code is
	// keyfile,certfile for TLS listening
	// and a host which is listening for
	Plugin struct {
		logger             *logger.Logger
		enabled            bool   // default true
		host               string // default 127.0.0.1
		port               int    // default 4444
		username, password string // based on Basic Auth, // default -nothing, for security reasons you have to set it otherwise editor is not opening.
		keyfile            string
		certfile           string
		directory          string // working directory

		// after alm started
		process *os.Process
	}
)

// New creates and returns a new (Editor)Plugin object
// accepts username and password, these are not optionally, you have to use that otherwise the editor will never actual run
// for security reasons and only
func New(username string, password string) *Plugin {
	e := &Plugin{enabled: true, port: 4444}
	e.username = username
	e.password = password
	return e
}

// User set a user, accepts two parameters: username (string), string (string)
func (e *Plugin) User(username string, password string) *Plugin {
	e.username = username
	e.password = password
	return e
}

// Dir sets the directory which the client side source code alive
func (e *Plugin) Dir(workingDir string) *Plugin {
	e.directory = workingDir
	return e
}

// Port sets the port (int) for the editor plugin's standalone server
func (e *Plugin) Port(port int) *Plugin {
	e.port = port
	return e
}

//

// SetEnable if true enables the editor plugin, otherwise disables it
func (e *Plugin) SetEnable(enable bool) {
	e.enabled = enable
}

// implement the IPlugin, IPluginPreListen & IPluginPreClose

// Activate ...
func (e *Plugin) Activate(container iris.IPluginContainer) error {
	return nil
}

// GetName returns the name of the Plugin
func (e *Plugin) GetName() string {
	return Name
}

// GetDescription EditorPlugin is a bridge between Iris and the alm-tools, the browser-based IDE for client-side sources.
func (e *Plugin) GetDescription() string {
	return Name + " is a bridge between Iris and the alm-tools, the browser-based IDE for client-side sources. \n"
}

// PreListen runs before the server's listens, saves the keyfile,certfile and the host from the Iris station to listen for
func (e *Plugin) PreListen(s *iris.Iris) {
	e.logger = s.Logger()
	e.keyfile = s.Server().Config.KeyFile
	e.certfile = s.Server().Config.CertFile
	e.host = s.Server().Config.ListeningAddr

	if idx := strings.Index(e.host, ":"); idx >= 0 {
		e.host = e.host[0:idx]
	}
	if e.host == "" {
		e.host = "127.0.0.1"
	}

	e.start()
}

// PreClose kills the editor's server when Iris is closed
func (e *Plugin) PreClose(s *iris.Iris) {
	if e.process != nil {
		err := e.process.Kill()
		if err != nil {
			e.logger.Printf("\nError while trying to terminate the (Editor)Plugin, please kill this process by yourself, process id: %d", e.process.Pid)
		}
	}
}

// start starts the job
func (e *Plugin) start() {

	if e.username == "" || e.password == "" {
		e.logger.Println("Error before running alm-tools. You have to set username & password for security reasons, otherwise this plugin won't run.")
		return
	}

	if !npm.Exists("alm/bin/alm") {
		e.logger.Println("Installing alm-tools, please wait...")
		res := npm.Install("alm")
		if res.Error != nil {
			e.logger.Print(res.Error.Error())
			return
		}
		e.logger.Print(res.Message)
	}

	cmd := utils.CommandBuilder("node", npm.Abs("alm/src/server.js"))
	cmd.AppendArguments("-a", e.username+":"+e.password, "-h", e.host, "-t", strconv.Itoa(e.port), "-d", e.directory[0:len(e.directory)-1])
	// for auto-start in the browser: cmd.AppendArguments("-o")
	if e.keyfile != "" && e.certfile != "" {
		cmd.AppendArguments("--httpskey", e.keyfile, "--httpscert", e.certfile)
	}

	//For debug only:
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	//os.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		e.logger.Println("Error while running alm-tools. Trace: " + err.Error())
		return
	}

	//we lose the internal error handling but ok...
	e.logger.Printf("Editor is running at %s:%d | %s", e.host, e.port, e.directory)

}
