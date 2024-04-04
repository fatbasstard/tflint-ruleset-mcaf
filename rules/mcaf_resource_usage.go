package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-terraform/project"
	"github.com/terraform-linters/tflint-ruleset-terraform/terraform"
)

// McafResourceUsageRule checks whether ...
type McafResourceUsageRule struct {
	tflint.DefaultRule
}

// NewMcafResourceUsageRule returns a new rule
func NewMcafResourceUsageRule() *McafResourceUsageRule {
	return &McafResourceUsageRule{}
}

// Name returns the rule name
func (r *McafResourceUsageRule) Name() string {
	return "mcaf_resource_usage"
}

// Enabled returns whether the rule is enabled by default
func (r *McafResourceUsageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *McafResourceUsageRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *McafResourceUsageRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether ...
func (r *McafResourceUsageRule) Check(rr tflint.Runner) error {
	runner := rr.(*terraform.Runner)

	resource_types := map[string]string{
		"aws_s3_bucket": "terraform-aws-mcaf-s3",
	}

	for resource_type, mcaf_module := range resource_types {
		resources, err := runner.GetResourceContent(resource_type, &hclext.BodySchema{}, nil)
		if err != nil {
			return err
		}

		if resources.IsEmpty() {
			continue
		}

		for _, resource := range resources.Blocks {
			runner.EmitIssue(
				r,
				fmt.Sprintf("resource %s should be replaced with MCAF module %s", resource.Labels[0]+"."+resource.Labels[1], mcaf_module),
				resource.DefRange,
			)
		}
	}

	return nil
}
