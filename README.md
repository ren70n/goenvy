# goenvy
A simple, lightweight, package to deal with ENVs in golang for those who prefere clear JSON over YAML, ENV etc

Ideal for switching between dev/test/prod variables

## How it works

1. goenvy takes a JSON file and transforms its content into environment variables. 
```json
{
	"KEY":"value"
}
```
2. Those last for the time of your project running
3. All the keys are automatically transformed to UPPERCASE so you don't have to bother
4. In case of arrays key is given the index of array element:
```json
{
	"KEY":["arr1","arr2"]
}
```
and result is:
```bash
KEY_0=arr1
KEY_1=arr2
```

## Available functions

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
