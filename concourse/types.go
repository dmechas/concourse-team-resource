package concourse

type Source struct {
	Target   string `json:"target"`
	Insecure string `json:"insecure"`
	Username string `json:"username"`
	Password string `json:"password"`
	MainTeam string `json:"main_team"`
}

type CheckRequest struct {
}

type Version map[string]string

type CheckResponse []Version

type InRequest struct {
}

type InResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}

type Metadata struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type OutRequest struct {
	Source Source    `json:"source"`
	Params OutParams `json:"params"`
}

type OutParams struct {
	GitHubTeam string `json:"github_team" yaml:"github_team"`
	LocalUser  string `json:"local_user" yaml:"local_user"`
	TeamName   string `json:"team_name" yaml:"team_name"`
}

type OutResponse struct {
	Version  Version    `json:"version"`
	Metadata []Metadata `json:"metadata"`
}
