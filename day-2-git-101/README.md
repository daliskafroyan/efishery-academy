# Day 2 - Git 101

a summary of github common command and workflow  from eFishery Academy 2022

## Git Command
### 1. `git init`
* initiate git repository

### 2. `git add <file>`
* use `git add .` to select all of the file you wanted to track
* or if you only want to track of the specific folder / file, use `git add <your-file>`. Path stated according to the folder where `.git` folder reside

### 3. `git commit`
* captures a shapshot of the project's currently staged changes
* usually used with adding a message explaining changes
* eg: `git commit -m "feat: initial commit"`

### 4. `git remote`
* Lets you create, view, and delete connections to other repositories. Remote connections are more like bookmarks rather than direct links into other repositories
* eg : `git remote add <remote name> <remote URL>`

### 5. `git status`
* The git status command displays the state of the working directory and the staging area. It lets you see which changes have been staged, which haven’t, and which files aren’t being tracked by Git.

### 6. `git fetch`
* Command to retrieve the latest meta-data info from the original to your local git. Yet it doesn’t do any file transferring. It’s more like just checking to see if there are any changes available
* eg: `git fetch all`

### 7. `git pull`
* The git pull command is used to fetch and download content from a remote repository and immediately update the local repository to match that content
* eg: `git pull origin master`

### 8. `git push`
* Send the committed changes to remote repositories for collaboration
* eg: `git push origin master`

### 8. `git branch`
* check branch available on local
* eg: `git branch`


## Commit Convention
````bash
<type>(<scope>): <subject>

<body>

<footer>
````
example :
````bash
fix(middleware): ensure Range headers adhere more closely to RFC 2616 

Add one new dependency, use `range-parser` (Express dependency) to compute range. It is more well-tested in the wild.

Fixes #2310

````

* you can follow this pattern from conventional commit homepage

## Semantic Versioning
`v<major>.<minor>.<patch>`
Patch: fix, perf
Minor:  feat
Major: breaking changes in the API 

## Git Management
* Using Trunk Based Development
* A source-control branching model, where developers collaborate on code in a single branch called ‘trunk’ (main or master)
* Avoid long-lived branches
* Use short-lived branches only for feature development
* Merge the pull request only if it is already reviewed
* Each and every commit made to the trunk is considered as a deployable unit

