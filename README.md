# goenvy
A simple, lightweight, package to deal with ENVs in golang for those who prefere clear JSON over YAML, ENV etc

goenv takes a JSON file and changes it into environment variables. Those last for the time of your project running

Ideal for switching between dev/test/prod variables

## Installation
simply use go package manager in the project folder
```bash
go get github.com/ren70n/goenvy
```

## Usage
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
## License
[MIT](https://choosealicense.com/licenses/mit/)
