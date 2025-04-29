# Dogipsum CLI

A Go CLI tool that generates "Lorem ipsum" style text with dog-themed words. This is a Go implementation of the original JavaScript code by Mathew Tinsley and Rick Viscomi.

## Installation

### Quick Install (Linux/macOS)

You can install dogipsum with the following one-liner:

```bash
curl -sSL https://raw.githubusercontent.com/jacobprice/dogipsum/main/install.sh | bash
```

Or with wget:

```bash
wget -qO- https://raw.githubusercontent.com/jacobprice/dogipsum/main/install.sh | bash
```

### Manual Installation

1. Download the appropriate binary for your platform from the [Releases page](https://github.com/jacobprice/dogipsum/releases)
2. Extract the archive
3. Move the binary to a location in your PATH:

```bash
# Example for Linux/macOS
chmod +x dogipsum
sudo mv dogipsum /usr/local/bin/
```

### Using Go

If you have Go installed:

```bash
go install github.com/jacobprice/dogipsum@latest
```

## Usage

Run the tool with default settings (100 words):

```bash
dogipsum
```

Specify the number of words to generate:

```bash
dogipsum -words 50
```

## Example Output

```
Doggo shibe puggorino shoob borking doggo. Doggorino extremely cuuuuuute, floofs smol pupper big ol pupper. Tungg extremely cuuuuuute smol borking doggo with a long snoot for pats extremely cuuuuuute borking doggo, such treat shooberino shoober doggo. Big ol pupper extremely cuuuuuute long woofer doge doggorino noodle horse, very taste wow fluffer blop floofs. Thicc shibe boofers, yapper heckin good boys and girls long doggo.
```

## Development

### Releasing New Versions

This project uses GitHub Actions to automate the release process:

1. Create a new tag and push it to GitHub:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. The GitHub Action will automatically:
   - Build binaries for multiple platforms (Linux, macOS, Windows)
   - Create a GitHub release with the binaries

## License

This software is provided under the MIT license. See the source code for full license details.

## Credits

- Original JavaScript implementation: Mathew Tinsley (tinsley@tinsology.net) and Rick Viscomi (rviscomi@gmail.com)
- Go implementation: [Jacob Price](https://github.com/jacobprice)