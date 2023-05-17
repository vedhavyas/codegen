import os

def cat_files(file_list, output=None):
    if output:
        with open(output, "w") as outfile:
            for file in file_list:
                with open(file, "r") as infile:
                    outfile.write(infile.read())
    else:
        for file in file_list:
            with open(file, "r") as infile:
                print(infile.read(), end="")