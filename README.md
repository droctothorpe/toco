# toco

`toco` is a CLI that automates TOC (table of contents) generation for GitHub
wikis. 

It generates a TOC based on the files in your wiki and injects them into your
homepage (`Home.md`) and sidebar (`_Sidebar.md`). 

You can see an example of what the resulting TOC looks like
[here](https://github.com/droctothorpe/example-wiki/wiki).

You can set `toco` up as a client-side pre-commit hook or integrate it into your
CI pipeline for even further automation.

## Usage
```
 ______   ______     ______     ______    
/\__  _\ /\  __ \   /\  ___\   /\  __ \   
\/_/\ \/ \ \ \/\ \  \ \ \____  \ \ \/\ \  
   \ \_\  \ \_____\  \ \_____\  \ \_____\ 
    \/_/   \/_____/   \/_____/   \/_____/ 
                                          
A CLI to automate TOC (table of contents) generation for GitHub wikis.

Usage:
  toco [command]

Available Commands:
  gen         Generate a table of contents and inject it into your wiki's homepage and sidebar
  help        Help about any command
  push        Consolidate git add, commit, and push
  version     Print version data

Flags:
      --config string   config file (default is $HOME/.toco.yaml)
  -d, --debug           verbose logging
  -h, --help            help for toco
```
### Generate
```bash
❯ toco gen    
Generating table of contents.
Injecting table of contents into [Home.md _Sidebar.md].
Injection complete. Run 'toco push' to push your changes.
```

You can then `toco push` to push changes to your remote GH wiki with a single
command. This assumes you commit directly to master on your wiki. If not, just
use git directly.

## Filename convention
Your markdown files must adhere to the following naming convention in order for
toco to work:
```
category:title.md
```
For example, the following files:
```
Pipeline:CD.md
Pipeline:CI.md
Standards:Development.md
```
would generate the following TOC:

**Pipeline**  
• [CD](./Pipeline%3ACD)  
• [CI](./Pipeline%3ACI)  
**Standards**  
• [Development](./Standards%3ADevelopment)  

## Installation

### Go get
```bash
go get github.com/droctothorpe/toco
```

### Download binary (OSX only)
```bash
curl https://www.github.com/droctothorpe/toco/toco -o /usr/local/bin/toco && \
chmod 700 /usr/local/bin/toco
```

### Makefile
```bash
git clone https://github.com/droctothorpe/toco.git
make install
```

### Configure
`toco` relies on the presence of the following block in your `Home.md` and
`_Sidebar.md` files to know where to inject the generated TOC:
```
<!--starttoc-->
<!--endtoc-->
```
If it's not present, just add it before running `toco`.

## To do
- Auto-append toc tags if they're not present.
- Make everything more configurable. 
- Automate cross-compiled binary builds. 
- Add pre-commit hook implementation documentation.
- Implement unit tests.
- Add target directory flag.

Contributions are welcome and appreciated.
