# Directory Add
This command is used to manage syncronized directory

## Command:

### Local to Local
`$ rcloud directory add <local_path> <local_path>` Add directory to sync

### Local to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add <local_path> <remote_path>` Add directory to sync

### Remote to Local
`$ rcloud config` Create a remote folder

`$ rcloud directory add <remote_path> <local_path>` Add directory to sync

### Remote to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add <remote_path> <remote_path>` Add directory to sync