# Uxmal

uxmal is a cli tool to create quick and dirty jobs in kubernetes without too much hassle. 
This is a debugging and troubleshotting tool, not for use in production environments.

Instead of creating helm charts, complex yamls files, uxmal can easily create jobs for testing permissions, programs or applications.

## Usage
```
USAGE:
   uxmal command [command options]

COMMANDS:
   cron:    uxmal cron
      create:  uxmal cron create --name myCron --img ubuntu --cmd "ls -ltr" --sch "* * * * *"
      delete:  uxmal cron delete --name myCron
   job:     uxmal job
      create:  uxmal job create --name myJob --img ubuntu --cmd "ls -ltr"
      status:  uxmal job status --name myJob
      logs:    uxmal job logs --name myJob
      run:     uxmal job run --name myJob --img ubuntu --cmd "ls -ltr"   
      delete:  uxmal job delete --name myJob
   help, h    Shows a list of commands or help for one command
```
## Installation

``` 
go build .
mv uxmal /usr/bin/local
```
