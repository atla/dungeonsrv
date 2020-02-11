echo "--- Backend:"
find . -name '*.go' | xargs wc -l
echo "--- Frontend:"
find ./dist/app/src -name '*.js' -or -name '*.vue' | xargs wc -l

