AppEnvs
======

A simple Cloud Foundry CLI plugin written in Go that exports application environment variables locally.

Usage
--
```
$ cf appenv APP_NAME
export VCAP_SERVICES='...'
export VCAP_APPLICATION='...'
```

Features
--
- VCAP_SERVICE support 
- VCAP_APPLCATION support 
- USER DEFINED Variable support (TODO)

Installation
-
Install the plugin:
```
git clone https://github.com/Samze/appenvs.git
cd appenvs

go build
cf plugin-install appenv 
```
