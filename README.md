# aest

> A command-line wrapper for AES

## Installation

Prerequisites:

- [Git](https://git-scm.com)
- [Go](https://go.dev) 1.22 or later

Run the following commands to build aest:

```sh
git clone https://github.com/interrrp/aest
cd aest
go build -ldflags "-s -w" -o aest .
# aest should now be built at this point, try it:
./aest
```

## Usage

```text
aest <encrypt|decrypt> <input file> <output file> <key>
```

### Examples

Encrypt an image named `image.png` with the key `verysecurekey` and save it to `image.enc.png`:

```sh
aest encrypt image.png image.enc.png verysecurekey
```

Decrypt the image above to `image.dec.png`:

```sh
aest decrypt image.enc.png image.dec.png verysecurekey
```
