package rules

import (
	"fmt"
	"log"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

// McafModuleUsageRule checks whether ...
type McafModuleUsageRule struct {
	tflint.DefaultRule
}

// NewMcafModuleUsageRule returns a new rule
func NewMcafModuleUsageRule() *McafModuleUsageRule {
	return &McafModuleUsageRule{}
}

// Name returns the rule name
func (r *McafModuleUsageRule) Name() string {
	return "mcaf_module_usage"
}

// Enabled returns whether the rule is enabled by default
func (r *McafModuleUsageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *McafModuleUsageRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *McafModuleUsageRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether ...
func (r *McafModuleUsageRule) Check(rr tflint.Runner) error {
	runner := rr.(*terraform.Runner)

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	calls, diags := runner.GetModuleCalls()
	if diags.HasErrors() {
		return diags
	}

	for _, call := range calls {
		if err := r.checkModule(runner, call); err != nil {
			return err
		}
	}

	return nil
}

func (r *McafModuleUsageRule) checkModule(runner tflint.Runner, module *terraform.ModuleCall) error {

	modules := map[string]string{
		"terraform-aws-modules/s3-bucket/aws": "terraform-aws-mcaf-s3",
	}

	log.Printf("[DEBUG] Checking module: %s", module.Name)

	for module_source, mcaf_module := range modules {
		if strings.Contains(module.Source, module_source) {
			return runner.EmitIssue(
				r,
				fmt.Sprintf("module %s should be replaced with MCAF module %s", module.Name, mcaf_module),
				module.SourceAttr.Range,
			)
		}
	}

	return nil
}
