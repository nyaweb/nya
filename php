#!/bin/bash
apt update 
apt upgrade -y
apt install libxml2-dev libsqlite3-dev build-essential curl -y

cd ~/
curl -OJL https://github.com/php/php-src/archive/refs/tags/php-8.5.0.tar.gz

tar -zxvf php*.tar.gz
rm -rf php*.tar.gz
cd php*/
./buildconf --force
./configure
make
make test
make install
cd ~/
rm -rf php*/

php -v