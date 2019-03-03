from distutils.core import setup
from Cython.Build import cythonize

setup(name="Day 12 Puzzle", 
    ext_modules=cythonize('cylogic.pyx'))