# ginstall

litte bin for run *go install package* on vscode build command (Cmd+Shift+B)

*ginstall /full/path/to/file.ext* runs *go install package*

## usage
in .vscode folder (folder for workspace specific settings) add next tasks.json:
```json
{
    "version": "0.1.0",
    "command": "ginstall",
    "isShellCommand": true,
    "args": ["${file}"],
    "showOutput": "always"
}
```
**where**: *${file}* - vscode substitution for absolute path for current open file

**important**: *ginstall* must be in PATH

now u can use vscode build command (Cmd+Shift+B), it runs *ginstall ${file}*
