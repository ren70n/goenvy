# goenvy
A simple package to deal with ENVs in golang

## Install
simply use go package manager in the project folder
```
go get https://github.com/ren70n/goenvy
```

## Sample code
```go
package main

import "github.com/ren70n/goenvy"

func main(){
  // gets variables from JSON
  envMap,err := goenvy.GetEnvsFromJSON("env.json")
  
  if err!=nil{
    // do what you want
  }
  
  // put envs into system
  goenvy.PushToOSEnvs(envMap)
}
```
