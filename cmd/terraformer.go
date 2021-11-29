package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type printfer interface {
	Printf(format string, v ...interface{})
}

// Terraformer represents the Terraformer CLI executable and working directory.
//
type Terraformer struct {
	execPath   string
	workingDir string
	env        map[string]string

	stdout  io.Writer
	stderr  io.Writer
	logger  printfer
	logPath string
}

// NewTerraform returns a Terraformer struct with default values for all fields.
func NewTerraformer(workingDir string, execPath string) (*Terraformer, error) {
	if workingDir == "" {
		return nil, fmt.Errorf("Terraformer cannot be initialised with empty workdir")
	}

	if _, err := os.Stat(workingDir); err != nil {
		return nil, fmt.Errorf("error initialising Terraform with workdir %s: %s", workingDir, err)
	}

	if execPath == "" {
		err := fmt.Errorf("NewTerraformerer: please supply the path to a Terraformer executable using execPath.")
		return nil, &ErrNoSuitableBinary{
			err: err,
		}
	}
	tf := Terraformer{
		execPath:   execPath,
		workingDir: workingDir,
		env:        nil, // explicit nil means copy os.Environ
		logger:     log.New(ioutil.Discard, "", 0),
	}

	return &tf, nil
}

// SetLogger specifies a logger for tfexec to use.
func (tf *Terraformer) SetLogger(logger printfer) {
	tf.logger = logger
}

// SetStdout specifies a writer to stream stdout to for every command.
//
// This should be used for information or logging purposes only, not control
// flow. Any parsing necessary should be added as functionality to this package.
func (tf *Terraformer) SetStdout(w io.Writer) {
	tf.stdout = w
}

// SetStderr specifies a writer to stream stderr to for every command.
//
// This should be used for information or logging purposes only, not control
// flow. Any parsing necessary should be added as functionality to this package.
func (tf *Terraformer) SetStderr(w io.Writer) {
	tf.stderr = w
}

// SetLogPath sets the TF_LOG_PATH environment variable for Terraform CLI
// execution.
func (tf *Terraformer) SetLogPath(path string) error {
	tf.logPath = path
	return nil
}

// WorkingDir returns the working directory for Terraform.
func (tf *Terraformer) WorkingDir() string {
	return tf.workingDir
}

// ExecPath returns the path to the Terraform executable.
func (tf *Terraformer) ExecPath() string {
	return tf.execPath
}
