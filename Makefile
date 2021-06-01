

sync:
	 rsync -avz ./nku-treehole-server ubuntu@120.53.250.69:~/app/
 	 rsync -avz ./nku-treehole-web ./nku-treehole-server  --exclude 'node_modules' ubuntu@120.53.250.69:~/app/



