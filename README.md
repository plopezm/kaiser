# Kaiser

Kaiser is a job executer, currently is under development but any help is very welcome. The language used is javascript for scripts (in workspace folder you can find some examples).

The idea is to create a job executer that can be modified or updated using a rest api. This can be used for IoT project or for dynamic programming.

# To be implemented

- [] [Job management] Rest API for creating, executing, removing tasks remotely.
- [x] [Job scheduling] Scheduling jobs using ISO 8601 Duration
- [] [Command interface] A commandline interface to manage kaiser 
- [] [Web interface] A web interface to manage kaiser
- [] [Plugins] More plugins

# What is implemented

- [x] Basic engine
- [x] Basic job foldering 
- [x] Script file dynamic read (scripts can be modified and reexecuted automatically)
- [x] Argument setting