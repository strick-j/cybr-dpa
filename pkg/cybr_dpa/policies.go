package cybr_dpa

import (
	"context"
	"fmt"
)

// Struct for the response from GET /api/access-policies
type Policies struct {
	Items []struct {
		PolicyID    string   `json:"policyId"`
		Status      string   `json:"status"`
		PolicyName  string   `json:"policyName"`
		Description string   `json:"description"`
		UpdatedOn   string   `json:"updatedOn"`
		RuleNames   []string `json:"ruleNames"`
		Platforms   []string `json:"platforms"`
	} `json:"items"`
	TotalCount int `json:"totalCount"`
}

// Struct for the response from GET, PATCH, PUT /api/access-policies/{policy-id} and POST /api/access-polcies
type Policy struct {
	PolicyName    string `json:"policyName,omitempty"`
	Status        string `json:"status,omitempty"`
	Description   string `json:"description,omitempty"`
	ProvidersData struct {
		Aws struct {
			Regions []string `json:"regions,omitempty"`
			Tags    []struct {
				Key   string   `json:"Key,omitempty"`
				Value []string `json:"Value,omitempty"`
			} `json:"tags,omitempty"`
			VpcIds     []any `json:"vpcIds,omitempty"`
			AccountIds []any `json:"accountIds,omitempty"`
		} `json:"AWS,omitempty"`
		Azure struct {
			Regions []string `json:"regions,omitempty"`
			Tags    []struct {
				Key   string   `json:"Key,omitempty"`
				Value []string `json:"Value,omitempty"`
			} `json:"tags,omitempty"`
			ResourceGroups []any `json:"resourceGroups,omitempty"`
			VnetIds        []any `json:"vnetIds,omitempty"`
			Subscriptions  []any `json:"subscriptions,omitempty"`
		} `json:"Azure,omitempty"`
		OnPrem struct {
			FqdnRulesConjunction string `json:"fqdnRulesConjunction,omitempty"`
			FqdnRules            []struct {
				Operator            string `json:"operator,omitempty"`
				ComputernamePattern string `json:"computernamePattern,omitempty"`
				Domain              string `json:"domain,omitempty"`
			} `json:"fqdnRules,omitempty"`
		} `json:"OnPrem,omitempty"`
	} `json:"providersData,omitempty"`
	StartDate       string `json:"startDate,omitempty"`
	EndDate         string `json:"endDate,omitempty"`
	UserAccessRules []struct {
		RuleName string `json:"ruleName,omitempty"`
		UserData struct {
			Roles []struct {
				Name   string `json:"name,omitempty"`
				Source string `json:"source,omitempty"`
			} `json:"roles,omitempty"`
			Groups []any `json:"groups,omitempty"`
			Users  []any `json:"users,omitempty"`
		} `json:"userData,omitempty"`
		ConnectionInformation struct {
			ConnectAs struct {
				Aws struct {
					SSH string `json:"ssh,omitempty"`
					Rdp struct {
						LocalEphemeralUser struct {
							AssignGroups []string `json:"assignGroups,omitempty"`
						} `json:"localEphemeralUser,omitempty"`
					} `json:"rdp,omitempty"`
				} `json:"AWS,omitempty"`
				Azure struct {
					SSH string `json:"ssh,omitempty"`
				} `json:"Azure,omitempty"`
				OnPrem struct {
					Rdp struct {
						LocalEphemeralUser struct {
							AssignGroups []string `json:"assignGroups,omitempty"`
						} `json:"localEphemeralUser,omitempty"`
					} `json:"rdp,omitempty"`
				} `json:"OnPrem,omitempty"`
			} `json:"connectAs,omitempty"`
			GrantAccess int      `json:"grantAccess,omitempty"`
			IdleTime    int      `json:"idleTime,omitempty"`
			DaysOfWeek  []string `json:"daysOfWeek,omitempty"`
			FullDays    bool     `json:"fullDays,omitempty"`
			HoursFrom   string   `json:"hoursFrom,omitempty"`
			HoursTo     string   `json:"hoursTo,omitempty"`
			TimeZone    string   `json:"timeZone,omitempty"`
		} `json:"connectionInformation,omitempty"`
	} `json:"userAccessRules,omitempty"`
}

var (
	policies Policies
	policy   Policy
)

// GetPolicies: Returns all authorization policies
//
// Example Usage:
//
//	getPolicies, err := s.GetPolicies(context.Background)
func (s *Service) GetPolicies(ctx context.Context) (*Policies, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s", "access-policies"), &policies); err != nil {
		return nil, fmt.Errorf("failed to get policies: %w", err)
	}

	return &policies, nil
}

// GetPolicyById: Retrieves authorization policy for the given ID
//
// Example Usage:
//
//	getPolicyById err := s.GetPolicyById(context.Background, "{policy_id}")
func (s *Service) GetPolicyById(ctx context.Context, policyId string) (*Policy, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "access-policies", policyId), &policy); err != nil {
		return nil, fmt.Errorf("failed to get policy %s: %w", policyId, err)
	}

	return &policy, nil
}

// PostPolicy: Adds a new authorization policy
//
// Example Usage:
//
//	policyDetails := Policy {
//	  PolicyName:  "ExamplePolicy",
//	  Status:      "Enabled",
//	  Description: "Example Policy",
//	  ProvidersData.
//	}
//
//	getPolicyById err := s.GetPolicyById(context.Background, policyDetails)
func (s *Service) PostPolicy(ctx context.Context, newPolicy Policy) (*Policy, error) {
	if err := s.client.Post(ctx, fmt.Sprintf("/%s", "access-policies"), newPolicy, &policy); err != nil {
		return nil, fmt.Errorf("failed to create new policy: %w", err)
	}

	return &policy, nil
}

// DeletePolicyById: Deletes the specified policy
//
// Example Usage:
//
//	err := s.DeletePolicyById(context.Background, "{policy_id}")
func (s *Service) DeletePolicyById(ctx context.Context, policyId string) error {
	if err := s.client.Delete(ctx, fmt.Sprintf("/%s/%s", "access-policies", policyId), nil); err != nil {
		return fmt.Errorf("failed to delete policy %s: %w", policyId, err)
	}

	return nil
}
