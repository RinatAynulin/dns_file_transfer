with open('sorted.out', 'w') as w:
	lines = (line.rstrip('\n')[20:][:-12] for line in open('dns.out'))
	for line in lines:
		w.write(line + '\n')
