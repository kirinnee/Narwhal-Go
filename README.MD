# Narwhal Go 

A Go CLI that allows users to load and save docker volumes as tarball

# Installation
Narwhal is available as to install in various platforms

### Ubuntu, apt-get
```bash
sudo add-apt-repository ppa:kirinnee/main
sudo apt-get update
sudo apt-get install narwhal
```

### Go installation
```bash
go install gitlab.com/kiringo/narwhal 
```

### Other installations
[.NET Core](https://gitlab.com/nuget-packages/narwhal-.net)


# Usage
## Saving
Saving docker volumes to tarball
```bash
narwhal save <volume> <tarball name> <path to save>
```
| var  | description  |
|---|---|
| volume | the name of a docker named volume|
| tarball name | the file name of the tarball. if `group` is given, the tarball generated will be `group.tar.gz`. Optional. Default:`data` | 
| path to save | the path to save the tarball on host to. relative to terminal directory. Optional. Default: `./`

## Loading
Loading a tarball to a docker volume.
Do note that if the volume have files, the will overwrite files with same path and name.

```bash
narwhal load <path to tarball> <volume>
```
| var  | description  |
|---|---|
| path to tarball | The path to the tarball to load on the host machine, relative to the terminal's directory. |
| volume | the name of a docker named volume|



## Contributing
Please read [CONTRIBUTING.md](CONTRIBUTING.MD) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning 
We use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository.etting

## Authors
* [kirinnee](mailto:kirinnee97@gmail.com) 

## License
This project is licensed under MIT - see the [LICENSE.md](LICENSE.MD) file for details