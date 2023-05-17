# cat - A Linux cat Command Drop-In Replacement

This project is a command line application called `cat`, which is designed as a drop-in replacement for the Linux `cat` command. The project is written in Python 3 and serves as both a CLI app and a library for other applications. Extensive documentation, CLI usage examples, and library usage examples are included.

## Requirements

- Python 3

## Installation

To install the project, run:

```bash
pip install -r requirements.txt
```

## Usage

### Command Line Interface

To use the `cat` command line application, simply run `./main.py` followed by the file(s) you would like to concatenate:

```bash
python main.py file1.txt file2.txt
```

This will output the content of `file1.txt` and `file2.txt`.

### Library Usage

To use the `cat` library in your project, first import it:

```python
from cat_lib import cat
```

Then, you can use the `cat` function to concatenate the content of any number of files:

```python
content = cat(["file1.txt", "file2.txt"])
print(content)
```

This will return the concatenated content of `file1.txt` and `file2.txt`.

## Examples

For examples of how to use the CLI and library versions of `cat`, refer to [`examples/cli_example.py`](examples/cli_example.py) and [`examples/library_example.py`](examples/library_example.py) respectively.

## License

This project is licensed under the terms of the [MIT License](LICENSE).