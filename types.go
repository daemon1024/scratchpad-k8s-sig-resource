package main

import (
	"fmt"
	"sort"
	"strings"

	"k8s.io/test-infra/prow/github"
)

// LabelTarget specifies the intent of the label (PR or issue)
type LabelTarget string

// Label holds declarative data about the label.
type Label struct {
	// Name is the current name of the label
	Name string `json:"name"`
	// Color is rrggbb or color
	// Color string `json:"color"`
	// // Description is brief text explaining its meaning, who can apply it
	// Description string `json:"description"`
	// // Target specifies whether it targets PRs, issues or both
	// Target LabelTarget `json:"target"`
	// // ProwPlugin specifies which prow plugin add/removes this label
	// ProwPlugin string `json:"prowPlugin"`
	// // IsExternalPlugin specifies if the prow plugin is external or not
	// IsExternalPlugin bool `json:"isExternalPlugin"`
	// // AddedBy specifies whether human/munger/bot adds the label
	// AddedBy string `json:"addedBy"`
	// // Previously lists deprecated names for this label
	// Previously []Label `json:"previously,omitempty"`
	// // DeleteAfter specifies the label is retired and a safe date for deletion
	// DeleteAfter *time.Time `json:"deleteAfter,omitempty"`
	// parent      *Label     // Current name for previous labels (used internally)
}

// Configuration is a list of Repos defining Required Labels to sync into them
// There is also a Default list of labels applied to every Repo
type Configuration struct {
	Repos   map[string]RepoConfig `json:"repos,omitempty"`
	Orgs    map[string]RepoConfig `json:"orgs,omitempty"`
	Default RepoConfig            `json:"default"`
}

// RepoConfig contains only labels for the moment
type RepoConfig struct {
	Labels []Label `json:"labels"`
}

// RepoLabels holds a repo => []github.Label mapping.
type RepoLabels map[string][]github.Label

// Update a label in a repo
type Update struct {
	repo    string
	Why     string
	Wanted  *Label `json:"wanted,omitempty"`
	Current *Label `json:"current,omitempty"`
}

// RepoUpdates Repositories to update: map repo name --> list of Updates
type RepoUpdates map[string][]Update

// copied from generator/app.go

// Person represents an individual person holding a role in a group.
type Person struct {
	GitHub  string
	Name    string
	Company string `yaml:"company,omitempty"`
}

// Meeting represents a regular meeting for a group.
type Meeting struct {
	Description   string
	Day           string
	Time          string
	TZ            string
	Frequency     string
	URL           string `yaml:",omitempty"`
	ArchiveURL    string `yaml:"archive_url,omitempty"`
	RecordingsURL string `yaml:"recordings_url,omitempty"`
}

// Contact represents the various contact points for a group.
type Contact struct {
	Slack              string       `yaml:",omitempty"`
	MailingList        string       `yaml:"mailing_list,omitempty"`
	PrivateMailingList string       `yaml:"private_mailing_list,omitempty"`
	GithubTeams        []GithubTeam `yaml:"teams,omitempty"`
	Liaison            Person       `yaml:"liaison,omitempty"`
}

// GithubTeam represents a specific Github Team.
type GithubTeam struct {
	Name        string
	//Description string `yaml:",omitempty"`
}

// Subproject represents a specific subproject owned by the group
type Subproject struct {
	Name        string
	Description string   `yaml:",omitempty"`
	Contact     *Contact `yaml:",omitempty"`
	Owners      []string
	Meetings    []Meeting `yaml:",omitempty"`
}

// LeadershipGroup represents the different groups of leaders within a group
type LeadershipGroup struct {
	Chairs         []Person
	TechnicalLeads []Person `yaml:"tech_leads,omitempty"`
	EmeritusLeads  []Person `yaml:"emeritus_leads,omitempty"`
}

// PrefixToPersonMap returns a map of prefix to persons, useful for iteration over all persons
func (g *LeadershipGroup) PrefixToPersonMap() map[string][]Person {
	return map[string][]Person{
		"chair":         g.Chairs,
		"tech_lead":     g.TechnicalLeads,
		"emeritus_lead": g.EmeritusLeads,
	}
}

// Owners returns a sorted and de-duped list of owners for a LeadershipGroup
func (g *LeadershipGroup) Owners() []Person {
	o := append(g.Chairs, g.TechnicalLeads...)

	// Sort
	sort.Slice(o, func(i, j int) bool {
		return o[i].GitHub < o[j].GitHub
	})

	// De-dupe
	seen := make(map[string]struct{}, len(o))
	i := 0
	for _, p := range o {
		if _, ok := seen[p.GitHub]; ok {
			continue
		}
		seen[p.GitHub] = struct{}{}
		o[i] = p
		i++
	}
	return o[:i]
}

// FoldedString is a string that will be serialized in FoldedStyle by go-yaml
type FoldedString string

// Group represents either a Special Interest Group (SIG) or a Working Group (WG)
type Group struct {
	Dir              string
	Name             string
	MissionStatement FoldedString `yaml:"mission_statement,omitempty"`
	CharterLink      string       `yaml:"charter_link,omitempty"`
	StakeholderSIGs  []string     `yaml:"stakeholder_sigs,omitempty"`
	Label            string
	Leadership       LeadershipGroup `yaml:"leadership"`
	Meetings         []Meeting
	Contact          Contact
	Subprojects      []Subproject `yaml:",omitempty"`
}

// DirName returns the directory that a group's documentation will be
// generated into. It is composed of a prefix (sig for SIGs and wg for WGs),
// and a formatted version of the group's name (in kebab case).
func (g *Group) DirName(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, strings.ToLower(strings.Replace(g.Name, " ", "-", -1)))
}

// LabelName returns the expected label for a given group
func (g *Group) LabelName(prefix string) string {
	return strings.Replace(g.DirName(prefix), fmt.Sprintf("%s-", prefix), "", 1)
}

// Context is the context for the sigs.yaml file.
type Context struct {
	Sigs          []Group
	WorkingGroups []Group
	UserGroups    []Group
	Committees    []Group
}


// Team is the list it contains
type Team struct {

}

// TeamList is the struct for teams in various sig folders and its teams.yaml file
type TeamList struct {
	Teams		   []Team
}