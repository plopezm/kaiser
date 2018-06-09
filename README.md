# Kaiser

Kaiser is a job executer, currently is under development but any help is very welcome. The language used is javascript for scripts (in workspace folder you can find some examples). The Javascript interpreter is simple, you can find more information [here](https://github.com/robertkrimen/otto)

The idea is to create a job executer that can be modified or updated using a graphQL api. This can be used for IoT project or for dynamic job programming.

# Status

- [x] Basic engine
- [x] Basic job foldering 
- [x] Script file dynamic read (scripts can be modified and reexecuted automatically)
- [x] Argument setting
- [ ] [Job management] GraphQL API for creating, executing, removing tasks remotely.
- [x] [Job scheduling] Scheduling jobs using ISO 8601 Duration
- [ ] [Command interface] A commandline interface to manage kaiser 
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
