[![Go](https://github.com/anotherhope/rcloud/actions/workflows/go.yml/badge.svg)](https://github.com/anotherhope/rcloud/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/anotherhope/rcloud.svg)](https://pkg.go.dev/github.com/anotherhope/rcloud) [![Go Report Card](https://goreportcard.com/badge/github.com/anotherhope/rcloud)](https://goreportcard.com/report/github.com/anotherhope/rcloud)
[![Maintainability](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/maintainability)](https://codeclimate.com/github/anotherhope/rcloud/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/d5102bdf5504b9ce56ce/test_coverage)](https://codeclimate.com/github/anotherhope/rcloud/test_coverage)

# Rcloud

**Rcloud** is a command-line program to sync files and directories to and from different cloud storage providers in real time and it dependent on **Rclone**.

# How to use

### Local To Local
Create a remote provider with rclone: 
```
$ rclone config
```
Create un synchronized folder
```
$ rcloud add <local_source> <local_destination>
```
Now check your folder is registered check this with
```
$ rcloud status
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
Now check your folder is registered check this with
```
$ rcloud status
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
Now check your folder is registered check this with
```
$ rcloud status
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
Now check your folder is registered check this with
```
$ rcloud status
```