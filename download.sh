#!/bin/bash

set -e          # Exit on error
set -u          # Treat unset variables as an error
set -o pipefail # Exit on errors in pipelines

if [ -z "$GO_CLI_VERSION" ]; then
	echo "GO_CLI_VERSION is not set. Using default version from file."
	GO_CLI_VERSION=$(yq '.version' "$CONFIG_FILE_PATH")
fi

s3Bucket="$DX_PACKAGES_S3_BUCKET"
filename="go-version-select.${GO_CLI_VERSION}.zip"
prefix="go-version-select/$filename"
localDir="/tmp/go-version-download"

aws s3 cp "s3://$s3Bucket/$prefix" "$localDir/" --force --debug

# Unzip the file
unzip -o "$localDir/$filename" -d "$localDir"
echo "Unzip completed. Extracted to: $localDir"

mkdir -p "$HOME/.local/bin/"

mv "$localDir/go-version-select" "$HOME/.local/bin/"
chmod +x "$HOME/.local/bin/go-version-select"

echo "$HOME/.local/bin" >>"$GITHUB_PATH"
