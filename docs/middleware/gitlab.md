# go-gitlab
github：https://github.com/xanzy/go-gitlab

## get access token
1. login gitlab
2. User Settings - Access Token
3. Add a personal access token
   name: zhongqiong
   Expires at: null is never expires
4. Scopes
   select api


## import package
import "github.com/xanzy/go-gitlab"

## download package
go mod tidy

## use password auth

>
```
	git, err := gitlab.NewBasicAuthClient(nil, "https://gitlab.icarbonx.cn", "your username", "your password")
	if err != nil {
		log.Fatal(err)
	}

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d projects", projects)
```


## use access token auth


```
	git := gitlab.NewClient(nil, "your personal access token")
	// access should be select api scope
	git.SetBaseURL("https://gitlab.icarbonx.cn")

	// List all projects
	projects, _, err := git.Projects.ListProjects(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d projects", projects)
```

