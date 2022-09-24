<div align="center">
    <h1>Shawty</h1>
    <strong>a basic url shortener</strong>
</div>

---

## Usage
### Server
Shawty needs a running postgres instance to work. You could either download and run the docker image from [ghcr](https://github.com/cloudybyte/shawty/pkgs/container/shawty) or use the [provided docker-compose config](https://github.com/cloudybyte/shawty/blob/main/docker-compose.yml)


⚠️ **__Note:__** Do **NOT** run the server without a reverse proxy in front of it. <br> Additionally: I cannot give any guarantees regarding the security of this project. __Run at your own risk.__


### CLI
Download the precompiled binary off of the releases section (or build it yourself), login and have fun.

### Utility Script
This script should only be needed for development. 
Its only dependency is [gum](https://github.com/charmbracelet/gum).<br>
The script will help you sticking to conventions, building images and stuff (look it up if you want to know what it exactly does. Thats what you should do before running sth off the internet anyways).
Running it is super simple:
Just invoke `./util.sh` without any arguments. Navigation should be pretty self explanatory.