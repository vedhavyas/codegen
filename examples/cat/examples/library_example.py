from cat_lib import cat_files

# List of file paths to concatenate
filepaths = ["file1.txt", "file2.txt", "file3.txt"]

# Concatenate the content of the given files
output = cat_files(filepaths)

# Print the concatenated content to the console
print(output)