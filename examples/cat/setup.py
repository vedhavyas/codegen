from setuptools import setup

with open("README.md", "r") as fh:
    long_description = fh.read()

requirements = [
    "click>=8.0.0",
]

setup(
    name="cat-cli-lib",
    version="0.1.0",
    author="Your Name",
    author_email="your_email@example.com",
    description="A drop-in replacement for the Linux 'cat' command as a CLI and library",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/your_username/cat-cli-lib",
    py_modules=["cat_lib", "main"],
    install_requires=requirements,
    entry_points="""
        [console_scripts]
        cat=main:cli
    """,
    classifiers=[
        "Development Status :: 3 - Alpha",
        "Intended Audience :: Developers",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
    ],
    python_requires=">=3.7",
)