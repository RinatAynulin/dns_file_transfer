from dnslib.server import *
import datetime

class TestResolver:
	def resolve(self,request,handler):
		if (not request.q.qname.matchGlob('*.lohcoin.ru')):
			return request.reply()
		reply_fromZone = request.q.qname.__str__() + ' 60 '
		if (request.q.qtype == 1):
			with open("dns.out", "a") as myfile:
    				myfile.write(datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S") + ' ' + request.q.qname.__str__() + '\n')
			reply_fromZone += 'A 62.109.2.225'
		if (request.q.qtype == 28):
                        # reply_fromZone += 'AAAA 0:0:0:0:0:0:0:1'		
			return request.reply()
		if (request.q.qtype == 15):
			return request.reply()
		reply = request.reply()
		reply.add_answer(*RR.fromZone(reply_fromZone.__str__()))
		return reply

resolver = TestResolver()
server = DNSServer(resolver,port=53,address="62.109.2.225",tcp=False)
server.start_thread()

while True:
	pass
