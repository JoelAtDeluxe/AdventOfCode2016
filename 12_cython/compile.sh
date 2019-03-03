#! /usr/bin/env bash
cython -a -3 cylogic.pyx
python3 setup.py build_ext --inplace