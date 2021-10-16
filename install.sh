#!/bin/sh -e

##
# This script makes certain assumptions about your deployment of SimpleStatus
# 1) You are deploying to at least Ubuntu 12.04 LTS
# 2) the user running the script has nopasswd access to sudo
# 3) using upstart to manage the daemonizing
# This script should be viewed as a template for your own deployment and should
# not be run blindly

##
# Set up build directory



cd $TEMP/
git clone https://github.com/oknors/status.git

cd status
## 
# We know where generate_cert.go lives when installed from ppa
# but, if you want to find it yourself, uncomment the line below
# and use the $CRT variable in the build stage
#export CRT=$(sudo find / -name generate_cert.go)

##


##
# Build package, generate certs, and modify permissions
go build -o ostat
sudo cp ostat /usr/local/bin/ostat
##
# This is just one way to deal changing the default configuration of the package
# we're using upstart here and setting some runtime flags

sudo cp ostat.conf /etc/systemd/system/ostat.service
sudo systemctl enable ostat.service
sudo systemctl daemon_restart
sudo systemctl start ostat.service
sudo systemctl status ostat.service
