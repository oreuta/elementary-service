# Glide

What is glide...

Glide is Vendor Package Management for Golang
Glide scans the source code of your application or library to determine the needed dependencies. To determine the versions and locations (such as aliases for forks) Glide reads a glide.yaml file with the rules. With this information Glide retrieves needed dependencies.

When a dependent package is encountered its imports are scanned to determine dependencies of dependencies (transitive dependencies). If the dependent project contains a glide.yaml file that information is used to help determine the dependency rules when fetching from a location or version to use. Configuration from Godep, GB, GOM, and GPM is also imported.    
#Install

Linux:

    curl https://glide.sh/get | sh
    
#Usage
    $ glide create                            # Start a new workspace
    $ open glide.yaml                         # and edit away!
    $ glide get github.com/Masterminds/cookoo # Get a package and add to glide.yaml
    $ glide install                           # Install packages and dependencies
    # work, work, work
    $ go build                                # Go tools work normally
    $ glide up                                # Update to newest versions of the package
    
#Detailed information

    https://github.com/Masterminds/glide
    
Fast HOWTO's

    https://glide.sh/
    
    
#How it worked at out project

To init glide yaml file

    glide init
    glide update
    
To add a new dependency

    glide get gopkg.in/alecthomas/gometalinter.v2  