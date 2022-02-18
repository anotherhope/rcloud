# Local To Local
Create un synchronized local folder

## Use case:
- Sync on mounted volume (SMB, AFP, NFS)
- Sync folder in other cloud tools (Synologie Drive, iCloud, ...)

```bash
$ rcloud add /path/to/source /path/to/destination
```

## To Do:

### Feature: 
- [x] Use fsnotify to detect real time change and make a sync
	- [ ] Optimization: Use only Sync at start, and copyto Only on file change
