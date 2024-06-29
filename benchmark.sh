# Launch 10 times the program ./main and time it

for i in {1..200}
do
    ./main input.txt 9 >/dev/null
done