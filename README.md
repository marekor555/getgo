# ABOUT
Getgo was designed as alternative to wget. 
It is easy to use, but because of that there isn't as many functionalities as in wget.
In future I plan to add more things to it.
# INSTALLATION
We will be compiling the project from source.
For newcomers that means you will need `git` and golang compiler(see [go.dev](https://go.dev/dl)). There are a lot of tutorials for that.
## windows, mac and linux non-packaged
1. You firstly must clone the git repo: `git clone https://github.com/marekor555/getgo`
2. Then move to the project: `cd getgo`
3. Get the libraries: `go get .`
4. Install the project: `go install .`
3. Now try the app: `getgo --save --as "duckduckgo.html" "duckduckgo.com"`

If the above command fails, saying that getgo command is not found. 
You need to add go bin to path. And try again.
## arch linux package build
For arch linux and relatives I prepared PKGBUILD.
To do it with the arch package:

1. Clone git repo: `git clone https://github.com/marekor555/getgo`
2. Then mote to the project: `cd getgo`
3. Then run install script: `sh clean.sh`

This script will clean the garbage that is left by makepkg.
