# Installation

for MacOS/Linux
```
sudo /bin/sh -c "$(curl -fsSL https://raw.githubusercontent.com/wendy-YW/color-converter-cli/test/godownloader.sh)" -- -b /usr/local/bin
```

for Windows
```
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/wendy-YW/color-converter-cli/test/godownloader.sh" -OutFile "godownloader.sh"
bash ./godownloader.sh
```
> if you encounter any problems regarding installing it on the windows, please open an issue

# Usage

you can use --help to check all the available commands
```
clr-cli --help
```
we added 4 colors in the log beforehand, you can run the command below to check
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

