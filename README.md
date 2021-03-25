# Fops CLI

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
go build -o fops
./fops
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
- [ ] verify file
- [ ] read file
- [ ] count lines in file

### checksum
- [ ] verify file: binary ok
- [ ] read algorithm flag
- [ ] implement 3 flags

### version
- [ ] get from build

### help
- [ ] implement
- [X] subcommands

### CI
- [ ] Unit tests
- [ ] Build
- [ ] Release

## Dependencies

- [mimetype](github.com/gabriel-vasile/mimetype)  v1.2.0
- [cobra](github.com/spf13/cobra) v0.0.5

## Limitations

- [mimetype](github.com/gabriel-vasile/mimetype) identification is deterministic and dependency free but [non-exhaustive](https://dev.to/sistoi/golang-mime-type-handling-3fnd)

## Contributing

Check out [CONTRIBUTING.md](./CONTRIBUTING.md) to learn how to contribute to this project.

## License

This library is licensed under the Apache 2.0 License.

