# WARNING
This project and the Documentation are currently in progress don't use this for the moment

# Rcloud [![Reference](https://pkg.go.dev/badge/github.com/anotherhope/rcloud.svg)](https://pkg.go.dev/github.com/anotherhope/rcloud) [![License](https://img.shields.io:/github/license/anotherhope/rcloud)](https://github.com/anotherhope/rcloud/blob/main/LICENSE.md)

[![Workflow](https://img.shields.io:/github/workflow/status/anotherhope/rcloud/Go)](https://github.com/anotherhope/rcloud/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/anotherhope/rcloud)](https://goreportcard.com/report/github.com/anotherhope/rcloud)
[![Maintainability](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/maintainability)](https://codeclimate.com/github/anotherhope/rcloud/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/test_coverage)](https://codeclimate.com/github/anotherhope/rcloud/test_coverage)

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=bugs)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=anotherhope_rcloud&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=anotherhope_rcloud)

**Rcloud** is a command-line program to sync files and directories to and from different cloud storage providers in real time and it dependent on **Rclone**.

[//]:![Windows](https://img.shields.io/badge/Windows%20(amd%7Carm)-595959?logo=windows&logoColor=F0F0F0)
![OSX](https://img.shields.io/badge/OSX%20(amd%7Carm)-595959?logo=apple&logoColor=F0F0F0)
![Linux](https://img.shields.io/badge/Linux%20(amd%7Carm)-595959?logo=linux&logoColor=F0F0F0)
![OpenBSD](https://img.shields.io/badge/OpenBSD%20(amd%7Carm)-595959?logo=openbsd&logoColor=F0F0F0)

# Why Rcloud exists ?
I initiated this project because I wanted to be able to synchronize on my NAS (or other cloud providers), my different projects in a secure way in real time with a wide choice of protocols and features.

this project wrap Rclone to add the real-time management layer that it lacks by trying as best as possible to respond to the different problems.

It is above all a personal tool, but if you wish to participate in the improvement, do not hesitate
### Command
- [rcloud](rcloud "rcloud")
	- [config](config "rcloud config")
	- [daemon](daemon "rcloud daemon")
	- [directory](directory "rcloud directory")
		- [add](directory/add "rcloud directory add")
		- [del](directory/del "rcloud directory del")
	- [rts](rts "rcloud rts")
		- [start](rts/start "rcloud rts start")
		- [stop](rts/stop "rcloud rts stop")
	- [selfupdate](selfupdate "rcloud selfupdate")
	- [status](status "rcloud status")

#### [Local To Local](local-to-local "Documentation: Rcloud Local To Local")
#### [Local To Remote](local-to-remote "Documentation: Rcloud Local To Remote")
#### [Remote To Local](remote-to-local "Documentation: Rcloud Remote To Local")
#### [Remote To Remote](remote-to-remote "Documentation: Rcloud Remote To Remote")

### Local To Remote
Create a remote provider with rclone: 
```
$ rclone config
```
Create un synchronized folder
```
$ rcloud add <local_source> <remote_destination:path>
```
### Remote To Local
Create a remote provider with rclone: 
```
$ rclone config
```
Create un synchronized folder
```
$ rcloud add <remote_source:path> <local_destination>
```
### Remote To Remote
Create a remote provider with rclone: 
```
$ rclone config
```
Create un synchronized folder
```
$ rcloud add <remote_source:path> <remote_destination:path>
```