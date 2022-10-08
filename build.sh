# check if golang is installed else ask to install
if ! [ -x "$(command -v go)" ]; then
  echo 'Error: golang is not installed.' >&2
  echo 'Please install golang and try again.' >&2
  exit 1
fi

if [ "$(uname)" == "Darwin" ]; then
    if [ -d "bin/macos" ]; then
        echo "Deleting old build..."
        rm -rf bin/macos
    fi
    echo "Building for MacOS"
    go build -o bin/macos/
    echo "Done"
fi