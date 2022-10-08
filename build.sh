if ! [ -x "$(command -v go)" ]; then
  echo 'Error: golang is not installed.' >&2
  echo 'Please install golang and try again.' >&2
  exit 1
fi

if [ "$(go version | awk '{print $3}' | cut -d. -f2)" -lt 11 ]; then
  echo 'Error: golang version is less than 1.11.' >&2
  echo 'Please upgrade golang and try again.' >&2
  exit 1
fi

if [ "$(uname)" == "Darwin" ]; then
    if [ -d "bin/macos" ]; then
        echo "Deleting old build..."
        rm -rf bin/macos
    fi
    echo "Building for  macOS..."
    go build -o bin/macos/
    echo "Done"
fi