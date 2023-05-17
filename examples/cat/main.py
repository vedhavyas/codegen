import argparse
from cat_lib import cat_files

def main():
    parser = argparse.ArgumentParser(description="A drop-in replacement for the Linux 'cat' command.")
    parser.add_argument("filepath", nargs="+", help="Filepath(s) for the files to be read.")
    args = parser.parse_args()

    result = cat_files(args.filepath)
    print(result, end="")

if __name__ == "__main__":
    main()