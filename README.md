## Vision

Vision is a light weight tool written purely in golang for viewing remote resources over HTTP. Vision allows you to view
config files, log files and other such files over HTTP via your browser or on your terminal. It allows you to set ACLs which
via which you can block view on certain resources and alow view on certain resources. It allows you to configure aliases
so that you do not have to type the entire path of the resource on server, view a file from top, or bottom, aplly regex
for filtering contents and specify number of lines to be read form desired files

## Features and use cases
- Viewing resources (log files, config files) on remote servers over http.

- During debugging, when multiple files has to be viewed in different servers, vision can be used to view such files
  on browser or terminal without ssh'ing into all the servers.
  
- Vision allows line limit, reading from head and tail (equivalent to **head** and **tail** in linux), and applying
  regex to filter content (similar to **grep** but limited)
  
- A sysad might not want all resources to be viewed. To address this, vision allows ACLs. You can define simple ACLs like
  allow_all, allow_for, block_for, via which you can allow certain files to be read or blocked from reading or you can block
  a directory altogether.
  
## Getting started

## Building from source (for Debian based systems)

You must have go binaries installed on your system and GOPATH and GOROOT set for bulding from source.

Get the repository using the following go command
```
go get github.com/djmgit/vision/main

```
Next enter into vision directory containing the source (you will find it in your GOPATH i,e, go/src/github.com/djmgit/vision)
and then execute the following make commands

```
make
sudo make install

```
The first command will build vision. Mostly there should not be any issue in this step as vision has no external dependencies.
It uses only the inbuild go packages.
The second make command installs the vision binary on your system.

Once vision is installed on your system, you need to start the service using the following

```
sudo systemctl start vision

```
Vision assumes you use systemd for managing your services. Otherwise you can run vision using :

```
/usr/local/bin/vision
```

By default vision listens on port 8080, however you can change that (see below in configuring vision section)

To verify that vision was installed properly and is working as expected, hit the following URL on browser :

```
[server_ip]:8080/path?=[path_to_a_file_on_your_system]
```
You should be able to view the last 10 lines of the file you have mentioned in the path parameter (provided its a valid paht)

## Uninstalling Vision

First stop the service using :

```
sudo systemctl stop vision
```

Then uninsall and clean your build using the following (execute them in the vision source folder):

```
sudo make uninstall
make clean
```

The above will uninstall vision, remove service and config files and clean the build (remove the dist directory containing
vision binary)

## Configuring Vision

The config file for vision will be present at /etc/vision/config.json
As you have already guessed, the config file is written format.

Given below is a sample config file

```{
	    "port": 8080,
	    "allow_all": true,
	    "block_for": [
		          "path1",
              "path2",
              "dir1",
              "dir2"
	    ],
	    "allow_for": [
		          ""
	    ],
	    "aliases": [
              {
                "alias_name" : "apache2",
                "alias_to" : "/var/log/apache2/access.log",
              },
              {
                "alias_name" : "kafka"
                "alias_to" : "/var/log/kafka/server.log"
              }
      ]
   }

```

Below are the description of each field :

- **port** : it is the port on which vision will listen. Make sure the port you are specifying over here is free

- **allow_all** : If true, all the files in your system can be view via vision except those present in the block_for
                  list. Its value can be true or false. If false, no file in your system can be viewed via vision, except
                  those present in allow_for list.
                  
- **allow_for** : Takes a list of paths. When **allow_all** is **false**, only the file paths present in this list will be
                  visible via vision. The list may also contain a directory path. In this case, all the files in that directory
                  will be visible.
                
- **block_for** : Takes a list of paths. When **allow_all** is **true**, the file paths present in this list will not be visible
                  via vision. The list may also contain a directory path. In that case, all the files in that directory will
                  be unaccessable via vision.
                  
- **aliases** :             
    




