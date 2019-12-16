gRPC ruby/go example
-------------------

### Minimal setup
#### Go

1. Install `goenv`
  ```
  git clone https://github.com/syndbg/goenv.git ~/.goenv
  ```
  ```
  echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bash_profile
  echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bash_profile
  ```
  ```
  echo 'eval "$(goenv init -)"' >> ~/.bash_profile
  ```

2. Install `dep`

  `brew install dep` or `sudo apt-get install go-dep`

3. Setup your `$GOPATH`

4. Vendor dependencies

  `cd <path-to-project>/go && dep ensure`

### Docker
