[Website](https://anotherhope.github.io/rcloud/) |
[Documentation](https://pkg.go.dev/github.com/anotherhope/rcloud) 

### Documentation
[![Go Reference](https://pkg.go.dev/badge/github.com/anotherhope/rcloud.svg)](https://pkg.go.dev/github.com/anotherhope/rcloud)

### Build
[![Go](https://github.com/anotherhope/rcloud/actions/workflows/go.yml/badge.svg)](https://github.com/anotherhope/rcloud/actions/workflows/go.yml)

### Test & Quality
[![Go Reference](https://pkg.go.dev/badge/github.com/anotherhope/rcloud.svg)](https://pkg.go.dev/github.com/anotherhope/rcloud)
[![Go Report Card](https://goreportcard.com/badge/github.com/anotherhope/rcloud)](https://goreportcard.com/report/github.com/anotherhope/rcloud)
[![Maintainability](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/maintainability)](https://codeclimate.com/github/anotherhope/rcloud/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/test_coverage)](https://codeclimate.com/github/anotherhope/rcloud/test_coverage)
[![License](https://img.shields.io:/github/license/anotherhope/rcloud)](https://github.com/anotherhope/rcloud/blob/main/LICENSE.md)


# Rcloud

**Rcloud** is a command-line program to sync files and directories to and from different cloud storage providers in real time and it dependent on **Rclone**.

# How to use

### Local To Local
Create un synchronized folder
```
$ rcloud add <local_source> <local_destination>
```
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