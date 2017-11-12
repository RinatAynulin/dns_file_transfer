current_file = '';
current_file_name = ''
current_id = 'START';
lines = list((line.rstrip('\n') for line in open('sorted.out')))
for i in range(len(lines)):
	line = lines[i]
	if (current_id == 'START' or current_id != line.split('.')[0]):
		if (current_id != 'START'): 
			with open('output/' + current_file_name, 'w') as w:
                		w.write(current_file)	
		current_id = line.split('.')[0]
		current_file = '';
		current_file_name = line.split('.')[2].decode('hex')
	else:
		current_file += line.split('.')[2].decode('hex')
	if (i == len(lines) - 1):
		with open('output/' + current_file_name, 'w') as w:
                                w.write(current_file)
