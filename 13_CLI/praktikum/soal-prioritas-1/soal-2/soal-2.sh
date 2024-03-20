#!/bin/zsh

mkdir -p foods drinks snacks

touch foods/menu.txt drinks/menu.txt snacks/menu.txt

echo "nasi goreng" >> foods/menu.txt
echo "ayam goreng" >> foods/menu.txt
echo "bubur ayam" >> foods/menu.txt

echo "kopi susu" >> drinks/menu.txt
echo "susu oat" >> drinks/menu.txt
echo "es teh" >> drinks/menu.txt

curl -s https://gist.githubusercontent.com/nadirbslmh/c84ee3527272e52a8e6257d46e627f91/raw/74593cc3fe61ca4ff36e11ed8e3fffcfb1d0c602/snacks.json | jq -r '.[].name' > snacks/menu.txt

echo "Tugas selesai."
