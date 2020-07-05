sudo apt update

sudo apt install build-essential nload

cd ~

# Install Go
wget -c https://storage.googleapis.com/golang/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.14.4.linux-amd64.tar.gz
sudo ln -s /usr/local/go/bin/go /usr/bin/go

cd ~

# Install ntttcp
git clone https://github.com/microsoft/ntttcp-for-linux
cd ntttcp-for-linux/src
make -j 4
sudo make install

cd ~

# Download Proximity
git clone https://github.com/scottshotgg/proximity
cd proximity
git checkout origin/event-loop