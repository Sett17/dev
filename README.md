<div align="center">

  <h1>Devfile</h1>

  <p>
    Tired of the weird quirks of make? Annoyed of making typos in long chained commands, or getting to them via reverse-i-search?
    Well, here is a solution that comes as just an executable for each platform and with an extensive help command.
  </p>

  <p>
    <i>This is a rewrite in golang from the original version written in kotlin. This old version can be accessed <a href="https://github.com/Sett17/DevfileKotlin">here</a>. Note that many of the issues still in the kotlin native version, are fixed in this version</i>
  </p>


<!-- Badges -->
<p>
  <a href="https://github.com/Sett17/Devfile/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/Sett17/Devfile" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/Sett17/Devfile" alt="last update" />
  </a>
  <a href="https://github.com/Sett17/Devfile/network/members">
    <img src="https://img.shields.io/github/forks/Sett17/Devfile" alt="forks" />
  </a>
  <a href="https://github.com/Sett17/Devfile/stargazers">
    <img src="https://img.shields.io/github/stars/Sett17/Devfile" alt="stars" />
  </a>
  <a href="https://github.com/Sett17/Devfile/issues/">
    <img src="https://img.shields.io/github/issues/Sett17/Devfile" alt="open issues" />
  </a>
  <a href="https://github.com/Sett17/Devfile/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/Sett17/Devfile.svg" alt="license" />
  </a>
</p>
</div>

<br />

<!-- Table of Contents -->

# Table of Contents

- [Table of Contents](#table-of-contents)
  - [About the Project](#about-the-project)
  - [Getting Started](#getting-started)
    - [Installation](#installation)
    - [Update](#update)
    - [Build](#build)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [License](#license)
  - [Acknowledgements](#acknowledgements)

<!-- About the Project -->

## About the Project

<div align="center"> 
  <img src="https://i.imgur.com/QotDt8R.png" alt="screenshot" />
</div>

<!-- Getting Started -->

## Getting Started

<!-- Installation -->

### Installation

To install Devfile from the command line you need to have [go installed](https://go.dev/doc/install).

```bash
go install github.com/Sett17/dev@latest
```

or download the latest release executable for your system [here](https://github.com/Sett17/dev/releases) and put it
somewhere in your PATH.

<!-- Update -->

### Update

To update Devfile you first need to remove thr currently installed version.

```bash
rm $GOPATH/bin/dev
```

Then you can follow the steps from [Installation](#installation).

<!--Build-->

### Build

Clone the project

```bash
  git clone https://github.com/Sett17/dev.git
```

Go to the project directory

```bash
  cd dev
```

And get the dependencies and build

```bash
  go get -x
  go build
```

<!-- Usage -->

## Usage

To get started, either create a file called `devfile`

And populate it with operations, according to the [REFERENCE.md](REFERENCE.md) section.

<!-- Contributing -->
## Contributing

<a href="https://github.com/Sett17/Devfile/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Sett17/Devfile" />
</a>


Contributions are always welcome!

There are already issues for enhancements, and everyone is welcome to create new ones, or work on some existing issues.


<!-- License -->

## License

Distributed under the GNU General Public License v3.0. See LICENSE for more information.

<!-- Acknowledgments -->

## Acknowledgements

Use this section to mention useful resources and libraries that you have used in your projects.

- [cfmt](https://github.com/i582/cfmt)
- [simpletable](https://github.com/alexeyco/simpletable)