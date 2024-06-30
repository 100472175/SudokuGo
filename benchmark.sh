# Launch 10 times the program ./main and time it

for i in {1..200}
do
    ./main input/9.txt 9 >/dev/null
done