# pjma

Project manager for Aveva E3D Design and Administration

## Installation

If you're using Go:

```bat
go install github.com/k-awata/pjma@latest
```

Otherwise, you can download a binary from [Releases](https://github.com/k-awata/pjma/releases).

## Usage

### Create a project directory

Command:

```bat
mkdir myproj
cd myproj
pjma init
pjma exec setup
```

Result:

```bash
myproj
    │  pjmaconf.yaml  # pjma config file
    │
    ├─cafuic    # to store UI customization files
    ├─pmllib    # to store PML2 macros
    ├─pmlui     # to store PML1 macros
    └─projects  # to store E3D project folders
```

### Launch an app

```bat
pjma run adm19
```

### Make a bat file to launch an app

```bat
pjma mkbat e3d31 launch.bat
```

### Add an existing project

```bat
xcopy /e C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\cpl\ projects\cpl\
pjma mkevars
```

## License

[MIT License](LICENSE)
