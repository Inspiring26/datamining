import os

for fname in os.listdir("./data"):
	with open("./data/"+fname) as f:
		tmp=f.readlines()