package types

// AddPolicy response from adding a policy
type AddPolicy struct {
	PolicyID string `json:"policyId"`
}

// List policies response from listing policies
type ListPolicies struct {
	Items      []Items `json:"items,omitempty"`
	TotalCount int     `json:"totalCount,omitempty"`
}
type Items struct {
	PolicyID    string   `json:"policyId,omitempty"`
	Status      string   `json:"status,omitempty"`
	PolicyName  string   `json:"policyName,omitempty"`
	Description string   `json:"description,omitempty"`
	UpdatedOn   string   `json:"updatedOn,omitempty"`
	RuleNames   []string `json:"ruleNames,omitempty"`
	Platforms   []string `json:"platforms,omitempty"`
}

// Get policy response from getting a policy
type Policy struct {
	PolicyName      string            `json:"policyName,omitempty"`
	Status          string            `json:"status,omitempty"`
	Description     string            `json:"description,omitempty"`
	ProvidersData   ProvidersData     `json:"providersData,omitempty"`
	StartDate       string            `json:"startDate,omitempty"`
	EndDate         string            `json:"endDate,omitempty"`
	UserAccessRules []UserAccessRules `json:"userAccessRules,omitempty"`
}
type Tags struct {
	Key   string   `json:"Key,omitempty"`
	Value []string `json:"Value,omitempty"`
}
type Aws struct {
	Regions    []string `json:"regions,omitempty"`
	Tags       []Tags   `json:"tags,omitempty"`
	VpcIds     []string `json:"vpcIds,omitempty"`
	AccountIds []string `json:"accountIds,omitempty"`
}
type Azure struct {
	Regions        []string `json:"regions,omitempty"`
	Tags           []Tags   `json:"tags,omitempty"`
	ResourceGroups []string `json:"resourceGroups,omitempty"`
	VnetIds        []string `json:"vnetIds,omitempty"`
	Subscriptions  []string `json:"subscriptions,omitempty"`
}
type FqdnRules struct {
	Operator            string `json:"operator,omitempty"`
	ComputernamePattern string `json:"computernamePattern,omitempty"`
	Domain              string `json:"domain,omitempty"`
}
type OnPrem struct {
	FqdnRulesConjunction string      `json:"fqdnRulesConjunction,omitempty"`
	FqdnRules            []FqdnRules `json:"fqdnRules,omitempty"`
}
type Labels struct {
	Key   string   `json:"Key,omitempty"`
	Value []string `json:"Value,omitempty"`
}
type Gcp struct {
	Regions  []string `json:"regions,omitempty"`
	Labels   []Labels `json:"labels,omitempty"`
	VpcIds   []string `json:"vpc_ids,omitempty"`
	Projects []string `json:"projects,omitempty"`
}
type ProvidersData struct {
	Aws    Aws    `json:"AWS,omitempty"`
	Azure  Azure  `json:"Azure,omitempty"`
	OnPrem OnPrem `json:"OnPrem,omitempty"`
	Gcp    Gcp    `json:"GCP,omitempty"`
}
type Roles struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type Groups struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type Users struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type UserData struct {
	Roles  []Roles  `json:"Roles,omitempty"`
	Groups []Groups `json:"Groups,omitempty"`
	Users  []Users  `json:"users,omitempty"`
}
type LocalEphemeralUser struct {
	AssignGroups []string `json:"assignGroups,omitempty"`
}
type Rdp struct {
	LocalEphemeralUser LocalEphemeralUser `json:"localEphemeralUser,omitempty"`
}
type ConnectAsAws struct {
	SSH string `json:"ssh,omitempty"`
	Rdp Rdp    `json:"rdp,omitempty"`
}
type ConnectAsAzure struct {
	SSH string `json:"ssh,omitempty"`
}
type ConnectAsOnPrem struct {
	Rdp Rdp `json:"rdp,omitempty"`
}
type ConnectAsGcp struct {
	SSH string `json:"ssh,omitempty"`
}
type ConnectAs struct {
	Aws    ConnectAsAws    `json:"AWS,omitempty"`
	Azure  ConnectAsAzure  `json:"Azure,omitempty"`
	OnPrem ConnectAsOnPrem `json:"OnPrem,omitempty"`
	Gcp    ConnectAsGcp    `json:"GCP,omitempty"`
}
type ConnectionInformation struct {
	ConnectAs   ConnectAs `json:"connectAs,omitempty"`
	GrantAccess int       `json:"grantAccess,omitempty"`
	IdleTime    int       `json:"idleTime,omitempty"`
	DaysOfWeek  []string  `json:"daysOfWeek,omitempty"`
	FullDays    bool      `json:"fullDays,omitempty"`
	HoursFrom   string    `json:"hoursFrom,omitempty"`
	HoursTo     string    `json:"hoursTo,omitempty"`
	TimeZone    string    `json:"timeZone,omitempty"`
}
type UserAccessRules struct {
	RuleName              string                `json:"ruleName,omitempty"`
	UserData              UserData              `json:"userData,omitempty"`
	ConnectionInformation ConnectionInformation `json:"connectionInformation,omitempty"`
}
