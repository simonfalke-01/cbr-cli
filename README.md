# cbr-cli
Submit your solutions to codebreaker.xyz without leaving the terminal.
This is a cli utility written in Go that allows you to directly submit your solutions to a problem to codebreaker.xyz,
all in the comfort of your commandline.

## Usage
`cbr <problem id> <path-to-solution>` </br>
Example: </br>
`cbr helloworld main.cpp`

## Installation
### Linux
`wget https://github.com/simonfalke-01/cbr-cli/releases/latest/download/cbr-linux && chmod +x ./cbr-linux && mv ./cbr-linux /usr/local/bin/cbr`
### macOS
`wget https://github.com/simonfalke-01/cbr-cli/releases/latest/download/cbr-darwin && chmod +x ./cbr-darwin && mv ./cbr-darwin /usr/local/bin/cbr`
### Windows
Check the GitHub Actions tab for Windows builds.

## Disclaimer
This is not an official utility. This is a personal project and is not affiliated with codebreaker.xyz in any way.
Do note, that the utility uses your browser's session cookie for codebreaker.xyz to authenticate you. 
The session cookie will **never** leave your machine. They are only sent to codebreaker.xyz to authenticate you.
