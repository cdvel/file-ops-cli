# Fops CLI 

[![CircleCI](https://img.shields.io/circleci/build/github/cdvel/fops-cli/main?logo=circleci&style=flat-square&token=7c35257d742623bed23bb4d07b4032aa6c38ad25)](https://circleci.com/gh/cdvel/fops-cli)

The File Operations (Fops) CLI, a resilient command line application written in Golang to count line and checksum files.

  - [Installation](#installation)
  - [Usage](#usage)
	  - [Commands](#commands)
 	  - [Flags](#flags)
  - [Roadmap](#Roadmap)
  - [Dependencies](#Dependencies)
  - [Limitations](#Limitations)
  - [Contributing](#Contributing)
  - [License](#license)

## Installation

```
git clone https://github.com/cdvel/fops-cli
cd fops-cli
make
./fops help
```

## Usage

Once installed, the CLI is used via the following standard format:

```sh
fops <command> <file> <flags>
```

### Commands

This CLI supports the following commands:

- `linecount` - Count the number of lines in a file

- `checksum` - Find the checksum of a file using md5, sha1 or sha256

- `version` - Get the current version of Fops

- `help` - Print help


### Flags

| Flag             | Shorthand | Description                                                                                                                                                         | Example              |
|------------------|-----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| `--file`  | `-f`      | Filename to be queried.                                                       | `linecount -f my-sample.txt` |
| `--md5`  | `NA`      | Use MD5 algorith to checksum a file.                                                                                                        | `checksum --md5`             |
| `--sha1`  | `NA`      | Use SHA1 algorith to checksum a file.                                                                                                        | `checksum --sha1`             |
| `--sha256`  | `NA`      | Use SHA256 algorith to checksum a file.                                                                                                        | `checksum --sha26`             |

## Roadmap

### basics
- [X] barebones cli with command placeholders
- [X] read file argument
- [ ] argument detection (file, algorithm)

### file verification
- [X] non-existent
- [X] is a dir
- [X] is binary

### linecount
- [X] verify file
- [X] read file
- [X] count lines in file

### checksum
- [X] verify file: binary ok
- [X] read algorithm flag
- [X] implement 3 flags

### version
- [X] get from build

### help
- [X] implement
- [X] subcommands

### CI
- [X] Unit tests
- [X] Integrate with CircleCI
- [X] Build
- [ ] Release

## Dependencies

- [mimetype](https://github.com/gabriel-vasile/mimetype)  v1.2.0
- [cobra](https://github.com/spf13/cobra) v0.0.5

## Limitations

- [mimetype](https://github.com/gabriel-vasile/mimetype) identification is deterministic and dependency free but [non-exhaustive](https://dev.to/sistoi/golang-mime-type-handling-3fnd)

## Contributing

Check out [CONTRIBUTING.md](./CONTRIBUTING.md) to learn how to contribute to this project.

## License

This library is licensed under the Apache 2.0 License.

