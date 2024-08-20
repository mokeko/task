package task

import (
	"fmt"

	"github.com/go-task/task/v3/errors"
	"github.com/go-task/task/v3/taskfile/ast"
)

func (e *Executor) areTaskRequiredVarsSet(t *ast.Task, call *ast.Call) error {
	requires, err := e.Compiler.GetRequires(t, call)
	if err != nil {
		return err
	}
	fmt.Printf("merged requires: %v\n", requires.Vars)
	if len(requires.Vars) == 0 {
		return nil
	}

	vars, err := e.Compiler.GetVariables(t, call)
	if err != nil {
		return err
	}

	var missingVars []string
	for _, requiredVar := range requires.Vars {
		if !vars.Exists(requiredVar) {
			missingVars = append(missingVars, requiredVar)
		}
	}

	if len(missingVars) > 0 {
		return &errors.TaskMissingRequiredVars{
			TaskName:    t.Name(),
			MissingVars: missingVars,
		}
	}

	return nil
}
