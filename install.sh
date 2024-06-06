#!/bin/bash

PROGRAM_NAME="gogithook"

INSTALL_DIR="./bin"


install() {
    echo "Building $PROGRAM_NAME..."
    go build -o "$INSTALL_DIR/$PROGRAM_NAME" .

    if [ $? -eq 0 ]; then
        echo "Installation complete."
        echo "Setting permissions..."
        chmod +x "$INSTALL_DIR/$PROGRAM_NAME"
        echo "Installation finished."
    else
        echo "Failed to build $PROGRAM_NAME."
        exit 1
    fi
}

uninstall() {
    if [ -e "$INSTALL_DIR/$PROGRAM_NAME" ]; then
        echo "Removing $PROGRAM_NAME from $INSTALL_DIR..."
        rm "$INSTALL_DIR/$PROGRAM_NAME"

        if [ $? -eq 0 ]; then
            echo "Uninstallation complete."
        else
            echo "Failed to remove $PROGRAM_NAME from $INSTALL_DIR."
            exit 1
        fi
    else
        echo "$PROGRAM_NAME is not installed in $INSTALL_DIR."
        exit 1
    fi
}

if [ $# -eq 0 ]; then
    echo "Usage: $0 [install|uninstall]"
    exit 1
fi

case "$1" in
    "install")
        install
        ;;
    "uninstall")
        uninstall
        ;;
    *)
        echo "Invalid argument. Usage: $0 [install|uninstall]"
        exit 1
        ;;
esac

