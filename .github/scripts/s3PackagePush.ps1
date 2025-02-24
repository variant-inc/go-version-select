# Exit on any error
$ErrorActionPreference = "Stop"

$SourcePackage = "build/go-version-select"
$ZipPackage = "build/go-version-select.zip"

$S3Bucket = $env:GO_CLI_S3_BUCKET

$Version = $env:IMAGE_VERSION
if (-not $Version) {
    Write-Error "Error: IMAGE_VERSION environment variable is not set."
    exit 1
}

# main is stable, else pre-release and append version
$Branch = if ($BranchName -in "main", "master") { "stable" } else { "pre-release" }
$S3Key = "$Branch/$Version.zip"

if (!(Test-Path -Path $SourcePackage)) {
    Write-Error "Error: $SourcePackage does not exist."
    exit 1
}

if (Test-Path -Path $ZipPackage) {
    Remove-Item -Path $ZipPackage -Force
}

Write-Host "Compressing $SourcePackage to $ZipPackage"
Compress-Archive -Path $SourcePackage -DestinationPath $ZipPackage

Write-Host "Uploading $ZipPackage to s3://$S3Bucket/$S3Key"
aws s3 cp $ZipPackage "s3://$S3Bucket/$S3Key"

Write-Host "Upload complete to s3://$S3Bucket/$S3Key"

Add-Content -Path ${env:GITHUB_STEP_SUMMARY} `
  -Encoding utf8 `
  -Value "## Uploaded Package"

Add-Content -Path ${env:GITHUB_STEP_SUMMARY} `
  -Encoding utf8 `
  -Value "s3://$S3Bucket/$S3Key"
