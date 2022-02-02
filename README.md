# pjma

Project manager for Aveva E3D Design and Administration

## Installation

If you're using Go:

```bat
go install github.com/k-awata/pjma@latest
```

Otherwise you can download a binary from [Releases](https://github.com/k-awata/pjma/releases).

## Usage

- Initialize a project directory

  ```bat
  mkdir myproj
  cd myproj
  pjma init
  pjma exec setup
  ```

- Launch Administration 1.9

  ```bat
  pjma run adm19
  ```

- Make a bat file to launch E3D 3.1

  ```bat
  pjma mkbat launch.bat e3d31
  ```

- Add an existing project

  ```bat
  xcopy /e C:\Users\Public\Documents\AVEVA\Projects\E3D3.1\cpl\ projects\cpl\
  pjma mkevars
  ```
