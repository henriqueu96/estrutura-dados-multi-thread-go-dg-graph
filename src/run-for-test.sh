sh ./run.sh 64 250 0 100 3 ./results/result-with-250-population-no-conflicts.txt
sh ./run.sh 64 250 0.1 100 3 ./results/result-with-250-population-10percent-conflicts.txt
sh ./run.sh 64 250 0.2 100 3 ./results/result-with-250-population-20percent-conflicts.txt
sh ./run.sh 64 250 1 100 3 ./results/result-with-250-population-all-conflicts.txt

sh ./run.sh 64 100 0 100 3 ./results/result-with-100-population-no-conflicts.txt
sh ./run.sh 64 100 0.1 100 3 ./results/result-with-100-population-10percent-conflicts.txt
sh ./run.sh 64 100 0.2 100 3 ./results/result-with-100-population-20percent-conflicts.txt
sh ./run.sh 64 100 1 100 3 ./results/result-with-100-population-all-conflicts.txt
