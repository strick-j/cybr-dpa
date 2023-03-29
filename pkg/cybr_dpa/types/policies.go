package types

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
		Aws *struct {
			Regions []string `json:"regions,omitempty"`
			Tags    []struct {
				Key   string   `json:"Key,omitempty"`
				Value []string `json:"Value,omitempty"`
			} `json:"tags,omitempty"`
			VpcIds     []any `json:"vpcIds,omitempty"`
			AccountIds []any `json:"accountIds,omitempty"`
		} `json:"AWS,omitempty"`
		Azure *struct {
			Regions []string `json:"regions,omitempty"`
			Tags    []struct {
				Key   string   `json:"Key,omitempty"`
				Value []string `json:"Value,omitempty"`
			} `json:"tags,omitempty"`
			ResourceGroups []any `json:"resourceGroups,omitempty"`
			VnetIds        []any `json:"vnetIds,omitempty"`
			Subscriptions  []any `json:"subscriptions,omitempty"`
		} `json:"Azure,omitempty"`
		OnPrem *struct {
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
			Roles  []string `json:"roles,omitempty"`
			Groups []any    `json:"groups,omitempty"`
			Users  []any    `json:"users,omitempty"`
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
