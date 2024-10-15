# Gitrocket

## Web-application to provide CI

### Stack
- Go 1.22
- SQLx
- go-migrate
- PostgreSQL
- TS + React for frontend application
- go-task

### Roadmap
- [ ] Load changes from git by time / by event and run in docker
- [ ] Public API via Swagger
- [ ] Authentication from database
- [ ] Simple frontend application
- [ ] Extended rules for Dockerfiles and configuration files
- [ ] Extended git-repos API support
- [ ] Additional sources of authentication via SSO
- [ ] Pipeline configuration scripts
- [ ] Additional pipeline workers (K8s, local)

### Deploy

### Git flow
#### Brief Description:

- ```master``` - production version of the code.
- ```dev``` - branch with the most up-to-date codebase.
- ```feat``` - branch for a specific task, created from the develop branch and merged into the develop branch (sub-branches can be created from it if multiple developers are working on the task, named feat-dev).
- ```fix``` - branch for bug fixes, created from the develop branch and merged into the develop branch (sub-branches can be created from it if multiple developers are working on the task, named fix-dev).
- ```hotfix``` - branch for production bug fixes, created from the main branch, and merged into both main and develop (sub-branches can be created from it if multiple developers are working on the task, named hotfix-dev).
- ```release``` - branch with the code for the next release, created from the main branch, into which develop is merged after the planned features are completed, and then merged into main.
- ```[branch type]-dev``` - sub-branch of feat / fix / hotfix for a single developer, created from the specified branch and merged back into it.

#### Branch Naming Conventions:
Branches are named according to the following formulas:
- ```feat``` - [branch type]/[short task name written in snake-case]
- ```feat-dev``` - [branch type]/[short task name written in snake-case]/[developer initials]
- ```hotfix``` - [branch type]/[date]
- ```hotfix-dev``` - [branch type]/[date]/[developer initials]
- ```release``` - [branch type]/[version]
- ```fix``` - [branch type]/[task description]
- ```fix-dev``` - [branch type]/[task description]/[developer initials]

#### Commit Naming Conventions:

Commits are named according to the formula: [task type]: [brief description of what was done]

**Task types**:

- ```feat``` - working on a feature
- ```fix``` - working on a bug
- ```docs``` - documentation changes
- ```style``` - formatting, missing semicolons, etc.; no production code changes
- ```refactor``` - refactoring production code, such as renaming variables
- ```test``` - adding missing tests, refactoring tests; no production code changes
- ```chore``` - updating grunt tasks, etc.; no production code changes