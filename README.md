# service-use 

A cf CLI plugin to display info about service use

## Installation
### Install from pre-built binaries
1. `git clone https://github.com/cdelashmutt-pivotal/service-use`
2. `cd service-use`
3. `cd bin`
4. Now cd to the appropriate architecture for your machine.
5. Install the plugin.
  * Mac/Linux: `cf install-plugin service-use`
  * Windows: `cf install-plugin service-use.exe`

## Usage
Calling `cf service-use` will print out a report of all the Services and their Plans, with the instances that have been created of those services.  You will see all the services and service instances that you have access to.
```
$ cf service-use
Getting service use information as auser@domain.com...

Service searchify:
 Plan small:
 Plan plus:
 Plan pro:

Service blazemeter:
 Plan free-tier:
 Plan basic1kmr:
 Plan pro5kmr:

Service rediscloud:
 Plan 100mb:
 Plan 250mb:
 Plan 500mb:
 Plan 1gb:
 Plan 2-5gb:
 Plan 5gb:
 Plan 10gb:
 Plan 50gb:
 Plan 30mb:
  Org: MyOrg, Space: staging, Instance: redis-service, Managers: [auser2@domain.com]
  Org: MyOrg, Space: production, Instance: df-redis, Managers: [auser2@domain.com]
...
```
