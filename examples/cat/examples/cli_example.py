import os
import sys
from main import main

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python cli_example.py <filename>")
        sys.exit(1)
    
    filename = sys.argv[1]
    if not os.path.isfile(filename):
        print(f"Error: {filename} does not exist")
        sys.exit(1)
    
    with open(filename, "r") as file:
        content = file.read()
        
    main(content)