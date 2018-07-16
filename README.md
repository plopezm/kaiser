# Kaiser Job Runner

Kaiser is a extensible job runner. It is responsible for execute jobs provided by the different implemented providers. The language used for developing script is javascript with the functions implemented with plugins. Kaiser can be used as workflow manager, IoT projects or for dynamic job programming.

Currently is under development but any help is very welcome. The language used is javascript (basic javascript) for scripts (in workspace folder you can find some examples).


# Status

- [x] Basic engine
- [x] Basic job foldering 
- [x] Script file dynamic read (scripts can be modified and reexecuted automatically)
- [x] Argument setting
- [x] [Job scheduling] Scheduling jobs using ISO 8601 Duration
- [ ] [Job launcher] Add a way to launch jobs from http events
- [ ] [Job management] GraphQL API for creating, executing, removing tasks remotely.
- [ ] [Web interface] A web interface to manage kaiser 
- [ ] [Plugins] More plugins

# Project structure

- cmd: Contains executables, currently only the engine is there, in the future a terminal connected with graphql will be implemented for managing the daemon
- config: Used for parging "kaiser.config.json" with the program input, in addition the log is configured to be used with files
- core: Contains the implementation of the engine, interpreter (otto) and provider. A Job provider is a thread that find new jobs and provide them to daemon
- interfaces: Different ways to manage, create and query the jobs
- plugins: Implementation of new functionalities for Kaiser interpreter
- utils: Common functions used in the whole project

# Package manager

This project is using dep as package management, to download all required libraries type "dep ensure"

# Configuration

Currently there is only one important configuration called "workspace" where the job folder is defined. 

This is an example of the current configuration file:

```json
{
    "workspace": "./workspace"
}
```

# Workspace

The workspace folder has some rules:

- Folder "disabled" is used to ignore jobs putting them there.
- Every job file has special names, *.job.json


# Examples (Some examples are outdated)

The file graphql.http contains some examples to use with the current API.