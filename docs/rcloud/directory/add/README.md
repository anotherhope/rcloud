# Directory Add
This command is used to manage syncronized directory

## Command:

### Local to Local
`$ rcloud directory add /path/to/local /path/to/local` Add to sync
### Local to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add /path/to/local remote:/path/to/remote` Add to sync
### Remote to Local
`$ rcloud config` Create a remote folder

`$ rcloud directory add remote:/path/to/remote /path/to/local` Add to sync

### Remote to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add remote:/path/to/remote remote:/path/to/remote` Add to sync