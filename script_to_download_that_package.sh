#!/usr/bin/env bash
# will rename package and download last version package from gitea
mv colors old_colors
# should be started from top of package
git clone --depth=1 --branch=master https://01.kood.tech/git/Gameline/colors.git
rm -rf ./colors/.git
