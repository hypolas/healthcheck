package main

type Builds struct {
	CommunityModules map[string]Modules `yaml:"community_modules"`
}

type Modules struct {
	Import  string `yaml:"import"`
	Version string `yaml:"version"`
}
