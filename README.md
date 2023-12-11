# Color-Converter-CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/wendy-YW/clr-cli)](https://goreportcard.com/report/github.com/wendy-YW/clr-cli)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/wendy-YW/clr-cli)


Color-Converter-CLI is a command-line interface (CLI) tool designed for color conversion.  
This application is a powerful tool that allows users to convert colors between different color models such as RGB, HEX.  
For example, if you have a color in RGB format like rgb(255, 0, 0) and you want to convert it to HEX format, you can simply use our tool like this:
```
clr-cli toHex -r 255,0,0 or  clr-cli 2x -rgb 255,0,0
```
And it will output the color in HEX format:` #FF0000`  
It also allows you to save the colors you have converted to a file and list them later.
```
clr-cli log list or clr-cli log l 
```
to list all the colors you have converted.  

You can also generate a token for your saved colors and they will be sent to our server, so you can see your saved colors from color converter website.
```
clr-cli log token or clr-cli log t
```
Color-Converter-CLI is built using Go and Cobra, a CLI library for Go that empowers applications.

# Installation

for MacOS/Linux
```
sudo /bin/sh -c "$(curl -fsSL https://raw.githubusercontent.com/wendy-YW/clr-cli/release/godownloader.sh)" -- -b /usr/local/bin
```

for Windows
```
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/wendy-YW/clr-cli/release/godownloader.sh" -OutFile "godownloader.sh"
bash ./godownloader.sh
```
> if you encounter any problems regarding installing it on the windows, please open an issue

# Usage

you can use --help to check all the available commands
```
clr-cli --help
```
you can run the command below to check the color you have converted
```
clr-cli log list or clr-cli log l
```
to clear the log, run
```
clr-cli log clear or clr-cli log c
```
to delete certain color in the log, run
```
clr-cli log delete [index] or clr-cli log d [index]
```

# Reminder

Before using `clr-cli log token`, you need to add your env variable first  

## Set Env

Type below in your terminal, and then run clr-cli log token
```
export POST_URL=the URL we offered on the website
clr-cli log token
```
  
# Contribution & Feedback

Contributing and feedbacks are more than welcome. Please feel free to submit a Pull Request.
  
# External Libraies
| Library Name | Link | License | 
|---|---|---|
| Cobra | https://github.com/spf13/cobra| Apache License 2.0 |
| go-pretty | https://github.com/jedib0t/go-pretty| MIT License |


