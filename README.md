# pjma

pjma is a project manager for Aveva E3D Design and Administration. It manages the environment to launch E3D Design or Administration in one specified directory.

## Installation

If you're using Go:

```bat
go install github.com/k-awata/pjma@latest
```

Otherwise you can download a binary from [Releases](https://github.com/k-awata/pjma/releases).

## Usage

### Create a project directory

- Command:

  ```bat
  mkdir myproj
  cd myproj
  pjma init
  pjma setup
  ```

- Result:

  ```bash
  myproj
      │  pjma.yaml  # pjma env file
      │
      ├─cafuic    # to store UI customization files
      ├─pmllib    # to store PML2 macros
      ├─pmlui     # to store PML1 macros
      └─projects  # to store E3D project folders
  ```

### Launch an app

```bat
pjma open adm
```

### Make a bat file to launch an app

```bat
pjma mkbat e3d > launch.bat
```

### Add an existing project

```bat
xcopy /e C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\cpl\ projects\cpl\
pjma evars
```

## License

[MIT License](LICENSE)
