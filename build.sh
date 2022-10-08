if [ "$(uname)" == "Darwin" ]; then
    if [ -d "bin/macos" ]; then
        echo "Deleting old build..."
        rm -rf bin/macos
    fi
    echo "Building for MacOS"
    go build -o bin/macos/
    echo "Done"
fi