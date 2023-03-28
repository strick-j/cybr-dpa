package types

type Policies []struct {
	PolicyID    string   `json:"policyId"`
	Status      string   `json:"status"`
	PolicyName  string   `json:"policyName"`
	Description string   `json:"description"`
	UpdatedOn   string   `json:"updatedOn"`
	RuleNames   []string `json:"ruleNames"`
	Platforms   []string `json:"platforms"`
}
