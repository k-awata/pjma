# pjma

Project manager for Aveva E3D Design and Administration

## Installation

```bat
go install github.com/k-awata/pjma
```

## Usage

- Initialize a project directory

  ```bat
  mkdir myproj
  cd myproj
  pjma init
  ```

- Run Hello world script

  ```bat
  pjma exec hello
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
