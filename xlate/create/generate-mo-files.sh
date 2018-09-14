#!/usr/local/bin/bash
# Script to generate MO (machine object) files in a temp locale folder
# 
rm -rf ../view/locale
mkdir ../view/locale
cd ../view/po
for lang in $(ls *.po); do
    lang=${lang::-3}
    mkdir -p ../locale/${lang//_/-}/LC_MESSAGES
    msgfmt -c -o ../locale/${lang//_/-}/LC_MESSAGES/ghostbsd-welcome.mo $lang.po
done
