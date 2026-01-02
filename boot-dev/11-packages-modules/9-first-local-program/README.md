```bash
mkdir hellogo
cd hellogo
go mod init {REMOTE}/{USERNAME}/hellogo



```
# GO run
quickly compile and runs a go package. The compiled binary is not saved in your working directory.

# GO build
The go build command compiles go code into a single, statically linked executable program. One of the beauties of Go is that you always go build for production, and because the output is a statically compiled binary, you can ship it to production or end users without them needing the Go toolchain installed.

Some new Go devs use go run on a server in production, which is a huge mistake.

# GO install
The go install command compiles and installs a package or packages on your local machine for your personal usage. It installs the package's compiled binary in the GOBIN directory. (We installed the bootdev cli with it, after all)

Assignment
Ensure you are in your hellogo repo, then run:
go install

Navigate out of your project directory:
cd ../

If all went well, Go compiled and installed the hellogo program globally on your machine. Run it with:
hellogo

Run and submit the CLI tests from anywhere on your machine.