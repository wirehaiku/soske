#!/bin/fish
# extra/build.fish: Build cross-platform Soske binaries.

set arches "386" "amd64"
set oses   "linux" "windows"

gox -arch=$arches -os=$oses -output="soske-{{ .OS }}-{{ .Arch }}"

for exe in soske-*-*
    set zip (string split "." $exe --fields 1 --max 1)
    zip -q -9 $zip changes.md license.md readme.md
    zip -q -9 -m $zip $exe
    echo "Created archive for $exe."
end
