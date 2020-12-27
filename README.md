<h1 align="center">genday</h1>
<p align="center"><i>Made with :heart: by <a href="https://github.com/GreatGodApollo">@GreatGodApollo</a></i></p>

`curday.dat` Generator for Prevue

## Installing

### Standard Download
Just head on over to the [releases](https://github.com/GreatGodApollo/genday/releases) page and download the latest release
for your platform. Extract it using something like [7-Zip](https://www.7-zip.org) for Windows or `tar` on other 
platforms (`tar -zxvf genday*.tar.gz`).

That's it! Although you'll probably want to also add the binary to your path for ease of use.

### Scoop
Do you happen to have [scoop](https://github.com/lukesampson/scoop) installed? Well lucky for you, I happen to have a [scoop](https://github.com/lukesampson/scoop) bucket. Said bucket's name is [Trough](https://github.com/GreatGodApollo/trough).
```shell
$ scoop bucket add trough https://github.com/GreatGodApollo/trough.git
$ scoop install genday
```

### Go Get
Do you have go installed? You can run just one simple command to install Genday!
```shell
$ go get -u github.com/GreatGodApollo/genday
```

## Usage

*For detailed usage please visit my [documentation](https://docs.brettb.xyz/prevue/genday/index.html)*

```shell
 $ genday --help
    Usage: genday [--version] [--verbose] COMMAND [arg...]
    
    Generate a curday.dat file
    
    Options:
      -v, --version   Show the version and exit
      -V, --verbose   Verbose debug mode
    
    Commands:
      json            generate a file from JSON
    
    Run 'genday COMMAND --help' for more information on a command.
```

## Bugs & Feature Requests
As you may have noticed, this repo has issues disabled, that's because everything is over on my bug
portal @ [pm.brettb.xyz](https://pm.brettb.xyz). You can browse around without an account, but you
do need to make one in order to file anything.

## Built With

- [fatih/color](https://github.com/ttacon/chalk)
- [jawher/mow.cli](https://github.com/jawher/mow.cli)

## Licensing

This project is licensed under the [MIT License](https://choosealicense.com/licenses/mit/)

## Authors

- [Brett Bender](https://github.com/GreatGodApollo)
