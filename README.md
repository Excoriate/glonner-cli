[![CI: Container](https://github.com/Excoriate/glonner-cli/actions/workflows/ci-docker.yml/badge.svg)](https://github.com/Excoriate/glonner-cli/actions/workflows/ci-docker.yml)
[![CI: CLI](https://github.com/Excoriate/glonner-cli/actions/workflows/ci-app.yml/badge.svg)](https://github.com/Excoriate/glonner-cli/actions/workflows/ci-app.yml)
[![CLI-Release](https://github.com/Excoriate/glonner-cli/actions/workflows/generate-release.yml/badge.svg)](https://github.com/Excoriate/glonner-cli/actions/workflows/generate-release.yml)
<h1 align="center">
  Glonner CLI
</h1>
<p align="center"> <img alt="cli-logo" src="./docs/images/logo.png" width="224px"/><br/> </p>

<p align="center">A <b>"glutton" Gopher mutant 🧟‍♂️</b> that helps in cloning massively certain (or all) repositories<b>of a given GitHub organization.</b> <br>Use it with precautions.
</p>
<div align="center">
  <p align="center">
    <a href="https://github.com/Excoriate/glonner-cli/tree/main/docs"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    ·
    <a href="https://github.com/Excoriate/glonner-cli/issues/new?assignees=alextorres-warner&labels=bug&template=bug_report.md&title=">Report Bug</a>
    ·
    <a href="https://github.com/Excoriate/glonner-cli/issues/new?assignees=alextorres-warner&labels=feature&template=feature_request.md&title=">Request Feature</a>
  </p>
</div>

---

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#What-it-does">What it does?</a></li>
        <li><a href="#Commands">Commands</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#for-developers-and-contributors">For Deverlopers & Contributors</a></li>
        <li><a href="#for-users">For Users</a></li>
      </ul>
    </li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#Roadmap">Roadmap</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project
It was created due to some lazziness and the need to clone all the repositories of a given organization. If you are looking for a more in-depth explanation of the project, sorry there isn't. 🤷🏻‍♀️
Some important details:
* It is a CLI tool.
* The term 'organization' and 'owner', in the context of [GitHub](https://github.com), are used interchangeably.
* Depending on the size (amount of repositories) of the organization, the cloning process can take a while.


### What it does?
It browses through all the repositories of a given organization and clones them locally. It also creates a `glonner.json` file with the information of the repositories that were cloned. This file is used to keep track of the repositories that were cloned and to avoid cloning them again. If you want to clone all the repositories again, you can use the `--force` flag.

### Commands

#### Features commands
| Command           | Description                                          |
|-------------------|------------------------------------------------------|
| `glonner clone`   | Clones all the repositories of a given organization. |
| `glonner list`    | List all the repositories of a given organization.   |

#### Required options

These options are required, in order to operate with the above commands.

| Command        | Description                                                                                                                                                                                                                                                                                             |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `organization` | [GitHub organization](https://docs.github.com/en/organizations/collaborating-with-groups-in-organizations/about-organizations) from which the repositories will be fetched from. Mapped to the option `org` (E.g.: `--org="cloudposse"`). Also, it can be set as an environment variable: `GITHUB_ORG`. |
| `token`        | [Personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) used to authenticate against the target GitHub organization..                                                                                                  |

> **_NOTE:_**  For the `clone` command, it's supported the `dry-run` option, set with the flag `--dry-run`. This option will print the list of repositories that would be cloned, without actually cloning them.


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

There are some prerequisites to get started with this project. Ready and follow the instructions below.

### For Developers  & Contributors 🚀
Follow [this guide](docs/guides/setup_development.md) to setup the project locally.

### For Users 🧑‍💻
If you want to clone repositories, it is as simple as running:
```bash
$ glonner clone --org="gruntwork-io" --token="your-token"

```
or
```bash
$ GITHUB_ORG="gruntwork-io" GITHUB_TOKEN="your-token" glonner clone

```
If you only want to list, and take a look to a given set of repositories, within an organization, you can run:
```bash
$ glonner list --org="gruntwork-io" --token="your-token"

```
or
```bash
$ GITHUB_ORG="gruntwork-io" GITHUB_TOKEN="your-token" glonner list

```

> **_NOTE:_**  For advanced options, and more information, please refer to the `help` command, running: `glonner --help`.



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- INSTALLATION -->
## Installation
### Using Homebrew
```bash
brew tap Excoriate/homebrew-tap
brew install glonner
```
Or, if you have [TaskFile](https://taskfile.dev/#/installation) installed, it's the same method, just wrapping it with it:
```bash
task install
```


<!-- USAGE -->
## Usage
![](docs/images/demo.gif)



<!-- ROADMAP -->
## Roadmap

- [ ] ❤️ Add Daeger as a 'portable' approach for CI and CD respectively.
- [ ] ❤️ Add Unit Test coverage.
- [ ] ❤️ Add a built-in configuration, using 'Viper', to keep tack of local repositories cloned.
- [ ] ❤️️ Add a way to selectively clone repositories, based on a given set of criteria.
- [ ] ❤️️ Allow custom output formats. Currently, it's supporting only 'table'. Potential formats: 'json', 'yaml', 'csv', etc.
- [ ] ❤️️ Add a _storage/capacity_ check, for larger organizations and less-space-in-disk sort of machines.

See the [open issues](https://github.com/HBOCodeLabs/hello-world-go/issues) for a full list of proposed features (and known issues). Also, your contributions are more than welcome — just ensure following the [contributing guidelines](docs/guides/contribution_guidelines.md).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact
- 📧 **Email**: [Alex T.](mailto:alex@ideaup.cl)
- 🧳 **Linkedin**: [Alex T.](https://www.linkedin.com/in/alextorresruiz/)

_made/with_ ❤️  🤟
