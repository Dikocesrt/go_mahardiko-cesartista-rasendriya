#!/bin/bash

echo -e "Masukkan nama proyek (tanpa spasi, boleh mengandung hypen): \c"
read project_name

if [[ "$project_name" == *" "* ]]; then
    echo "Nama proyek tidak boleh mengandung spasi."
    exit 1
fi

echo -e "Pilih jenis proyek:\n1. Simple\n2. API\nMasukkan pilihan (1/2): \c"
read project_type

if [[ "$project_type" != "1" && "$project_type" != "2" ]]; then
    echo "Pilihan tidak valid. Silakan masukkan 1 atau 2."
    exit 1
fi

mkdir "$project_name"
cd "$project_name"

go mod init "$project_name"

if [[ "$project_type" == "1" ]]; then
    touch main.go
    echo "package main" > main.go
    echo "" >> main.go
    echo "import \"fmt\"" >> main.go
    echo "" >> main.go
    echo "func main() {" >> main.go
    echo "    fmt.Println(\"Hello, World!\")" >> main.go
    echo "}" >> main.go

    echo "Proyek simple dengan nama $project_name telah dibuat."
    exit 0
fi

if [[ "$project_type" == "2" ]]; then
    mkdir routes models repositories services configs controllers
    touch main.go

    echo "package main" > main.go
    echo "" >> main.go
    echo "import (" >> main.go
    echo "    \"fmt\"" >> main.go
    echo ")" >> main.go
    echo "" >> main.go
    echo "func main() {" >> main.go
    echo "    fmt.Println(\"API $project_name is running\")" >> main.go
    echo "}" >> main.go

    echo "Proyek API dengan nama $project_name telah dibuat."
    exit 0
fi
