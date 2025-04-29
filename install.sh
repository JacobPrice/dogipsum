#!/bin/bash
# Script to install dogipsum binary

set -e

# Default version
VERSION=${VERSION:-"latest"}
# Default install directory
INSTALL_DIR=${INSTALL_DIR:-"/usr/local/bin"}

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Installing dogipsum ${VERSION}...${NC}"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
case $ARCH in
  x86_64)
    ARCH="x86_64"
    ;;
  arm64 | aarch64)
    ARCH="arm64"
    ;;
  *)
    echo -e "${RED}Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

# Set file extension based on OS
if [ "$OS" = "windows" ]; then
  EXT=".zip"
  BIN_EXT=".exe"
else
  EXT=".tar.gz"
  BIN_EXT=""
fi

# Construct the GitHub release URL
if [ "$VERSION" = "latest" ]; then
  RELEASE_URL="https://github.com/jacobprice/dogipsum/releases/latest/download/dogipsum_${OS}_${ARCH}${EXT}"
else
  RELEASE_URL="https://github.com/jacobprice/dogipsum/releases/download/${VERSION}/dogipsum_${OS}_${ARCH}${EXT}"
fi

# Create temp directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo -e "${BLUE}Downloading from ${RELEASE_URL}...${NC}"

# Download the binary
if command -v curl > /dev/null; then
  curl -SL "$RELEASE_URL" -o "dogipsum$EXT"
elif command -v wget > /dev/null; then
  wget -q "$RELEASE_URL" -O "dogipsum$EXT"
else
  echo -e "${RED}Neither curl nor wget found. Please install one of them.${NC}"
  exit 1
fi

# Extract the binary
if [ "$OS" = "windows" ]; then
  unzip "dogipsum$EXT"
else
  tar xzf "dogipsum$EXT"
fi

# Check if the binary exists
if [ ! -f "dogipsum$BIN_EXT" ]; then
  # Look for the binary in a subdirectory
  BINARY_PATH=$(find . -name "dogipsum$BIN_EXT" -type f | head -n 1)
  if [ -z "$BINARY_PATH" ]; then
    echo -e "${RED}Could not find dogipsum binary in the downloaded archive.${NC}"
    exit 1
  fi
  cp "$BINARY_PATH" "dogipsum$BIN_EXT"
fi

# Make it executable
chmod +x "dogipsum$BIN_EXT"

# Create the install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Copy the binary to the install directory
if [ -w "$INSTALL_DIR" ]; then
  cp "dogipsum$BIN_EXT" "$INSTALL_DIR/dogipsum$BIN_EXT"
else
  echo -e "${BLUE}Copying dogipsum to $INSTALL_DIR requires root privileges...${NC}"
  sudo cp "dogipsum$BIN_EXT" "$INSTALL_DIR/dogipsum$BIN_EXT"
fi

echo -e "${GREEN}dogipsum has been installed to $INSTALL_DIR/dogipsum${BIN_EXT}${NC}"

# Clean up
cd - > /dev/null
rm -rf "$TMP_DIR"

echo -e "${GREEN}Installation complete!${NC}"
echo -e "${BLUE}Run 'dogipsum' to generate some dog-themed lorem ipsum text.${NC}"