package gissue


const(
	token = "05bdc375fd240bcaf3a81a5bdee36dd008ffd6e2"
	rootUrl = "https://api.github.com"
	repoURI = rootUrl + "/repos"
	searchIssue = rootUrl + "/search/issues"
	repoIssues = repoURI + "/%s/%s/issues"
	singleIssue = repoIssues + "/%d"
	headerAccept = "application/vnd.github.symmetra-preview+json"
)

type ListResult []Issue

// gen by another project : https://github.com/dudd/json2struct
type SearchResult struct {
	TotalCount               int64          `json:"total_count"`
	Items                    []Issue        `json:"items"`
	IncompleteResults        bool           `json:"incomplete_results"`
}

type PullRequest struct {
	Url                      string         `json:"url,omitempty"`
	DiffUrl                  string         `json:"diff_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	PatchUrl                 string         `json:"patch_url,omitempty"`
}

type Creator struct {
	FollowingUrl             string         `json:"following_url,omitempty"`
	EventsUrl                string         `json:"events_url,omitempty"`
	AvatarUrl                string         `json:"avatar_url,omitempty"`
	Url                      string         `json:"url,omitempty"`
	GistsUrl                 string         `json:"gists_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	SubscriptionsUrl         string         `json:"subscriptions_url,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ReposUrl                 string         `json:"repos_url,omitempty"`
	ReceivedEventsUrl        string         `json:"received_events_url,omitempty"`
	GravatarId               string         `json:"gravatar_id,omitempty"`
	StarredUrl               string         `json:"starred_url,omitempty"`
	SiteAdmin                bool           `json:"site_admin"`
	Login                    string         `json:"login,omitempty"`
	Type                     string         `json:"type,omitempty"`
	Id                       int64          `json:"id"`
	FollowersUrl             string         `json:"followers_url,omitempty"`
	OrganizationsUrl         string         `json:"organizations_url,omitempty"`
}

type Assignees struct {
	FollowingUrl             string         `json:"following_url,omitempty"`
	EventsUrl                string         `json:"events_url,omitempty"`
	AvatarUrl                string         `json:"avatar_url,omitempty"`
	Url                      string         `json:"url,omitempty"`
	GistsUrl                 string         `json:"gists_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	SubscriptionsUrl         string         `json:"subscriptions_url,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ReposUrl                 string         `json:"repos_url,omitempty"`
	ReceivedEventsUrl        string         `json:"received_events_url,omitempty"`
	GravatarId               string         `json:"gravatar_id,omitempty"`
	StarredUrl               string         `json:"starred_url,omitempty"`
	SiteAdmin                bool           `json:"site_admin"`
	Login                    string         `json:"login,omitempty"`
	Type                     string         `json:"type,omitempty"`
	Id                       int64          `json:"id"`
	FollowersUrl             string         `json:"followers_url,omitempty"`
	OrganizationsUrl         string         `json:"organizations_url,omitempty"`
}

type Labels struct {
	Default                  bool           `json:"default"`
	Description              string         `json:"description,omitempty"`
	Url                      string         `json:"url,omitempty"`
	Color                    string         `json:"color,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	Id                       int64          `json:"id"`
	Name                     string         `json:"name,omitempty"`
}

type ClosedBy struct {
	FollowingUrl             string         `json:"following_url,omitempty"`
	EventsUrl                string         `json:"events_url,omitempty"`
	AvatarUrl                string         `json:"avatar_url,omitempty"`
	Url                      string         `json:"url,omitempty"`
	GistsUrl                 string         `json:"gists_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	SubscriptionsUrl         string         `json:"subscriptions_url,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ReposUrl                 string         `json:"repos_url,omitempty"`
	ReceivedEventsUrl        string         `json:"received_events_url,omitempty"`
	GravatarId               string         `json:"gravatar_id,omitempty"`
	StarredUrl               string         `json:"starred_url,omitempty"`
	SiteAdmin                bool           `json:"site_admin"`
	Login                    string         `json:"login,omitempty"`
	Type                     string         `json:"type,omitempty"`
	Id                       int64          `json:"id"`
	FollowersUrl             string         `json:"followers_url,omitempty"`
	OrganizationsUrl         string         `json:"organizations_url,omitempty"`
}

type Assignee struct {
	FollowingUrl             string         `json:"following_url,omitempty"`
	EventsUrl                string         `json:"events_url,omitempty"`
	AvatarUrl                string         `json:"avatar_url,omitempty"`
	Url                      string         `json:"url,omitempty"`
	GistsUrl                 string         `json:"gists_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	SubscriptionsUrl         string         `json:"subscriptions_url,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ReposUrl                 string         `json:"repos_url,omitempty"`
	ReceivedEventsUrl        string         `json:"received_events_url,omitempty"`
	GravatarId               string         `json:"gravatar_id,omitempty"`
	StarredUrl               string         `json:"starred_url,omitempty"`
	SiteAdmin                bool           `json:"site_admin"`
	Login                    string         `json:"login,omitempty"`
	Type                     string         `json:"type,omitempty"`
	Id                       int64          `json:"id"`
	FollowersUrl             string         `json:"followers_url,omitempty"`
	OrganizationsUrl         string         `json:"organizations_url,omitempty"`
}

type User struct {
	FollowingUrl             string         `json:"following_url,omitempty"`
	EventsUrl                string         `json:"events_url,omitempty"`
	AvatarUrl                string         `json:"avatar_url,omitempty"`
	Url                      string         `json:"url,omitempty"`
	GistsUrl                 string         `json:"gists_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	SubscriptionsUrl         string         `json:"subscriptions_url,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ReposUrl                 string         `json:"repos_url,omitempty"`
	ReceivedEventsUrl        string         `json:"received_events_url,omitempty"`
	GravatarId               string         `json:"gravatar_id,omitempty"`
	StarredUrl               string         `json:"starred_url,omitempty"`
	SiteAdmin                bool           `json:"site_admin"`
	Login                    string         `json:"login,omitempty"`
	Type                     string         `json:"type,omitempty"`
	Id                       int64          `json:"id"`
	FollowersUrl             string         `json:"followers_url,omitempty"`
	OrganizationsUrl         string         `json:"organizations_url,omitempty"`
}

type Milestone struct {
	Description              string         `json:"description,omitempty"`
	Title                    string         `json:"title,omitempty"`
	Url                      string         `json:"url,omitempty"`
	LabelsUrl                string         `json:"labels_url,omitempty"`
	CreatedAt                string         `json:"created_at,omitempty"`
	Creator                  Creator        `json:"creator"`
	Number                   int64          `json:"number"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	UpdatedAt                string         `json:"updated_at,omitempty"`
	DueOn                    string         `json:"due_on,omitempty"`
	State                    string         `json:"state,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	ClosedIssues             int64          `json:"closed_issues"`
	OpenIssues               int64          `json:"open_issues"`
	ClosedAt                 string         `json:"closed_at,omitempty"`
	Id                       int64          `json:"id"`
}

type Issue struct {
	ActiveLockReason         string         `json:"active_lock_reason,omitempty"`
	Labels                   []Labels       `json:"labels"`
	Number                   int64          `json:"number"`
	Assignee                 Assignee       `json:"assignee"`
	RepositoryUrl            string         `json:"repository_url,omitempty"`
	ClosedAt                 string         `json:"closed_at,omitempty"`
	Id                       int64          `json:"id"`
	Title                    string         `json:"title,omitempty"`
	Comments                 int64          `json:"comments"`
	State                    string         `json:"state,omitempty"`
	Body                     string         `json:"body,omitempty"`
	LabelsUrl                string         `json:"labels_url,omitempty"`
	Milestone                Milestone      `json:"milestone"`
	EventsUrl                string         `json:"events_url,omitempty"`
	CommentsUrl              string         `json:"comments_url,omitempty"`
	HtmlUrl                  string         `json:"html_url,omitempty"`
	UpdatedAt                string         `json:"updated_at,omitempty"`
	NodeId                   string         `json:"node_id,omitempty"`
	User                     User           `json:"user"`
	PullRequest              PullRequest    `json:"pull_request"`
	ClosedBy                 ClosedBy       `json:"closed_by"`
	Locked                   bool           `json:"locked"`
	Url                      string         `json:"url,omitempty"`
	CreatedAt                string         `json:"created_at,omitempty"`
	Assignees                []Assignees    `json:"assignees"`
	Score					 float64		`json:"score,omitempty"`
}

type IssueReq struct {
	Title                    string         `json:"title"`
	Labels                   []interface{}  `json:"labels"`
	Body                     string         `json:"body"`
	State                    string         `json:"state,omitempty"`
	Assignee                 interface{}    `json:"assignee,omitempty"`
	Assignees                []string       `json:"assignees"`
	Milestone                interface{}    `json:"milestone,omitempty"`
}

