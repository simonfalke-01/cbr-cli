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
`wget https://github.com/simonfalke-01/cbr-cli/releases/latest/download/cbr-darwin && chmod +x ./cbr-linux && mv ./cbr-darwin /usr/local/bin/cbr`
### Windows
If you are a windows user I don't get why you would ever prefer the cmd prompt or PowerShell over the actual codebreaker webpage.
Like seriously. Imagine having such a god awful "terminal". </br>
But if you really want to use it? **Clone the repository and build it from source.** </br>
**I will not provide Windows builds because I don't have access to a Windows machine.**

## Disclaimer
This is not an official utility. This is a personal project and is not affiliated with codebreaker.xyz in any way.
Do note, that the utility uses your browser's session cookie for codebreaker.xyz to authenticate you. 
The session cookie will **never** leave your machine. They are only sent to codebreaker.xyz to authenticate you.