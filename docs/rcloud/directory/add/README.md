# Directory Add
This command is used to manage syncronized directory

## Command:

### Local to Local
`$ rcloud directory add <local_id> <local_id>` Add directory to sync

### Local to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add <local_id> <remote_id>` Add directory to sync

### Remote to Local
`$ rcloud config` Create a remote folder

`$ rcloud directory add <remote_id> <local_id>` Add directory to sync

### Remote to Remote
`$ rcloud config` Create a remote folder

`$ rcloud directory add <remote_id> <remote_id>` Add directory to sync