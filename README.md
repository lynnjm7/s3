Simple Static Server
====================
A very simple static file server for local development with HTML, JavaScript, 
and CSS. The primary motivation of this is because I needed a simple static 
file server for WebGL development.

# Usage 
Once you've installed this tool with `go get`, using it is pretty simple. The 
most straightforward usage is with the default values for the directory and the 
port. The default directory is your current working directory and the default 
port is 8080.

```bash
%> s3
```

To set the directory and port, use the command line options

```bash
%> s3 -dir <path_to_directory> -port 3030
```

# LICENSE 
This project is licensed under the MIT license. For more information see LICENSE.


