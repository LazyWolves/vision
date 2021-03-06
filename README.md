[![Build Status](https://travis-ci.org/LazyWolves/vision.svg?branch=dev)](https://travis-ci.org/LazyWolves/vision)

## Vision

Vision is a light weight tool written purely in golang for viewing and fetching information on your system's state remotely. Vision allows you to view
config files, log files and other such files over HTTP via your browser or on your terminal. It allows you to set ACLs via 
which you can block view on certain resources and alow view on certain resources. It allows you to configure aliases
so that you do not have to type the entire path of the resource on server, view a file from top, or bottom, apply regex
for filtering contents and specify number of lines to be read form desired files.

Apart from viewing file resources it also allows you to view information about your remote host, processes running and their state and information on them, system metrcis like CPU and Memory, status of your systemd services running and option to
start and stop them.

## Features and use cases

### Viewing remote resource files

- Viewing resources (log files, config files) on remote servers over http.

- During debugging, when multiple files has to be viewed in different servers, vision can be used to view such files
  on browser or terminal without ssh'ing into all the servers.
  
- Vision allows line limit, reading from head and tail (equivalent to **head** and **tail** in linux), and applying
  regex to filter content (similar to **grep** but limited)
  
- A sysad might not want all resources to be viewed. To address this, vision allows ACLs. You can define simple ACLs like
  allow_all, allow_for, block_for, via which you can allow certain files to be read or blocked from reading or you can block
  a directory altogether.
  
### Fetching remote Host information

- Vision allows you to view basic informations about your remote hosts over HTTP
- Host information includes ```hostname```, ```uptime```, ```bootTime```, ```Porcs```, ```OS```, ```Platform```, ```Arch```,
  ```KernelVersion```, ```Virtualisation Type```, ```Virtualisation Role```, etc.

### Fetching System Metrics from your remote system

- Vision allows you to view systems metrics corresponding to your rempte host
- As of now it allows you to fetch ```CPU utilisation``` and ```memory utilisation```.

### Fetching process information

- Vision allows you to view all the processes running presently on your remote system.
- You can query details of a process via its PID.
- For a queried process it shows ```Name```, ```command line arguments```, ```executable path```,
  ```current working directory```, ```Status```, ```Number of Threads```, ```GIDS and UIDS```, etc.

### Fetching systemd services information

- Vision allows you to view the ```status``` of the various systemd services running in your remote system
- It allows you to list all the systemd services running in your system, and also allows you to filter via name.
- You can also ```start``` or ```stop``` a service via vision remotely over HTTP.
  
## Getting started

## Installing from debian package

### NOTE : The debian package does not yet support querying host info, process info, systemd info and system metric info. These features will be included in the package in the next release.

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
The first command will build vision.

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
As you have already guessed, the config file is written in json format.

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

The following section describes how to use Vision endpoints in order to use its various features.

### Viewing remote resource files

The Base path is [server_ip]:[port]/

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

### Viewing remote host information

The Base path is [server_ip]:[port]/hostInfo

Example:

```

API call : http://[server-ip]:[port]:8080/hostInfo

Response:

{
  "HostInfo": {
    "hostname": "vision-test",
    "uptime": 6747,
    "bootTime": 1580965085,
    "procs": 322,
    "os": "linux",
    "platform": "ubuntu",
    "platformFamily": "debian",
    "platformVersion": "18.04",
    "kernelVersion": "4.15.0-76-generic",
    "kernelArch": "x86_64",
    "virtualizationSystem": "kvm",
    "virtualizationRole": "host",
    "hostid": "4bdc21cc-2895-11b2-a85c-9a4598d18182"
  },
  "Timestamp": 1580971832,
  "TimestampUTC": "2020-02-06 06:50:32.614724886 +0000 UTC"
}

```

### Fetching System Metrics from your remote system

The Base path is [server_ip]:[port]/systemMetrics

Example:

```
API call : http://[serve_ip]:[port]/systemMetrics

Response:

{
  "Metrics": {
    "CPU": {
      "LoadAvg": {
        "Load1": 0.77,
        "Load5": 0.66,
        "Load15": 0.57
      }
    },
    "Memory": {
      "VirtualMemory": {
        "MemTotal": 16329969664,
        "MemFree": 9131687936,
        "UsedPercent": 23.65463206288538
      }
    }
  },
  "Timestamp": 1580972115,
  "TimestampUTC": "2020-02-06 06:55:15.481317283 +0000 UTC"
}

```

### Fetching process information

The Base path is [server_ip]:[port]/procs

All the options are passed as URL params. Only GET method is required.

|      Param      |      Type        |      Description       |
|:----------------|:-----------------|:-----------------------|
|   pid	  |      int      |   Pid of the desired process to be queried|

Example:

```
API call : http://[server_ip]:[port]/procs

Response:

{
  "ProcList": [
    {
      "Pid": 0,
      "Name": "",
      "CmdLine": ""
    },
    {
      "Pid": 0,
      "Name": "",
      "CmdLine": ""
    },
    {
      "Pid": 1,
      "Name": "systemd",
      "CmdLine": "/sbin/init splash"
    },
    {
      "Pid": 2,
      "Name": "kthreadd",
      "CmdLine": ""
    }
  ],
  "Timestamp": 1580972613,
  "TimestampUTC": "2020-02-06 07:03:33.870976948 +0000 UTC"
}

API call : http://[server_ip]:[port]/procs?pid=1234

Response:

{
  "ProcDesc": {
    "Pid": 1234,
    "Ppid": 1,
    "Name": "mongod",
    "CmdLine": "/usr/bin/mongod --config /etc/mongod.conf",
    "ExePath": "/usr/bin/mongod",
    "Cwd": "/",
    "Status": "S",
    "Uids": [
      122,
      122,
      122,
      122
    ],
    "Gids": [
      127,
      127,
      127,
      127
    ],
    "Nice": 20,
    "NumThreads": 26
  },
  "Timestamp": 1580972908,
  "TimestampUTC": "2020-02-06 07:08:28.947810077 +0000 UTC"
}

```

### Fetching systemd services information

The Base path is [server_ip]:[port]/systemd

|      Param      |      Type        |      Description       |
|:----------------|:-----------------|:-----------------------|
|   operation	  |      string      |   Expected values are **start**, **stop** and **list**. list will list the systemd services running, start will start a given systemd service, and stop will stop the service|
|   filterBy      |      string      |   Excepts a regex, will list systemd services whose name matches the given regex, should be used with operation=list |
|   serviceName   |      string      |   Excepts a service name, used to specify service to be started or stopped, should be used with operation=start|stop |

Example :

```
The following API call can be used to list all systemd services running:

http://[server_ip]:[port]/systemd?operation=lisy&filterBy=*.service

Response: 

{
  "Services": [
    {
      "ServiceName": "festival.service",
      "ServiceState": "inactive"
    },
    {
      "ServiceName": "user@1000.service",
      "ServiceState": "active"
    },
    {
      "ServiceName": "openvpn.service",
      "ServiceState": "active"
    },
    {
      "ServiceName": "ModemManager.service",
      "ServiceState": "active"
    },
    {
      "ServiceName": "sssd.service",
      "ServiceState": "inactive"
    },
    {
      "ServiceName": "unattended-upgrades.service",
      "ServiceState": "inactive"
    },
    {
      "ServiceName": "systemd-fsck-root.service",
      "ServiceState": "inactive"
    },
    {
      "ServiceName": "kerneloops.service",
      "ServiceState": "active"
    },
    {
      "ServiceName": "console-setup.service",
      "ServiceState": "active"
    }
  ],
  "NumServices": 131,
  "Timestamp": 1580974430,
  "TimestampUTC": "2020-02-06 07:33:50.024080951 +0000 UTC"
}

The following API call can be used to stop haproxy, (if haproxy is running in the system, just an example)

API call : http://[server_ip]:[port]/systemd?operation=stopt&serviceName=haproxy

Response:

{
  "Status": "OK",
  "Timestamp": 1580974652,
  "TimestampUTC": "2020-02-06 07:37:32.048756127 +0000 UTC"
}

Similarly the following can be used to start it back

API call : http://[server_ip]:[port]/systemd?operation=start&serviceName=haproxy

Response:

{
  "Status": "OK",
  "Timestamp": 1580974704,
  "TimestampUTC": "2020-02-06 07:38:24.708853208 +0000 UTC"
}

```

## Contributing to Vision

All opensource enthusiasts and Golang lovers are most welcome to add more features and fix bugs in vision.

Please have a look at <a href="https://github.com/djmgit/vision/blob/master/CONTRIBUTING.md">CONTRIBUTING.md</a>






