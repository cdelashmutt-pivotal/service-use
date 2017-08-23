# service-use 
A cf CLI plugin to display info about service use.
[![wercker status](https://app.wercker.com/status/757c88a99800e1317f3becdc9410ef4a/s/master "wercker status")](https://app.wercker.com/project/byKey/757c88a99800e1317f3becdc9410ef4a)

## Installation
### Install from CF Community Plugins Repo
1. `cf install-plugin -r CF-Community "service-use"`

### Install from project releases section
1. Download the version of plugin for your architecture from the "Releases" tab of the service-use project.
2. If you are on Linux or Mac, you will need to `chmod +x` the file you downloaded.
3. If you already had the plugin installed, uninstall it first with `cf uninstall-plugin ServiceUsePlugin`
4. Now, install the plugin with `cf install-plugin <path-to-downloaded-plugin>` replacing `<path-to-downloaded-plugin>` with the path to the plugin binary you downloaded.

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
