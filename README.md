## Vision

Vision is a light weight tool written purely in golang for viewing remote resources over HTTP. Vision allows you to view
config files, log files and other such files over HTTP via your browser or on your terminal. It allows you to set ACLs via 
which you can block view on certain resources and alow view on certain resources. It allows you to configure aliases
so that you do not have to type the entire path of the resource on server, view a file from top, or bottom, apply regex
for filtering contents and specify number of lines to be read form desired files.

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

## Installing from debian package

Get the debian package using the following:

```
wget -O vision_0.1-1_amd64.deb https://github.com/djmgit/vision_debian/blob/master/vision_0.1-1_amd64.deb?raw=true

```

Once the package is downloaded, installing the package using the following

```
sudo dpkg -i vision_0.1-1_amd64.deb

```

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
Vision uses only the inbuild go packages.
The second make command installs the vision binary on your system.

### Building DOC locally

To build the HTML version of the doc and serve locally, execute the following in the repo folder

```

make doc

```
Then hit ```http://127.0.0.1:6060/pkg/github.com/djmgit/vision/``` on the browser to view the documentaion of the source code.

## Post Installation : Testing it out

Once vision is installed on your system, you need to start the service using the following

```
sudo systemctl start vision

```
Vision assumes you use systemd for managing your services. Otherwise you can run vision using :

```
/usr/bin/vision
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
                  
- **aliases** :  Takes a list of objects. An object contains two keys - **alias_name** and **alias_to**. alias_name is the
		 name of the alias (basically a short name for a long path). Once a alias has been added, the corresponding
		 resource can be queried using the alias without providing the full path all the time.
		 
It is to be noted that allow_for is used when allow_all is set to **false** and block_for is used when allow_all is set to
**true**. Providing value for both allow_for and block_for will not have any affect.

## Usage and API endpoints

The Base path is simply [server_ip]:[port]/

All the options are simply passed as URL params. Only GET method is required.

Following are the options that can be used with vision :
    
|      Param      |      Type        |      Description       |
|:----------------|:-----------------|:-----------------------|
|   path	  |      String      |   Absolute path of the resource file on the remote system|
|   readFrom 	  |      String      |   Specifies from where to read the file, can be wither of head ot tail|
|   limit         |      Integer     |   Specifies the number of lines to be read|
|   filterBy      |  String (regex)  |   A regex to filter out desired lines from the given file. Only thoe lines containing patterns matched by the given regex will be returned.|
|   ignore        |  String (regex)  |   A regex to exclude lines containing patterns matching the regex |
|   alias         |  String          |   An alias name. Must be configured beforehand |

Some examples :

```
The following will return the last 10 lines of apache access log. If limit and readFrom is not mentioned, by default 
readFrom is tail and limit is 10.

http://[server-ip]:[port]?path=/var/log/apache2/access.log

The following will return the first 100 lines of apache access log

http://[server-ip]:[port]?path=/var/log/apache2/access.log&readFrom=head&limit=100

The following will return only those lines which contain the word INFO from kafka log

http://[server-ip]:[port]?path=/var/log/kafka/server.log&readFrom=head&limit=100&filterBy=INFO

The following will return only those lines which does not contain the word WARNING from kafka log

http://[server-ip]:[port]?path=/var/log/kafka/server.log&readFrom=head&limit=100&ignore=WARNING

The following will read the first 20 lines of the file aliases by 'nginx'

http://[server-ip]:[port]?alias=nginx&readFrom=head&limit=100&ignore=WARNING

```

## Contributing to Vision

All opensource enthusiasts and Golang lovers are most welcome to add more features and fix bugs in vision.

Please have a look at <a href="https://github.com/djmgit/vision/blob/master/CONTRIBUTING.md">CONTRIBUTING.md</a>






